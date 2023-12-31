// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Expense struct {
	ExpenseID     int32            `json:"ExpenseID"`
	ExpenseTypeID pgtype.Int4      `json:"ExpenseTypeID"`
	Amount        pgtype.Int4      `json:"Amount"`
	Description   pgtype.Text      `json:"Description"`
	CreatedAt     pgtype.Timestamp `json:"Created_At"`
}

type ExpenseType struct {
	ExpenseTypeID int32       `json:"ExpenseTypeID"`
	Name          pgtype.Text `json:"Name"`
	Description   pgtype.Text `json:"Description"`
}

type User struct {
	Username          pgtype.Text      `json:"Username"`
	FullName          pgtype.Text      `json:"Full_Name"`
	Email             pgtype.Text      `json:"Email"`
	PasswordChangedAt pgtype.Timestamp `json:"Password_Changed_At"`
	CreatedAt         pgtype.Timestamp `json:"Created_At"`
}
