package main

import (
	"fmt"

	"github.com/cesardelgadom/go-patterns/factory/shoe"
)

func main() {
	adidasFactory, _ := getSportFactory("adidas")
	adidasShoe := adidasFactory.MakeShoe()
	PrintShoeDetails(adidasShoe)

	nikeFactory, _ := getSportFactory("nike")
	nikeShoe := nikeFactory.MakeShoe()
	PrintShoeDetails(nikeShoe)
}

func PrintShoeDetails(s shoe.IShoe) {
	fmt.Printf("Logo: %s\n", s.GetLogo())
	fmt.Printf("Size: %d\n", s.GetSize())
}
