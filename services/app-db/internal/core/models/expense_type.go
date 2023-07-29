package models

import "github.com/jackc/pgx/v5/pgtype"

type ExpenseType struct {
	ExpenseTypeID int32       `json:"ExpenseTypeID"`
	Name          pgtype.Text `json:"Name"`
	Description   pgtype.Text `json:"Description"`
}
