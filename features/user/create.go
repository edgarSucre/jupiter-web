package user

import (
	"context"
	"fmt"

	"github.com/edgarSucre/jw/features/user/model"
)

func (h *Handler) Create(ctx context.Context, params model.CreateParams) (model.User, error) {
	user, err := h.useCase.Create(ctx, params)
	if err != nil {
		return model.User{}, fmt.Errorf("%w, usecase.Create(%+v)", err, params)
	}

	return user, nil
}
