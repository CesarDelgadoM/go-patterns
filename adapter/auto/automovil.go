package auto

import "fmt"

type Automovil struct {
	Marca     string
	Model     uint8
	Encendido bool
}

func (a *Automovil) Encender() {
	if a.Encendido {
		fmt.Println("Ya esta encendido!")

	} else {
		fmt.Println("Encendido!")
	}
}

func (a *Automovil) Acelerar() {
	fmt.Println("Acelerando...")
}
