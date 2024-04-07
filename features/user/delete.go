package user

import (
	"context"
	"fmt"

	"github.com/edgarSucre/jw/domain"
)

func (uc *UseCase) Delete(ctx context.Context, userID int) error {
	user, err := uc.repo.GetUserByID(ctx, userID)
	if err != nil {
		return fmt.Errorf("%w, repo.GetUserByID", err)
	}

	err = uc.repo.WithTx(func(r domain.Repository) error {
		if err := uc.repo.DeleteUser(ctx, userID); err != nil {
			return fmt.Errorf("%w, repo.DeleteUser", err)
		}

		if err := uc.jClient.DeleteUser(ctx, user.Email); err != nil {
			return fmt.Errorf("%w, jClient.DeleteUser", err)
		}

		return nil
	})

	if err != nil {
		return fmt.Errorf("%w, tx.DeleteUser", err)
	}

	return nil
}
