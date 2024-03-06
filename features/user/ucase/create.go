package ucase

import (
	"context"
	"fmt"

	"github.com/edgarSucre/jw/features/user/model"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Create(ctx context.Context, params model.CreateParams) (model.User, error) {
	user := new(model.User)
	if err := params.Validate(); err != nil {
		return *user, fmt.Errorf("%w, params.Validate", err)
	}

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
