package local

import (
	"github.com/cesardelgadom/go-patterns/proxy/book"
	"github.com/cesardelgadom/go-patterns/proxy/remote"
)

type Local struct {
	Remote *remote.Data
	cache  book.Books
}

func New() Proxy {
	return &Local{
		Remote: remote.New("http://algunaDB.com", 8080, "pointer", "******"),
		cache:  make(book.Books, 0),
	}
}

func (l *Local) GetByID(ID uint) book.Book {
	//Buscamos primero en la lista de cache
	for _, book := range l.cache {
		if book.ID == ID {
			return book
		}
	}
	//Si no se encuentra lo traemos de
	//la db y lo agregamos a la cache
	b := l.Remote.ByID(ID)
	l.cache = append(l.cache, b)
	return b
}

func (l *Local) GetAll() book.Books {
	return l.Remote.All()
}
