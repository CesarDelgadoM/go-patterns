package nike

import (
	"github.com/cesardelgadom/go-patterns/factory/shirt"
	"github.com/cesardelgadom/go-patterns/factory/shoe"
)

type Nike struct {
}

func (n *Nike) MakeShoe() shoe.IShoe {
	return &NikeShoe{
		Shoe: shoe.Shoe{
			Logo: "nike",
			Size: 14,
		},
	}
}

func (n *Nike) MakeShirt() shirt.IShirt {
	return &NikeShirt{
		Shirt: shirt.Shirt{
			Logo: "nike",
			Size: 14,
		},
	}
}
