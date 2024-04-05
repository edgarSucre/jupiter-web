package user

import (
	"context"

	"github.com/edgarSucre/jw/domain"
)

type Repository interface {
	CreateUser(context.Context, domain.CreateUserParams) (domain.User, error)
	DeleteUser(context.Context, int) error
	ListUsers(context.Context) ([]domain.User, error)
	WithTx(fn func(domain.Repository) error) error
}

type JupiterCLient interface {
	CreateUser(context.Context, string) (domain.JupyterUser, error)
}

type UseCase struct {
	jClient JupiterCLient
	repo    Repository
}

func NewUseCase(repo Repository, client JupiterCLient) *UseCase {
	return &UseCase{
		jClient: client,
		repo:    repo,
	}
}
