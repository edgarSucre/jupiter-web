package ucase

import (
	"context"
	"fmt"

	"github.com/edgarSucre/jw/features/user/model"
)

func (uc *UseCase) List(ctx context.Context) ([]model.User, error) {
	dUsers, err := uc.repo.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w, repo.ListUsers", err)
	}

	return model.UsersFromDomain(dUsers), nil
}
