package user

import (
	"context"
	"fmt"

	"github.com/edgarSucre/jw/domain"
	"golang.org/x/crypto/bcrypt"
)

func (uc *UseCase) Create(ctx context.Context, params CreateParams) (User, error) {
	user := new(User)

	hp, err := hashPassword(params.Password)
	if err != nil {
		return *user, fmt.Errorf("%w, hashPassword", err)
	}

	params.Password = hp

	var du domain.User

	err = uc.repo.WithTx(func(r domain.Repository) error {
		du, err = r.CreateUser(ctx, params.Domain())
		if err != nil {
			return fmt.Errorf("%w, repo.CreateUser", err)
		}

		_, err = uc.jClient.CreateUser(ctx, params.Email)
		if err != nil {
			return fmt.Errorf("%w, jClient.CreateUser", err)
		}

		return nil
	})

	if err != nil {
		return User{}, err
	}

	user.FromDomain(du)

	return *user, nil
}

func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}
