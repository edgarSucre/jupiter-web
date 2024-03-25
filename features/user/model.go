package user

import (
	"fmt"

	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/features/user/view"
	"github.com/edgarSucre/jw/pkg/escape"
	"github.com/edgarSucre/jw/pkg/validate"
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

func (u User) toView() view.User {
	return view.User{
		Admin:   u.Admin,
		Name:    u.Name,
		TemplID: fmt.Sprintf("user-%v", u.ID),
	}
}

func usersFromDomain(dUsers []domain.User) []User {
	users := make([]User, len(dUsers))
	for i, du := range dUsers {
		user := new(User)
		user.FromDomain(du)

		users[i] = *user
	}

	return users
}

func usersToView(users []User) []view.User {
	viewUsers := make([]view.User, len(users))
	for i, u := range users {
		viewUsers[i] = u.toView()
	}

	return viewUsers
}

type CreateParams struct {
	Admin          bool   `form:"admin"`
	Name           string `form:"name"`
	Password       string `form:"password"`
	RepeatPassword string `form:"repeat_password"`
	UserName       string `form:"username"`
}

const (
	errMsgNoName          = "nombre es requerido"
	errMsgNoUserName      = "nombre de usuario es requerido"
	errMSgNoPassword      = "contraseña es requerido"
	errMsgPasswordNoMatch = "contraseñas no son iguales"
)

func (params CreateParams) Validate() map[string]string {
	errors := make(map[string]string)

	if validate.IsEmpty(params.Name) {
		errors["name"] = errMsgNoName
	}

	if validate.IsEmpty(params.Password) {
		errors["password"] = errMSgNoPassword
	}

	if params.Password != params.RepeatPassword {
		errors["repeat_password"] = errMsgPasswordNoMatch
	}

	if validate.IsEmpty(params.UserName) {
		errors["username"] = errMsgNoUserName
	}

	return errors
}

func (params *CreateParams) Sanitize() {
	params.Name = escape.Alpha(params.Name, true)
	params.UserName = escape.AlphaNumeric(params.UserName, false)
	params.Password = escape.Password(params.Password)
	params.RepeatPassword = escape.Password(params.RepeatPassword)
}

func (params CreateParams) View() view.UserForm {
	return view.UserForm{
		Name:           params.Name,
		Password:       params.Password,
		RepeatPassword: params.RepeatPassword,
		UserName:       params.UserName,
	}
}

func (params CreateParams) Domain() domain.CreateUserParams {
	return domain.CreateUserParams{
		Admin:    true,
		Name:     params.Name,
		Password: params.Password,
		UserName: params.UserName,
	}
}
