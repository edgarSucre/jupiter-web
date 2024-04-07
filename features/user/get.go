package user

import (
	"context"
	"fmt"
)

func (uc *UseCase) Get(ctx context.Context, userID int) (User, error) {
	user, err := uc.repo.GetUserByID(ctx, userID)
	if err != nil {
		return User{}, fmt.Errorf("%w, repo.GetUserByID", err)
	}

	response := new(User)
	response.FromDomain(user)

	return *response, nil
}
