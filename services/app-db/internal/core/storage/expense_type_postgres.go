package storage

import (
	"context"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"
	"github.com/jackc/pgx/v5"
)

type expensetypeStorage struct {
	client *pgx.Conn
}

const createExpenseType = `-- name: CreateExpenseType :one
INSERT INTO "ExpenseType" (
    "ExpenseTypeID", "Name", "Description"
) VALUES (
             $1, $2, $3
         )
    RETURNING "ExpenseTypeID", "Name", "Description"
`

func (e expensetypeStorage) CreateExpenseType(ctx context.Context, arg *db.CreateExpenseTypeParams) (*db.ExpenseType, error) {
	row := e.client.QueryRow(ctx, createExpenseType, arg.ExpenseTypeID, arg.Name, arg.Description)
	var expensetype db.ExpenseType // Use db.ExpenseType here
	err := row.Scan(&expensetype.ExpenseTypeID, &expensetype.Name, &expensetype.Description)
	if err != nil {
		return nil, err
	}
	return &expensetype, nil
}

const deleteExpenseType = `-- name: DeleteExpenseType :exec
DELETE FROM "ExpenseType"
WHERE "ExpenseTypeID" = $1
`

func (e expensetypeStorage) DeleteExpenseType(ctx context.Context, expensetypeid int32) error {
	_, err := e.client.Exec(ctx, deleteExpenseType, expensetypeid)
	return err
}

const getExpenseTypeByID = `-- name: GetExpenseTypeByID :one
SELECT "ExpenseTypeID", "Name", "Description" FROM "ExpenseType"
WHERE "ExpenseTypeID" = $1
    LIMIT 1
`

func (e expensetypeStorage) GetExpenseTypeByID(ctx context.Context, expensetypeid int32) (*db.ExpenseType, error) {
	row := e.client.QueryRow(ctx, getExpenseTypeByID, expensetypeid)
	var expensetype db.ExpenseType // Use db.ExpenseType here
	err := row.Scan(&expensetype.ExpenseTypeID, &expensetype.Name, &expensetype.Description)
	if err != nil {
		return nil, err
	}
	return &expensetype, nil
}

const listExpenseTypes = `-- name: ListExpenseTypes :many
SELECT "ExpenseTypeID", "Name", "Description" FROM "ExpenseType"
ORDER BY "Name"
`

func (e expensetypeStorage) ListExpenseTypes(ctx context.Context) ([]db.ExpenseType, error) {
	rows, err := e.client.Query(ctx, listExpenseTypes)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	expensetypes := []db.ExpenseType{}
	for rows.Next() {
		var expensetype db.ExpenseType // Use db.ExpenseType here
		if err := rows.Scan(&expensetype.ExpenseTypeID, &expensetype.Name, &expensetype.Description); err != nil {
			return nil, err
		}
		expensetypes = append(expensetypes, expensetype)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return expensetypes, nil
}

const updateExpenseType = `-- name: UpdateExpenseType :one
UPDATE "ExpenseType"
SET "Name" = $2, "Description" = $3
WHERE "ExpenseTypeID" = $1
    RETURNING "ExpenseTypeID", "Name", "Description"
`

func (e expensetypeStorage) UpdateExpenseType(ctx context.Context, arg *db.UpdateExpenseTypeParams) (*db.ExpenseType, error) {
	row := e.client.QueryRow(ctx, updateExpenseType, arg.ExpenseTypeID, arg.Name, arg.Description)
	var expensetype db.ExpenseType
	err := row.Scan(&expensetype.ExpenseTypeID, &expensetype.Name, &expensetype.Description)
	return &expensetype, err
}

func NewExpenseTypeStorage(client *pgx.Conn) ports.ExpenseTypeStorage {
	return &expensetypeStorage{client: client}
}
