package adidas

import (
	"github.com/cesardelgadom/go-patterns/factory/shirt"
	"github.com/cesardelgadom/go-patterns/factory/shoe"
)

type Adidas struct {
}

func (a *Adidas) MakeShoe() shoe.IShoe {
	return &adidasShoe{
		Shoe: shoe.Shoe{
			Logo: "adidas",
			Size: 16,
		},
	}
}

func (a *Adidas) MakeShirt() shirt.IShirt {
	return &adidasShirt{
		Shirt: shirt.Shirt{
			Logo: "adidas",
			Size: 16,
		},
	}
}
