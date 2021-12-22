package bici

type BicicletaAdapter struct {
	Bicicleta *Bicicleta
}

func (b *BicicletaAdapter) Mover() {
	b.Bicicleta.Avanzar()
}
