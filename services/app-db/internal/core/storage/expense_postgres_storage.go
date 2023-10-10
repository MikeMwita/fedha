package storage

import (
	"context"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"
	"github.com/jackc/pgx/v5"
)

type expenseStorage struct {
	client *pgx.Conn
}

const createExpense = `-- name: CreateExpense :one
INSERT INTO "Expense" (
    "ExpenseID", "ExpenseTypeID", "Amount", "Description", "Created_At"
) VALUES (
             $1, $2, $3, $4, $5
         )
    RETURNING "ExpenseID", "ExpenseTypeID", "Amount", "Description", "Created_At"
`

func (e expenseStorage) CreateExpense(ctx context.Context, arg *db.CreateExpenseParams) (*db.Expense, error) {
	var expense *db.Expense
	err := e.client.QueryRow(ctx, createExpense,
		arg.ExpenseID,
		arg.ExpenseTypeID,
		arg.Amount,
		arg.Description,
		arg.CreatedAt,
	).Scan(
		&expense.ExpenseID,
		&expense.ExpenseTypeID,
		&expense.Amount,
		&expense.Description,
		&expense.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return expense, nil
}

const deleteExpense = `-- name: DeleteExpense :exec
DELETE FROM "Expense"
WHERE "ExpenseID" = $1
`

func (e expenseStorage) DeleteExpense(ctx context.Context, expenseid int32) error {
	_, err := e.client.Exec(ctx, deleteExpense, expenseid)
	return err
}

const listExpenses = `-- name: ListExpenses :many
SELECT "ExpenseID", "ExpenseTypeID", "Amount", "Description", "Created_At" FROM "Expense"
ORDER BY "Created_At" DESC
`

func (e expenseStorage) ListExpenses(ctx context.Context) ([]db.Expense, error) {
	rows, err := e.client.Query(ctx, listExpenses)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []db.Expense{}
	for rows.Next() {
		var i db.Expense
		if err := rows.Scan(
			&i.ExpenseID,
			&i.ExpenseTypeID,
			&i.Amount,
			&i.Description,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateExpense = `-- name: UpdateExpense :one
UPDATE "Expense"
SET "ExpenseTypeID" = $2, "Amount" = $3, "Description" = $4, "Created_At" = $5
WHERE "ExpenseID" = $1
    RETURNING "ExpenseID", "ExpenseTypeID", "Amount", "Description", "Created_At"
`

func (e expenseStorage) UpdateExpense(ctx context.Context, arg *db.UpdateExpenseParams) (*db.Expense, error) {
	row := e.client.QueryRow(ctx, updateExpense,
		arg.ExpenseID,
		arg.ExpenseTypeID,
		arg.Amount,
		arg.Description,
		arg.CreatedAt,
	)
	var updatedExpense db.Expense // Use db.Expense here
	err := row.Scan(
		&updatedExpense.ExpenseID,
		&updatedExpense.ExpenseTypeID,
		&updatedExpense.Amount,
		&updatedExpense.Description,
		&updatedExpense.CreatedAt,
	)
	return &updatedExpense, err
}

const getExpenseByID = `-- name: GetExpenseByID :one
SELECT "ExpenseID", "ExpenseTypeID", "Amount", "Description", "Created_At" FROM "Expense"
WHERE "ExpenseID" = $1
    LIMIT 1
`

func (e expenseStorage) GetExpenseByID(ctx context.Context, expenseid int32) (*db.Expense, error) {
	row := e.client.QueryRow(ctx, getExpenseByID, expenseid)
	var expense db.Expense
	err := row.Scan(
		&expense.ExpenseID,
		&expense.ExpenseTypeID,
		&expense.Amount,
		&expense.Description,
		&expense.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &expense, nil
}

func NewExpenseStorage(client *pgx.Conn) ports.ExpenseStorage {
	return &expenseStorage{
		client: client,
	}
}
