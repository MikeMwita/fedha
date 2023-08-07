package repository

import (
	"context"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"
)

type expenseRepo struct {
	db ports.ExpenseStorage
}

func (e expenseRepo) CreateExpense(ctx context.Context, arg *db.CreateExpenseParams) (*db.Expense, error) {
	return e.db.CreateExpense(ctx, arg)
}

func (e expenseRepo) DeleteExpense(ctx context.Context, expenseid int32) error {
	return e.db.DeleteExpense(ctx, expenseid)
}

func (e expenseRepo) ListExpenses(ctx context.Context) ([]db.Expense, error) {
	return e.db.ListExpenses(ctx)
}

func (e expenseRepo) UpdateExpense(ctx context.Context, arg *db.UpdateExpenseParams) (*db.Expense, error) {
	return e.db.UpdateExpense(ctx, arg)
}

func (e expenseRepo) GetExpenseByID(ctx context.Context, expenseid int32) (*db.Expense, error) {
	return e.db.GetExpenseByID(ctx, expenseid)
}

func NewExpenseRepo(db ports.ExpenseStorage) ports.ExpenseRepo {
	return &expenseRepo{
		db: db,
	}
}
