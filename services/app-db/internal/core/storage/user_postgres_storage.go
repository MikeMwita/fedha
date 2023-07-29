package storage

import (
	"context"
	"errors"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgtype"
)

var (
	ErrUserAlreadyExists = errors.New("user not found")
	ErrUserNotFound      = errors.New("user not found")
)

type userStorage struct {
	client *pgx.Conn
}

func (u *userStorage) CreateUser(ctx context.Context, arg *db.CreateUserParams) (*db.User, error) {
	//check if user already exists
	_, err := u.GetUser(ctx, arg.Username)
	if err == nil {
		return nil, ErrUserAlreadyExists
	}
	if err != nil && err != pgx.ErrNoRows {
		return nil, err
	}

	user := &db.User{}
	err = u.client.QueryRow(ctx, createUser,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.PasswordChangedAt,
		arg.CreatedAt,
	).Scan(
		&user.Username,
		&user.FullName,
		&user.Email,
		&user.PasswordChangedAt,
		&user.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return user, nil
}

const createUser = `-- name: CreateUser :one
INSERT INTO "User" (
    "Username",
    "Full_Name",
    "Email",
    "Password_Changed_At",
    "Created_At"
) VALUES (
             $1, $2, $3, $4, $5
         ) RETURNING "Username", "Full_Name", "Email", "Password_Changed_At", "Created_At"
`

func (u userStorage) GetUser(ctx context.Context, username pgtype.Text) (*db.User, error) {
	var user db.User
	err := u.client.QueryRow(ctx, getUser, username).Scan(
		&user.Username,
		&user.FullName,
		&user.Email,
		&user.PasswordChangedAt,
		&user.CreatedAt,
	)
	if err != nil {
		if err == pgx.ErrNoRows {

			return nil, ErrUserNotFound
		}

		return nil, err
	}
	return &user, nil
}

const getUser = `-- name: GetUser :one
SELECT "Username", "Full_Name", "Email", "Password_Changed_At", "Created_At" FROM "User"
WHERE "Username" = $1 LIMIT 1
`

func NewUserStorage(client *pgx.Conn) ports.UserStorage {
	return &userStorage{
		client: client,
	}
}
