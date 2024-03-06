package db

import "github.com/edgarSucre/jw/domain"

func (u User) domain() domain.User {
	return domain.User{
		ID:       u.ID,
		Name:     u.Name,
		UserName: u.Username,
		Admin:    u.Admin > 0,
	}
}

func (u User) forAuth() domain.User {
	return domain.User{
		ID:       u.ID,
		Name:     u.Name,
		UserName: u.Username,
		Admin:    u.Admin > 0,
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
	up.Username = params.UserName
}
