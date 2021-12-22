package adapter

import (
	"github.com/cesardelgadom/go-patterns/adapter/auto"
	"github.com/cesardelgadom/go-patterns/adapter/bici"
)

func Factory(op string) Adapter {
	switch op {
	case "automovil":
		a := auto.Automovil{}
		return &auto.AutomovilAdapter{
			Automovil: &a,
		}
	case "bicicleta":
		b := bici.Bicicleta{}
		return &bici.BicicletaAdapter{
			Bicicleta: &b,
		}
	}
	return nil
}
