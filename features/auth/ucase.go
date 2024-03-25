package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/edgarSucre/jw/domain"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	GetUserByUsername(context.Context, string) (domain.User, error)
}

type UseCase struct {
	repo Repository
}

func NewUseCase(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) Login(ctx context.Context, params LoginParams) (User, error) {
	user := new(User)

	du, err := uc.repo.GetUserByUsername(ctx, params.UserName)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			return *user, ErrBadCredentials
		}

		return *user, fmt.Errorf("%w: repo.GetUserByUsername(%s)", err, params.UserName)
	}

	if !passwordMatchHash(params.Password, du.Password) {
		return *user, ErrBadCredentials
	}

	user.FromDomain(du)

	return *user, nil
}

func passwordMatchHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
