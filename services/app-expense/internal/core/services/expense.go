package services

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type ExpenseService struct {
	expenseRepository ports.ExpenseRepository
}

func (e ExpenseService) CreateExpense(ctx context.Context, in *db.ExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseService) GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseService) UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (e ExpenseService) DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewExpenseService(expenseRepository ports.ExpenseRepository) ports.ExpenseService {
	return &ExpenseService{
		expenseRepository: expenseRepository,
	}
}
