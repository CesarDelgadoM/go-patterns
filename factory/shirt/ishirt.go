package shirt

type IShirt interface {
	setLogo(Logo string)
	setSize(Size int)
	getLogo() string
	getSize() int
}

type Shirt struct {
	Logo string
	Size int
}

func (s *Shirt) setLogo(Logo string) {
	s.Logo = Logo
}

func (s *Shirt) setSize(Size int) {
	s.Size = Size
}

func (s *Shirt) getLogo() string {
	return s.Logo
}

func (s *Shirt) getSize() int {
	return s.Size
}
