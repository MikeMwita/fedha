package ports

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"google.golang.org/grpc"
)

type DbStorage interface {

	//HealthCheck(ctx context.Context, in *db.DefaultRequest) (*db.HealthResponse, error)

	CreateExpense(ctx context.Context, in *db.ExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error)
	// INCOMES
	CreateIncome(ctx context.Context, in *db.CreateIncomeRequest, opts ...grpc.CallOption) (*db.CreateIncomeResponse, error)
	GetIncome(ctx context.Context, in *db.GetIncomeRequest, opts ...grpc.CallOption) (*db.GetIncomeResponse, error)
	UpdateIncome(ctx context.Context, in *db.UpdateIncomeRequest, opts ...grpc.CallOption) (*db.UpdateIncomeResponse, error)
	DeleteIncome(ctx context.Context, in *db.DeleteIncomeRequest, opts ...grpc.CallOption) (*db.DeleteIncomeResponse, error)
	// BALANCES
	GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error)
	// MONTHLY SUMMARY
	GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error)
}
