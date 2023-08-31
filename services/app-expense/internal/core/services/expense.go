package services

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type ExpenseService struct {
	expenseRepository ports.ExpenseRepository
}

func (e ExpenseService) CreateExpense(ctx context.Context, in *db.CreateExpenseRequest, opts ...grpc.CallOption) (*db.CreateExpenseResponse, error) {
	expenseId, err := e.expenseRepository.CreateExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return &db.CreateExpenseResponse{
		Expense: &db.Expense{
			ExpenseId: expenseId.String(),
			Title:     in.Expense.Title,
			Amount:    in.Expense.Amount,
			Category:  in.Expense.Category,
			Date:      in.Expense.Date,
		},
	}, nil
}

func (e ExpenseService) GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.GetExpenseResponse, error) {

	expense, err := e.expenseRepository.GetExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return &db.GetExpenseResponse{
		Expense: &db.Expense{
			ExpenseId: expense.Expense.ExpenseId,
			Title:     expense.Expense.Title,
			Amount:    expense.Expense.Amount,
			Category:  expense.Expense.Category,
			Date:      expense.Expense.Date,
		},
	}, nil
}

func (e ExpenseService) UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.UpdateExpenseResponse, error) {
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
