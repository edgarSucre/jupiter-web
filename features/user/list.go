package user

import (
	"context"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/user/view"
)

// ListUsers shows the list of users
func (h *Handler) ListUsers(ctx context.Context) templ.Component {
	mUsers, err := h.useCase.List(ctx)
	if err != nil {
		return view.ListErr()
	}

	return view.List(mUsers)
}
