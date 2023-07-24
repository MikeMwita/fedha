package storage

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type dbStorage struct {
	client db.DbServiceClient
}

func (d dbStorage) CreateExpense(ctx context.Context, in *db.ExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) CreateIncome(ctx context.Context, in *db.CreateIncomeRequest, opts ...grpc.CallOption) (*db.CreateIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) GetIncome(ctx context.Context, in *db.GetIncomeRequest, opts ...grpc.CallOption) (*db.GetIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) UpdateIncome(ctx context.Context, in *db.UpdateIncomeRequest, opts ...grpc.CallOption) (*db.UpdateIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) DeleteIncome(ctx context.Context, in *db.DeleteIncomeRequest, opts ...grpc.CallOption) (*db.DeleteIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewDbStorage(client db.DbServiceClient) ports.DbStorage {
	return &dbStorage{
		client: client,
	}
}
