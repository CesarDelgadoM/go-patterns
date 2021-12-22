package auto

type AutomovilAdapter struct {
	Automovil *Automovil
}

func (a *AutomovilAdapter) Mover() {
	a.Automovil.Encender()
	a.Automovil.Acelerar()
}
