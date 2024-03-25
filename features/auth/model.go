package auth

import (
	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/pkg/escape"
	"github.com/edgarSucre/jw/pkg/validate"
)

type LoginParams struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

const (
	errMsgNoUserName = "usuario es requerido"
	errMSgNoPassword = "contraseña es requerido"
)

func (params LoginParams) Validate() map[string]string {
	errors := make(map[string]string)

	if validate.IsEmpty(params.UserName) {
		errors["username"] = errMsgNoUserName
	}

	if validate.IsEmpty(params.Password) {
		errors["password"] = errMSgNoPassword
	}

	return errors
}

func (params *LoginParams) Sanitize() {
	params.UserName = escape.AlphaNumeric(params.UserName, false)
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
