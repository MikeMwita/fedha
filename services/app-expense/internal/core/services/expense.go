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
	expenseId, err := e.expenseRepository.CreateExpense(ctx, in)
	if err != nil {
		return nil, err
	}

	return &db.ExpenseResponse{
		ExpenseId: expenseId.String(),
		Title:     in.Title,
		Amount:    in.Amount,
		Category:  in.Category,
		Date:      in.Date,
	}, nil
}

func (e ExpenseService) GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {

	expense, err := e.expenseRepository.GetExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return &db.ExpenseResponse{
		ExpenseId: expense.ExpenseId,
		Title:     expense.Title,
		Amount:    expense.Amount,
		Category:  expense.Category,
		Date:      expense.Date,
	}, nil
}

func (e ExpenseService) UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	return e.expenseRepository.UpdateExpense(ctx, in)
}

func (e ExpenseService) DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error) {
	return e.expenseRepository.DeleteExpense(ctx, in)
}

func NewExpenseService(expenseRepository ports.ExpenseRepository) ports.ExpenseService {
	return &ExpenseService{
		expenseRepository: expenseRepository,
	}
}
