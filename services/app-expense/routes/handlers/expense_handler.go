package handlers

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"google.golang.org/grpc"
)

func (h *Handler) CreateExpense(ctx context.Context, in *expense.CreateExpenseRequest, opts ...grpc.CallOption) (*expense.CreateExpenseResponse, error) {
	expenseRequest := &db.CreateExpenseRequest{
		Expense: &db.Expense{
			ExpenseId: in.GetExpense().GetExpenseId(),
			Title:     in.GetExpense().GetTitle(),
			Amount:    in.GetExpense().GetAmount(),
			Category:  in.GetExpense().GetCategory(),
			Date:      in.GetExpense().GetDate(),
		},
	}
	res, err := h.expenseService.CreateExpense(ctx, expenseRequest)
	if err != nil {
		return nil, err
	}
	expenseResult := &expense.CreateExpenseResponse{
		Expense: &expense.Expense{
			ExpenseId: res.Expense.ExpenseId,
			Title:     res.Expense.Title,
			Amount:    res.Expense.Amount,
			Category:  res.Expense.Category,
			Date:      res.Expense.Date,
		},
	}
	return expenseResult, nil
}

func (h *Handler) GetExpense(ctx context.Context, in *expense.GetExpenseRequest, opts ...grpc.CallOption) (*expense.GetExpenseResponse, error) {

	expenseId := in.GetExpenseId()

	expenseRequest := &db.GetExpenseRequest{
		ExpenseId: expenseId,
	}
	expenseResponse, err := h.expenseService.GetExpense(ctx, expenseRequest)
	if err != nil {
		return nil, err
	}
	getExpenseResponse := &expense.GetExpenseResponse{
		Expense: &expense.Expense{
			ExpenseId: expenseResponse.Expense.ExpenseId,
			Title:     expenseResponse.Expense.Title,
			Amount:    expenseResponse.Expense.Amount,
			Category:  expenseResponse.Expense.Category,
			Date:      expenseResponse.Expense.Date,
		},
	}

	return getExpenseResponse, nil
}
func (h *Handler) UpdateExpense(ctx context.Context, in *expense.UpdateExpenseRequest, opts ...grpc.CallOption) (*expense.UpdateExpenseResponse, error) {
	updatedExpense := in.GetExpense()
	updatedExpenseRequest := &db.UpdateExpenseRequest{
		Expense: &db.Expense{
			ExpenseId: updatedExpense.GetExpenseId(),
			Title:     updatedExpense.GetTitle(),
			Amount:    updatedExpense.GetAmount(),
			Category:  updatedExpense.GetCategory(),
			Date:      updatedExpense.GetDate(),
		},
	}
	updatedExpenseResponse, err := h.expenseService.UpdateExpense(ctx, updatedExpenseRequest)
	if err != nil {
		return nil, err
	}
	updateExpenseResponse := &expense.UpdateExpenseResponse{
		Expense: &expense.Expense{
			ExpenseId: updatedExpenseResponse.Expense.ExpenseId,
			Title:     updatedExpenseResponse.Expense.Title,
			Amount:    updatedExpenseResponse.Expense.Amount,
			Category:  updatedExpenseResponse.Expense.Category,
			Date:      updatedExpenseResponse.Expense.Date,
		},
	}

	return updateExpenseResponse, nil
}
func (h *Handler) DeleteExpense(ctx context.Context, in *expense.DeleteExpenseRequest, opts ...grpc.CallOption) (*expense.DeleteExpenseResponse, error) {
	expenseId := in.GetExpenseId()
	deleteExpenseRequest := &db.DeleteExpenseRequest{
		ExpenseId: expenseId,
	}
	_, err := h.expenseService.DeleteExpense(ctx, deleteExpenseRequest)
	if err != nil {
		return nil, err
	}
	deleteExpenseResponse := &expense.DeleteExpenseResponse{
		Success: true,
	}

	return deleteExpenseResponse, nil
}
