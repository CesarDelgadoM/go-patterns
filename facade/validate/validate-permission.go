package validate

import (
	"errors"
	"fmt"
)

var ErrPermissionNotValid = errors.New("el usuario no esta autorizado")

type ValidatorPermission struct {
	userID string
}

func NewValidatorPermission(ID string) ValidatorPermission {
	return ValidatorPermission{
		userID: ID,
	}
}

func (vp *ValidatorPermission) Validate() error {
	if vp.userID != "user-blog" {
		return ErrPermissionNotValid
	}
	fmt.Println("Usuario autorizado")
	return nil
}
