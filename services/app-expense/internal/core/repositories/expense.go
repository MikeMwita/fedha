package repositories

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type ExpenseRepository struct {
	dbStorage ports.DbStorage
}

func (e ExpenseRepository) CreateIncome(ctx context.Context, in *db.CreateIncomeRequest, opts ...grpc.CallOption) (*db.CreateIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseRepository) GetIncome(ctx context.Context, in *db.GetIncomeRequest, opts ...grpc.CallOption) (*db.GetIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseRepository) UpdateIncome(ctx context.Context, in *db.UpdateIncomeRequest, opts ...grpc.CallOption) (*db.UpdateIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseRepository) DeleteIncome(ctx context.Context, in *db.DeleteIncomeRequest, opts ...grpc.CallOption) (*db.DeleteIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseRepository) CreateExpense(ctx context.Context, in *db.ExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseRepository) GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseRepository) UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseRepository) DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewExpenseRepository(dbStorage ports.DbStorage) ports.ExpenseRepository {

	return &ExpenseRepository{
		dbStorage: dbStorage,
	}
}
