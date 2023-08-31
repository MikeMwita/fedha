package adapters

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/gin-gonic/gin"
)

type AuthUseCase interface {
	Login(c *gin.Context, data dto.LoginRequest) (*dto.LoginResponseData, error)
	Register(c *gin.Context, data dto.RegisterRequest) (*dto.RegisterResponseData, error)
	RefreshToken(c *gin.Context, data dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
	UserLogout(c *gin.Context)
	VerifyAccessToken(token string) (interface{}, interface{})
	GetUserById(c *gin.Context, id string) (*entity.User, error)
}

type ExpenseStorageUseCase interface {
	CreateExpense(context.Context, *expense.CreateExpenseRequest) (*expense.CreateExpenseResponse, error)
	GetExpense(context.Context, *expense.GetExpenseRequest) (*expense.GetExpenseResponse, error)
	UpdateExpense(context.Context, *expense.UpdateExpenseRequest) (*expense.UpdateExpenseResponse, error)
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
