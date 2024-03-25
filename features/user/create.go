package user

import (
	"context"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Create(ctx context.Context, params CreateParams) (User, error) {
	user := new(User)

	hp, err := hashPassword(params.Password)
	if err != nil {
		return *user, fmt.Errorf("%w, hashPassword", err)
	}

	params.Password = hp

	du, err := uc.repo.CreateUser(ctx, params.Domain())
	if err != nil {
		return *user, fmt.Errorf("%w, repo.CreateUser", err)
	}

	user.FromDomain(du)

	return *user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
