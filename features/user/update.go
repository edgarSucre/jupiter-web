package user

import (
	"context"
	"fmt"
)

func (uc *UseCase) Update(ctx context.Context, params UpdateParams) error {
	_, err := uc.repo.GetUserByID(ctx, params.ID)
	if err != nil {
		return fmt.Errorf("%w, repo.GetUserByID", err)
	}

	if params.Password != "" {
		hp, err := hashPassword(params.Password)
		if err != nil {
			return fmt.Errorf("%w, hashPassword", err)
		}

		params.Password = hp
	}

	if err := uc.repo.UpdateUser(ctx, params.Domain()); err != nil {
		return fmt.Errorf("%w, repo.UpdateUser", err)
	}

	return nil
}
