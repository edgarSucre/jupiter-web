package domain

import "context"

type Repository interface {
	WithTx(fn func(Repository) error) error
	CreateUser(context.Context, CreateUserParams) (User, error)
	ListUsers(context.Context) ([]User, error)
	GetUserByEmail(context.Context, string) (User, error)
}
