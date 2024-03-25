package user

import (
	"context"
	"fmt"
)

func (uc *UseCase) List(ctx context.Context) ([]User, error) {
	dUsers, err := uc.repo.ListUsers(ctx)
	if err != nil {
		return nil, fmt.Errorf("%w, repo.ListUsers", err)
	}

	return usersFromDomain(dUsers), nil
}
