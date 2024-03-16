package user

import (
	"context"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/user/model"
	"github.com/edgarSucre/jw/features/user/view"
)

type UseCase interface {
	Create(context.Context, model.CreateParams) (model.User, error)
	List(context.Context) ([]model.User, error)
}

type Handler struct {
	useCase UseCase
}

func NewHandler(uc UseCase) *Handler {
	return &Handler{
		useCase: uc,
	}
}

func (h *Handler) IndexCmp(ctx context.Context) templ.Component {
	return view.IndexCmp()
}

func (h *Handler) IndexFull(ctx context.Context) templ.Component {
	return view.IndexFull()
}
