package model

import "github.com/edgarSucre/jw/domain"

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

type CreateUserParams struct {
	Admin          bool   `json:"admin"`
	Name           string `json:"name"`
	Password       string `json:"password"`
	RepeatPassword string `json:"repeat_password"`
	UserName       string `json:"username"`
}

func (params CreateUserParams) Validate() error {
	if IsEmpty(params.Name) {
		return ErrNoName
	}

	if IsEmpty(params.Password) {
		return ErrNoPassword
	}

	if params.Password != params.RepeatPassword {
		return ErrPasswordNoMatch
	}

	if IsEmpty(params.UserName) {
		return ErrNoUserName
	}

	return nil
}

func (params CreateUserParams) Domain() domain.CreateUserParams {
	return domain.CreateUserParams{
		Admin:    params.Admin,
		Name:     params.Name,
		Password: params.Password,
		UserName: params.UserName,
	}
}
