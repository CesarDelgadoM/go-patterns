package singleton

import "sync"

var (
	p    *person
	mux  sync.Mutex
	once sync.Once
)

func GetInstance() *person {
	once.Do(func() {
		p = &person{}
	})
	return p
}

type person struct {
	name string
	age  int16
}

func (p *person) SetName(name string) {
	p.name = name
}

func (p *person) GetName() string {
	return p.name
}

func (p *person) IncrementAge() {
	mux.Lock()
	p.age++
	mux.Unlock()
}

func (p *person) GetAge() int16 {
	mux.Lock()
	defer mux.Unlock()
	return p.age
}
