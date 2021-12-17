package client_two

import "github.com/cesardelgadom/go-patterns/singleton"

func IncrementAge() {
	p := singleton.GetInstance()
	p.IncrementAge()
}
