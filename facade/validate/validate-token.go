package validate

import (
	"errors"
	"fmt"
)

var ErrTokenNotValid = errors.New("el usuario no esta logueado")

type ValidatorToken struct {
	token string
}

func NewValidatorToken(t string) ValidatorToken {
	return ValidatorToken{
		token: t,
	}
}

func (vt *ValidatorToken) Validate() error {
	if vt.token != "token-valido" {
		return ErrTokenNotValid
	}
	fmt.Println("token valido")
	return nil
}
