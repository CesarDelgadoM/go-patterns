package facade

import (
	"github.com/cesardelgadom/go-patterns/facade/email"
	"github.com/cesardelgadom/go-patterns/facade/storage"
	"github.com/cesardelgadom/go-patterns/facade/validate"
)

type Facade struct {
	to                  string
	comment             string
	validatorToken      validate.ValidatorToken
	validatorPermission validate.ValidatorPermission
	store               storage.Storage
	notificator         email.Email
}

func New(to, comment, token, user string) Facade {
	return Facade{
		to:                  to,
		comment:             comment,
		validatorToken:      validate.NewValidatorToken(token),
		validatorPermission: validate.NewValidatorPermission(user),
		store:               storage.NewStorage("redis"),
		notificator:         email.NewEmail(),
	}
}

//Toda la responsabilidad de operaciones
//lo tiene la fachada y no el cliente
func (f *Facade) Comment() error {
	if err := f.validatorToken.Validate(); err != nil {
		return err
	}

	if err := f.validatorPermission.Validate(); err != nil {
		return err
	}
	f.store.Save(f.comment)
	f.notificator.Send(f.to, f.comment)
	return nil
}

//Puede hacer procesos especificos
func (f *Facade) Notify() {
	f.notificator.Send(f.to, f.comment)
}
