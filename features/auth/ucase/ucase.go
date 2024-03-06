package ucase

import (
	"context"
	"fmt"

	"github.com/edgarSucre/jw/domain"
	"github.com/edgarSucre/jw/features/auth/model"
	"golang.org/x/crypto/bcrypt"
)

type Repository interface {
	GetUserByUsername(context.Context, string) (domain.User, error)
}

type UseCase struct {
	repo Repository
}

func New(repo Repository) *UseCase {
	return &UseCase{
		repo: repo,
	}
}

func (uc *UseCase) Login(ctx context.Context, params model.LoginParams) (model.User, error) {
	user := new(model.User)

	du, err := uc.repo.GetUserByUsername(ctx, params.UserName)
	if err != nil {
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
