package ucase

import (
	"context"
	"fmt"

	"github.com/edgarSucre/jw/features/user/model"
)

func (uc *UseCase) Create(ctx context.Context, params model.CreateUserParams) (model.User, error) {
	user := new(model.User)
	if err := params.Validate(); err != nil {
		return *user, fmt.Errorf("%w, params.Validate", err)
	}

	du, err := uc.repo.CreateUser(ctx, params.Domain())
	if err != nil {
		return *user, fmt.Errorf("%w, repo.CreateUser", err)
	}

	user.FromDomain(du)

	return *user, nil
}
