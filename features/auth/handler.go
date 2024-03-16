package auth

import (
	"context"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/auth/model"
	"github.com/edgarSucre/jw/features/auth/view"
)

type useCase interface {
	Login(context.Context, model.LoginParams) (model.User, error)
}

type Handler struct {
	useCase useCase
}

func NewHandler(uc useCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (h *Handler) Login(ctx context.Context, err error) templ.Component {
	return view.Login(err)
}

func (h *Handler) Authenticate(ctx context.Context, params model.LoginParams) (model.User, error) {
	if err := params.Validate(); err != nil {
		return model.User{}, err
	}

	user, err := h.useCase.Login(ctx, params)

	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
