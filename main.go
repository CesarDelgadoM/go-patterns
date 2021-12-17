package main

import (
	"fmt"
	"os"
	"sync"

	"github.com/cesardelgadom/go-patterns/builder"
	"github.com/cesardelgadom/go-patterns/factory"
	"github.com/cesardelgadom/go-patterns/singleton"
	"github.com/cesardelgadom/go-patterns/singleton/client_one"
	"github.com/cesardelgadom/go-patterns/singleton/client_two"
)

func main() {
	testSingleton()
	testFactory()
	testBuilder()
}

func testSingleton() {
	wg := sync.WaitGroup{}
	wg.Add(200)
	for i := 0; i < 100; i++ {
		go func() {
			defer wg.Done()
			client_one.IncrementAge()
		}()
		go func() {
			defer wg.Done()
			client_two.IncrementAge()
		}()
	}
	wg.Wait()
	p := singleton.GetInstance()
	fmt.Println(p.GetAge())
}

func testFactory() {
	var typeDb int8
	fmt.Print("Digite la conexion deseada: ")
	_, err := fmt.Scan(&typeDb)
	if err != nil {
		fmt.Printf("Error al leer la opcion: %v", err)
		os.Exit(1)
	}
	conn := factory.Factory(int(typeDb))
	if conn == nil {
		fmt.Println("Motor DB no valido!")
		os.Exit(1)
	}
	conn.Connect()
	fmt.Println(conn.GetNow())
	conn.Close()
}

func testBuilder() {
	var op int8
	var typeObj builder.MessageBuilder
	for typeObj == nil {
		fmt.Println("Digite el formato a construir (0:JSON, 1:XML): ")
		_, err := fmt.Scan(&op)
		if err != nil {
			fmt.Printf("Error al leer la opcion: %v", err)
			os.Exit(3)
		}
		typeObj = messageFactory(op)
	}
	sender := builder.Sender{}
	sender.SetBuilder(typeObj)
	data, err := sender.BuildMessage("PATTERNS-GO", "Pattern builder")
	if err != nil {
		fmt.Printf("Error al construir el tipo dato: %v", err)
		os.Exit(3)
	}
	fmt.Println(string(data.Body))
}

func messageFactory(op int8) builder.MessageBuilder {
	switch op {
	case 0:
		fmt.Println("Ha seleccionado JSON")
		return &builder.JSONMessageBuilder{}
	case 1:
		fmt.Println("Ha seleccionado XML")
		return &builder.XMLMessageBuilder{}
	default:
		return nil
	}
}
