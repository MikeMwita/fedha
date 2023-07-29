package ports

import (
	"context"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
)

type ExpenseTypeStorage interface {
	CreateExpenseType(ctx context.Context, arg *db.CreateExpenseTypeParams) (*db.ExpenseType, error)
	DeleteExpenseType(ctx context.Context, expensetypeid int32) error
	GetExpenseTypeByID(ctx context.Context, expensetypeid int32) (*db.ExpenseType, error)
	ListExpenseTypes(ctx context.Context) ([]db.ExpenseType, error)
	UpdateExpenseType(ctx context.Context, arg *db.UpdateExpenseTypeParams) (*db.ExpenseType, error)
}
type ExpenseTypeRepo interface {
	CreateExpenseType(ctx context.Context, arg *db.CreateExpenseTypeParams) (*db.ExpenseType, error)
	DeleteExpenseType(ctx context.Context, expensetypeid int32) error
	GetExpenseTypeByID(ctx context.Context, expensetypeid int32) (*db.ExpenseType, error)
	ListExpenseTypes(ctx context.Context) ([]db.ExpenseType, error)
	UpdateExpenseType(ctx context.Context, arg *db.UpdateExpenseTypeParams) (*db.ExpenseType, error)
}
