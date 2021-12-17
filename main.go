package main

import (
	"fmt"
	"os"
	"sync"

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

func testBuilder() {}
