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
	resExpense, err := d.client.CreateExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return resExpense, nil

}

func (d dbStorage) GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	res, err := d.client.GetExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	res, err := d.client.UpdateExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error) {
	res, err := d.client.DeleteExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) CreateIncome(ctx context.Context, in *db.CreateIncomeRequest, opts ...grpc.CallOption) (*db.CreateIncomeResponse, error) {
	res, err := d.client.CreateIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) GetIncome(ctx context.Context, in *db.GetIncomeRequest, opts ...grpc.CallOption) (*db.GetIncomeResponse, error) {
	res, err := d.client.GetIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) UpdateIncome(ctx context.Context, in *db.UpdateIncomeRequest, opts ...grpc.CallOption) (*db.UpdateIncomeResponse, error) {
	res, err := d.client.UpdateIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) DeleteIncome(ctx context.Context, in *db.DeleteIncomeRequest, opts ...grpc.CallOption) (*db.DeleteIncomeResponse, error) {
	res, err := d.client.DeleteIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	res, err := d.client.GetRemainingBalance(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (d dbStorage) GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error) {
	res, err := d.client.GenerateMonthlySummary(ctx, in)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func NewDbStorage(client db.DbServiceClient) ports.DbStorage {
	return &dbStorage{
		client: client,
	}
}
