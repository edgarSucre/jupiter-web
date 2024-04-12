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

func (repo *SqliteRepository) WithTx(fn func(domain.Repository) error) error {
	// TODO: handle not ok
	db, _ := repo.q.db.(*sql.DB)

	// handle error
	tx, _ := db.Begin()
	defer tx.Rollback()

	if err := fn(NewSqliteRepository(tx)); err != nil {
		return err
	}

	tx.Commit()

	return nil
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

func (repo *SqliteRepository) DeleteUser(ctx context.Context, userID int) error {
	return repo.q.DeleteUser(ctx, int64(userID))
}

func (repo *SqliteRepository) ListUsers(ctx context.Context) ([]domain.User, error) {
	users, err := repo.q.ListUsers(ctx)
	if err != nil {
		return []domain.User{}, err
	}

	return UsersToDomain(users), nil
}

func (repo *SqliteRepository) GetUserByEmail(
	ctx context.Context,
	email string,
) (domain.User, error) {
	user, err := repo.q.GetUserByEmail(ctx, email)
	if err != nil {
		user := domain.User{}
		if errors.Is(err, sql.ErrNoRows) {
			return user, domain.ErrNotFound
		}

		return user, err
	}

	return user.forAuth(), nil
}

func (repo *SqliteRepository) GetUserByID(
	ctx context.Context,
	id int,
) (domain.User, error) {
	user, err := repo.q.GetUserByID(ctx, int64(id))
	if err != nil {
		user := domain.User{}
		if errors.Is(err, sql.ErrNoRows) {
			return user, domain.ErrNotFound
		}

		return user, err
	}

	return user.domain(), nil
}

func (repo *SqliteRepository) UpdateUser(
	ctx context.Context,
	dp domain.UpdateUserParams,
) error {
	params := new(UpdateUserParams)
	params.copyDomain(dp)

	return repo.q.UpdateUser(ctx, *params)
}
