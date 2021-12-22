package local

import "github.com/cesardelgadom/go-patterns/proxy/book"

type Proxy interface {
	GetByID(ID uint) book.Book
	GetAll() book.Books
}
