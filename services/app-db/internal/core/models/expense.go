package models

import "github.com/jackc/pgx/v5/pgtype"

type Expense struct {
	ExpenseID     int32            `json:"ExpenseID"`
	ExpenseTypeID pgtype.Int4      `json:"ExpenseTypeID"`
	Amount        pgtype.Int4      `json:"Amount"`
	Description   pgtype.Text      `json:"Description"`
	CreatedAt     pgtype.Timestamp `json:"Created_At"`
}
