package main

import (
	"errors"

	"github.com/cesardelgadom/go-patterns/factory/adidas"
	"github.com/cesardelgadom/go-patterns/factory/nike"
	"github.com/cesardelgadom/go-patterns/factory/shirt"
	"github.com/cesardelgadom/go-patterns/factory/shoe"
)

type ISportsFactory interface {
	MakeShoe() shoe.IShoe
	MakeShirt() shirt.IShirt
}

func getSportFactory(brand string) (ISportsFactory, error) {
	switch brand {
	case "adidas":
		return &adidas.Adidas{}, nil
	case "nike":
		return &nike.Nike{}, nil
	default:
		return nil, errors.New("Not brand exist")
	}
}
