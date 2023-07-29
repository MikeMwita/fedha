package ports

import (
	"context"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
)

type ExpenseStorage interface {
	CreateExpense(ctx context.Context, arg *db.CreateExpenseParams) (*db.Expense, error)
	DeleteExpense(ctx context.Context, expenseid int32) error
	ListExpenses(ctx context.Context) ([]db.Expense, error)
	UpdateExpense(ctx context.Context, arg *db.UpdateExpenseParams) (*db.Expense, error)
	GetExpenseByID(ctx context.Context, expenseid int32) (*db.Expense, error)
}

type ExpenseRepo interface {
	CreateExpense(ctx context.Context, arg *db.CreateExpenseParams) (*db.Expense, error)
	DeleteExpense(ctx context.Context, expenseid int32) error
	ListExpenses(ctx context.Context) ([]db.Expense, error)
	UpdateExpense(ctx context.Context, arg *db.UpdateExpenseParams) (*db.Expense, error)
	GetExpenseByID(ctx context.Context, expenseid int32) (*db.Expense, error)
}
