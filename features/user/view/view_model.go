package view

import (
	"fmt"

	"github.com/edgarSucre/jw/features/user/model"
)

type User struct {
	Admin   bool
	TemplID string
	Name    string
}

func (u *User) fromModel(du model.User) {
	u.Admin = du.Admin
	u.Name = du.Name
}

func usersFromModel(mUsers []model.User) []User {
	users := make([]User, len(mUsers))
	for i, mu := range mUsers {
		user := new(User)
		user.fromModel(mu)
		user.TemplID = fmt.Sprintf("user-%v", i)

		users[i] = *user
	}

	return users
}
