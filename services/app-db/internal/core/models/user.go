package models

import "github.com/jackc/pgx/v5/pgtype"

type User struct {
	Username          pgtype.Text      `json:"Username"`
	FullName          pgtype.Text      `json:"Full_Name"`
	Email             pgtype.Text      `json:"Email"`
	PasswordChangedAt pgtype.Timestamp `json:"Password_Changed_At"`
	CreatedAt         pgtype.Timestamp `json:"Created_At"`
}
