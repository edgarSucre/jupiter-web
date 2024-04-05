package user

import (
	"fmt"
	"net/mail"

	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/features/user/view"
	"github.com/edgarSucre/jw/pkg/escape"
	"github.com/edgarSucre/jw/pkg/validate"
)

type User struct {
	Admin bool   `json:"admin"`
	Email string `json:"email"`
	ID    int64  `json:"id"`
	Name  string `json:"name"`
}

func (u *User) FromDomain(du domain.User) {
	u.Admin = du.Admin
	u.Email = du.Email
	u.ID = du.ID
	u.Name = du.Name
}

func (u User) toView() view.User {
	return view.User{
		Admin:   u.Admin,
		Email:   u.Email,
		ID:      fmt.Sprint(u.ID),
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
	Email          string `form:"email"`
}

const (
	errMsgBadEmail         = "email es invalido"
	errMsgNoName           = "nombre es requerido"
	errMsgNoEmail          = "email es requerido"
	errMSgNoPassword       = "contraseña es requerido"
	errMsgPasswordNoMatch  = "contraseñas no son iguales"
	errMsgPasswordTooShort = "contraseña debe tener almenos 8 caracteres"
)

func (params CreateParams) Validate() map[string]string {
	errors := make(map[string]string)

	if validate.IsEmpty(params.Name) {
		errors["name"] = errMsgNoName
	}

	if validate.IsEmpty(params.Password) {
		errors["password"] = errMSgNoPassword
	}

	if len([]rune(params.Password)) <= 7 {
		errors["password"] = errMsgPasswordTooShort
	}

	if params.Password != params.RepeatPassword {
		errors["repeat_password"] = errMsgPasswordNoMatch
	}

	if validate.IsEmpty(params.Email) {
		errors["email"] = errMsgNoEmail
	}

	if _, err := mail.ParseAddress(params.Email); err != nil {
		errors["email"] = errMsgBadEmail
	}

	return errors
}

func (params *CreateParams) Sanitize() {
	params.Name = escape.Alpha(params.Name, true)
	params.Email = escape.Email(params.Email)
	params.Password = escape.Password(params.Password)
	params.RepeatPassword = escape.Password(params.RepeatPassword)
}

func (params CreateParams) View() view.UserForm {
	return view.UserForm{
		Admin:          fmt.Sprint(params.Admin),
		Email:          params.Email,
		Name:           params.Name,
		Password:       params.Password,
		RepeatPassword: params.RepeatPassword,
	}
}

func (params CreateParams) Domain() domain.CreateUserParams {
	return domain.CreateUserParams{
		Admin:    params.Admin,
		Email:    params.Email,
		Name:     params.Name,
		Password: params.Password,
	}
}
