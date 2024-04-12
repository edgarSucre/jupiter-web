package db

import (
	"database/sql"

	"github.com/edgarSucre/jw/domain"
)

func (u User) domain() domain.User {
	return domain.User{
		Admin: u.Admin > 0,
		Email: u.Email,
		ID:    u.ID,
		Name:  u.Name,
	}
}

func (u User) forAuth() domain.User {
	return domain.User{
		Admin:    u.Admin > 0,
		Email:    u.Email,
		ID:       u.ID,
		Name:     u.Name,
		Password: u.Password,
	}
}

func UsersToDomain(users []User) []domain.User {
	dUsers := make([]domain.User, len(users))
	for i, user := range users {
		dUsers[i] = user.domain()
	}

	return dUsers
}

func (up *CreateUserParams) copyDomain(params domain.CreateUserParams) {
	if params.Admin {
		up.Admin = 1
	}

	up.Name = params.Name
	up.Password = params.Password
	up.Email = params.Email
}

func (up *UpdateUserParams) copyDomain(params domain.UpdateUserParams) {
	if params.Admin {
		up.Admin = 1
	}

	up.ID = int64(params.ID)
	up.Name = params.Name

	if params.Password != nil {
		up.Password = sql.NullString{
			String: *params.Password,
			Valid:  true,
		}
	}
}
