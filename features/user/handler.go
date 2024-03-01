package user

import (
	"context"

	"github.com/edgarSucre/jw/features/user/model"
	"github.com/edgarSucre/jw/features/user/ucase"
)

type UseCase interface {
	Create(context.Context, model.CreateUserParams) (model.User, error)
	List(context.Context) ([]model.User, error)
}

type Handler struct {
	useCase UseCase
}

func NewHandler(ur ucase.Repository) *Handler {
	return &Handler{
		useCase: ucase.New(ur),
	}
}
