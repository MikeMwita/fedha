package ports

import (
	"context"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
	"github.com/jackc/pgx/v5/pgtype"
)

type DatabaseRepository interface {
	CreateExpense(ctx context.Context, arg *db.CreateExpenseParams) (*db.Expense, error)
	CreateExpenseType(ctx context.Context, arg *db.CreateExpenseTypeParams) (*db.ExpenseType, error)
	DeleteExpense(ctx context.Context, expenseid int32) error
	DeleteExpenseType(ctx context.Context, expensetypeid int32) error
	GetExpenseByID(ctx context.Context, expenseid int32) (*db.Expense, error)
	GetExpenseTypeByID(ctx context.Context, expensetypeid int32) (*db.ExpenseType, error)
	ListExpenseTypes(ctx context.Context) ([]db.ExpenseType, error)
	ListExpenses(ctx context.Context) ([]db.Expense, error)
	UpdateExpense(ctx context.Context, arg *db.UpdateExpenseParams) (*db.Expense, error)
	UpdateExpenseType(ctx context.Context, arg *db.UpdateExpenseTypeParams) (*db.ExpenseType, error)
}

type UserStorage interface {
	CreateUser(ctx context.Context, arg *db.CreateUserParams) (*db.User, error)
	GetUser(ctx context.Context, username pgtype.Text) (*db.User, error)
}

type UserRepo interface {
	CreateUser(ctx context.Context, arg *db.CreateUserParams) (*db.User, error)
	GetUser(ctx context.Context, username pgtype.Text) (*db.User, error)
}

type UserService interface {
	CreateUser(ctx context.Context, arg *db.CreateUserParams) (*db.User, error)
	GetUser(ctx context.Context, username pgtype.Text) (*db.User, error)
}
