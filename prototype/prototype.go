package prototype

import "fmt"

type Prototype struct {
	Name    string
	Age     int
	Friends []string
	Color   *string
	Phones  map[string]string
}

func (p *Prototype) ToString() string {
	return fmt.Sprintf(
		"Nombre: %s, Edad: %d, Amigos: %v, Color: %s, Phones: %v",
		p.Name,
		p.Age,
		p.Friends,
		*p.Color,
		p.Phones,
	)
}

func (p *Prototype) Clone() Prototype {
	friends := make([]string, len(p.Friends))
	copy(friends, p.Friends)

	color := *p.Color

	phones := make(map[string]string)
	for i, phone := range p.Phones {
		phones[i] = phone
	}

	return Prototype{
		Name:    p.Name,
		Age:     p.Age,
		Friends: friends,
		Color:   &color,
		Phones:  phones,
	}
}
