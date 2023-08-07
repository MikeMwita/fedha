package repository

import (
	"context"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"
)

type expenseTypeRepo struct {
	db ports.ExpenseTypeStorage
}

func (e expenseTypeRepo) CreateExpenseType(ctx context.Context, arg *db.CreateExpenseTypeParams) (*db.ExpenseType, error) {
	expenseType, err := e.db.CreateExpenseType(ctx, arg)
	if err != nil {
		return nil, err
	}

	return expenseType, nil
}

func (e expenseTypeRepo) DeleteExpenseType(ctx context.Context, expensetypeid int32) error {
	err := e.db.DeleteExpenseType(ctx, expensetypeid)
	if err != nil {
		return err
	}

	return nil
}

func (e expenseTypeRepo) GetExpenseTypeByID(ctx context.Context, expensetypeid int32) (*db.ExpenseType, error) {
	return e.db.GetExpenseTypeByID(ctx, expensetypeid)
}

func (e expenseTypeRepo) ListExpenseTypes(ctx context.Context) ([]db.ExpenseType, error) {
	expenseTypes, err := e.db.ListExpenseTypes(ctx)
	if err != nil {
		return nil, err
	}

	return expenseTypes, nil
}

func (e expenseTypeRepo) UpdateExpenseType(ctx context.Context, arg *db.UpdateExpenseTypeParams) (*db.ExpenseType, error) {
	res, err := e.db.UpdateExpenseType(ctx, arg)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func NewExpenseTypeRepo(db ports.ExpenseTypeStorage) ports.ExpenseTypeRepo {
	return &expenseTypeRepo{
		db: db,
	}
}
