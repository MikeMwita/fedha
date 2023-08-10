package handlers

import (
	"context"
	"errors"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrEmptyRequest = errors.New("The request is empty")
)

func (h *Handler) CreateUser(ctx context.Context, arg *db.CreateUserParams) (*db.User, error) {
	if arg == nil {
		return nil, ErrEmptyRequest
	}
	user := &db.CreateUserParams{
		Username:          arg.Username,
		FullName:          arg.FullName,
		Email:             arg.Email,
		PasswordChangedAt: arg.PasswordChangedAt,
		CreatedAt:         arg.CreatedAt,
	}
	createdUser, err := h.userRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, ErrEmptyRequest
	}
	res := &db.User{
		Username:          createdUser.Username,
		FullName:          createdUser.FullName,
		Email:             createdUser.Email,
		PasswordChangedAt: createdUser.PasswordChangedAt,
		CreatedAt:         createdUser.CreatedAt,
	}

	return res, nil
}

func (h *Handler) GetUser(ctx context.Context, username pgtype.Text) (*db.User, error) {
	user, err := h.userRepo.GetUser(ctx, username)
	if err != nil {
		return nil, err
	}

	return user, nil
}
