package db

import (
	"context"
	"database/sql"
	"errors"

	"github.com/edgarSucre/jw/domain"
)

type SqliteRepository struct {
	q *Queries
}

func NewSqliteRepository(db DBTX) *SqliteRepository {
	return &SqliteRepository{
		q: &Queries{
			db: db,
		},
	}
}

func (repo *SqliteRepository) CreateUser(
	ctx context.Context,
	dp domain.CreateUserParams,
) (domain.User, error) {
	params := new(CreateUserParams)
	params.copyDomain(dp)

	u, err := repo.q.CreateUser(ctx, *params)

	return u.domain(), err
}

func (repo *SqliteRepository) ListUsers(ctx context.Context) ([]domain.User, error) {
	users, err := repo.q.ListUsers(ctx)
	if err != nil {
		return []domain.User{}, err
	}

	return UsersToDomain(users), nil
}

func (repo *SqliteRepository) GetUserByUsername(
	ctx context.Context,
	userName string,
) (domain.User, error) {
	user, err := repo.q.GetUserByUserName(ctx, userName)
	if err != nil {
		user := domain.User{}
		if errors.Is(err, sql.ErrNoRows) {
			return user, domain.ErrNotFound
		}

		return user, err
	}

	return user.forAuth(), nil
}
