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
	//Register(ctx context.Context, data dto.RegisterRequest) (*dto.RegisterResponseData, error)
	//Login(ctx context.Context, data dto.LoginRequest) (*dto.LoginResponseData, error)
	GetUserById(c *gin.Context, userId string) (string, error)
	RefreshToken(c *gin.Context, data dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
	UserLogout(c *gin.Context)
	VerifyAccessToken(token string) (interface{}, interface{})

	//grpc
	//RegisterUser(context.Context, *RegUserReq) (*RegUserRes, error)
	//UpdateUser(context.Context, *UpdateUserReq) (*UpdateUserRes, error)
	//GetPagedUsers(context.Context, *GetPagedUsersReq) (*GetPagedUsersRes, error)
	//GetUserByField(context.Context, *GetByfieldReq) (*GetByfieldRes, error)
	//GetUserByUsername(context.Context, *GetUserByUsernameRequest) (*RegUserRes, error)
	//GetUserByID(context.Context, *GetUserByIDRequest) (*RegUserRes, error)
	//SaveUser(context.Context, *SaveUserRequest) (*User, error)
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
