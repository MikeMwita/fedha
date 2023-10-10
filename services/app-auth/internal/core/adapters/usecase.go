package adapters

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/google/uuid"
)

//go:generate mockgen -source usecase.go -destination mock/usecase_mock.go -package mock

type AuthUseCase interface {
	Register(ctx context.Context, user *entity.User) (*entity.UserWithToken, error)
	Login(ctx context.Context, user *entity.User) (*entity.UserWithToken, error)
	Update(ctx context.Context, user *entity.User) (*entity.User, error)
	Delete(ctx context.Context, userID uuid.UUID) error
	GetByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
}

type ExpenseStorageUseCase interface {
	CreateExpense(context.Context, *expense.ExpenseRequest) (*expense.ExpenseResponse, error)
	GetExpense(context.Context, *expense.GetExpenseRequest) (*expense.ExpenseResponse, error)
	UpdateExpense(context.Context, *expense.UpdateExpenseRequest) (*expense.ExpenseResponse, error)
	DeleteExpense(context.Context, *expense.DeleteExpenseRequest) (*expense.DeleteExpenseResponse, error)
	// INCOMES
	CreateIncome(context.Context, *expense.CreateIncomeRequest) (*expense.CreateIncomeResponse, error)
	GetIncome(context.Context, *expense.GetIncomeRequest) (*expense.GetIncomeResponse, error)
	UpdateIncome(context.Context, *expense.UpdateIncomeRequest) (*expense.UpdateIncomeResponse, error)
	DeleteIncome(context.Context, *expense.DeleteIncomeRequest) (*expense.DeleteIncomeResponse, error)
	// BALANCES
	GetRemainingBalance(context.Context, *expense.RemainingBalanceRequest) (*expense.RemainingBalanceResponse, error)
	// MONTHLY SUMMARY
	GenerateMonthlySummary(context.Context, *expense.MonthlySummaryRequest) (*expense.MonthlySummaryResponse, error)
}
