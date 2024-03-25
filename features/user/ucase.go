package user

import (
	"context"

	"github.com/edgarSucre/jw/domain"
)

type Repository interface {
	CreateUser(context.Context, domain.CreateUserParams) (domain.User, error)
	ListUsers(context.Context) ([]domain.User, error)
}

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}
