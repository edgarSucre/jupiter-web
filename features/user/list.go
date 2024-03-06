package user

import (
	"context"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/user/view"
)

func (h *Handler) List(ctx context.Context) templ.Component {
	mUsers, err := h.useCase.List(ctx)
	if err != nil {
		return view.ListErr()
	}

	return view.List(mUsers)
}
