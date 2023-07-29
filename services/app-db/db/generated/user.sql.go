// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: user.sql

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

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

type CreateUserParams struct {
	Username          pgtype.Text      `json:"Username"`
	FullName          pgtype.Text      `json:"Full_Name"`
	Email             pgtype.Text      `json:"Email"`
	PasswordChangedAt pgtype.Timestamp `json:"Password_Changed_At"`
	CreatedAt         pgtype.Timestamp `json:"Created_At"`
}

func (q *Queries) CreateUser(ctx context.Context, arg CreateUserParams) (User, error) {
	row := q.db.QueryRow(ctx, createUser,
		arg.Username,
		arg.FullName,
		arg.Email,
		arg.PasswordChangedAt,
		arg.CreatedAt,
	)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}

const getUser = `-- name: GetUser :one
SELECT "Username", "Full_Name", "Email", "Password_Changed_At", "Created_At" FROM "User"
WHERE "Username" = $1 LIMIT 1
`

func (q *Queries) GetUser(ctx context.Context, username pgtype.Text) (User, error) {
	row := q.db.QueryRow(ctx, getUser, username)
	var i User
	err := row.Scan(
		&i.Username,
		&i.FullName,
		&i.Email,
		&i.PasswordChangedAt,
		&i.CreatedAt,
	)
	return i, err
}
