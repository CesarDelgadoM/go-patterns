package remote

import (
	"time"

	"github.com/cesardelgadom/go-patterns/proxy/book"
)

type Data struct {
	books    book.Books
	server   string
	port     uint16
	user     string
	password string
}

func New(server string, port uint16, user string, password string) *Data {
	data := &Data{
		server:   server,
		port:     port,
		user:     user,
		password: password,
	}
	data.load()
	return data
}

func (data *Data) ByID(ID uint) book.Book {
	time.Sleep(2 * time.Second)
	for _, book := range data.books {
		if book.ID == ID {
			return book
		}
	}
	return book.Book{}
}

func (data *Data) All() book.Books {
	time.Sleep(4 * time.Second)
	return data.books
}

func (data *Data) load() {
	data.books = make(book.Books, 0, 5)
	data.books = append(data.books,
		book.Book{ID: 1, Name: "El arte de la guerra", Author: "Sun Tzu"},
		book.Book{ID: 2, Name: "La pelota no entra por azar", Author: "Ferran Soriano"},
		book.Book{ID: 3, Name: "Jesus, CEO", Author: "Laurie Beth"},
		book.Book{ID: 4, Name: "La biografia de Steve Jobs", Author: "Walter Isaacson"},
		book.Book{ID: 5, Name: "Pequeño cerdo capitalista", Author: "Sofia Macias"},
	)
}
