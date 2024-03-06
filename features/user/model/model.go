package model

import (
	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/pkg/validation"
)

type User struct {
	Admin    bool   `json:"admin"`
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
}

func (u *User) FromDomain(du domain.User) {
	u.Admin = du.Admin
	u.ID = du.ID
	u.Name = du.Name
	u.UserName = du.UserName
}

func UsersFromDomain(dUsers []domain.User) []User {
	users := make([]User, len(dUsers))
	for i, du := range dUsers {
		user := new(User)
		user.FromDomain(du)

		users[i] = *user
	}

	return users
}

type CreateParams struct {
	Admin          bool   `form:"admin"`
	Name           string `form:"name"`
	Password       string `form:"password"`
	RepeatPassword string `form:"repeat_password"`
	UserName       string `form:"username"`
}

func (params CreateParams) Validate() error {
	if validation.IsEmpty(params.Name) {
		return ErrNoName
	}

	if validation.IsEmpty(params.Password) {
		return ErrNoPassword
	}

	if params.Password != params.RepeatPassword {
		return ErrPasswordNoMatch
	}

	if validation.IsEmpty(params.UserName) {
		return ErrNoUserName
	}

	return nil
}

func (params CreateParams) Domain() domain.CreateUserParams {
	return domain.CreateUserParams{
		Admin:    true,
		Name:     params.Name,
		Password: params.Password,
		UserName: params.UserName,
	}
}
