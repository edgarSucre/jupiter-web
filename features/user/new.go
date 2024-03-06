package user

import (
	"context"

	"github.com/a-h/templ"
	"github.com/edgarSucre/jw/features/user/view"
)

// New shows the new user form
func (h *Handler) New(ctx context.Context) templ.Component {
	return view.NewUser()
}
