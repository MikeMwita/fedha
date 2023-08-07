package repository

import (
	"context"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"
	"github.com/jackc/pgx/v5/pgtype"
)

type userRepo struct {
	db ports.UserStorage
}

func (u userRepo) CreateUser(ctx context.Context, arg *db.CreateUserParams) (*db.User, error) {
	createAcc, err := u.db.CreateUser(ctx, arg)
	if err != nil {
		return nil, err
	}
	return createAcc, nil
}

func (u userRepo) GetUser(ctx context.Context, username pgtype.Text) (*db.User, error) {
	res, err := u.db.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewUserRepo(db ports.UserStorage) ports.UserRepo {
	return &userRepo{
		db: db,
	}
}
