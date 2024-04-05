package user

import (
	"context"
	"fmt"
)

func (uc *UseCase) Delete(ctx context.Context, userID int) error {
	if err := uc.repo.DeleteUser(ctx, userID); err != nil {
		return fmt.Errorf("%w, repo.DeleteUser", err)
	}

	return nil
}
