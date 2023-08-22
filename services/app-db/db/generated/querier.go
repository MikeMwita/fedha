// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1

package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

type Querier interface {
	CreateExpense(ctx context.Context, arg CreateExpenseParams) (Expense, error)
	CreateExpenseType(ctx context.Context, arg CreateExpenseTypeParams) (ExpenseType, error)
	CreateUser(ctx context.Context, arg CreateUserParams) (User, error)
	DeleteExpense(ctx context.Context, expenseid int32) error
	DeleteExpenseType(ctx context.Context, expensetypeid int32) error
	GetExpenseByID(ctx context.Context, expenseid int32) (Expense, error)
	GetExpenseTypeByID(ctx context.Context, expensetypeid int32) (ExpenseType, error)
	GetUser(ctx context.Context, username pgtype.Text) (User, error)
	ListExpenseTypes(ctx context.Context) ([]ExpenseType, error)
	ListExpenses(ctx context.Context) ([]Expense, error)
	UpdateExpense(ctx context.Context, arg UpdateExpenseParams) (Expense, error)
	UpdateExpenseType(ctx context.Context, arg UpdateExpenseTypeParams) (ExpenseType, error)
}

var _ Querier = (*Queries)(nil)