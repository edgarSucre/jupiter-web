package model

import "github.com/edgarSucre/jw/domain"

type LoginParams struct {
	UserName string `form:"username"`
	Password string `form:"password"`
}

type User struct {
	Admin bool
	Name  string
}

func (u *User) FromDomain(du domain.User) {
	u.Admin = du.Admin
	u.Name = du.Name
}
