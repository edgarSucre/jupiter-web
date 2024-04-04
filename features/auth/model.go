package auth

import (
	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/pkg/escape"
	"github.com/edgarSucre/jw/pkg/validate"
)

type LoginParams struct {
	Email    string `form:"email"`
	Password string `form:"password"`
}

const (
	errMsgNoEmail    = "email es requerido"
	errMSgNoPassword = "contraseña es requerido"
)

func (params LoginParams) Validate() map[string]string {
	errors := make(map[string]string)

	if validate.IsEmpty(params.Email) {
		errors["email"] = errMsgNoEmail
	}

	if validate.IsEmpty(params.Password) {
		errors["password"] = errMSgNoPassword
	}

	return errors
}

func (params *LoginParams) Sanitize() {
	params.Email = escape.Email(params.Email)
	params.Password = escape.Password(params.Password)
}

type User struct {
	Admin bool
	Name  string
}

func (u *User) FromDomain(du domain.User) {
	u.Admin = du.Admin
	u.Name = du.Name
}
