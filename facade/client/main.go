package main

import (
	"log"

	"github.com/cesardelgadom/go-patterns/facade"
)

func main() {
	token := "token-valido"
	user := "user-blog"
	to := "correo@dominio.com"
	comment := "aprendiendo patrones con golang"

	f := facade.New(to, comment, token, user)
	err := f.Comment()
	if err != nil {
		log.Fatal(err)
	}

	f.Notify()
}
