package model

import "github.com/edgarSucre/jw/domain"

type LoginParams struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

func (lp LoginParams) Validate() error {
	err := domain.NewValidationError("Datos incompletos, por favor revisar formulario")

	if lp.UserName == "" {
		err.Append("username", "Nombre de usuario es requerido")
	}

	if lp.Password == "" {
		err.Append("password", "Password es requeredio")
	}

	return err
}

type User struct {
	Admin bool
	Name  string
}

func (u *User) FromDomain(du domain.User) {
	u.Admin = du.Admin
	u.Name = du.Name
}
