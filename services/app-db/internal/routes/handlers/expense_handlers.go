package handlers

import (
	"context"
	"errors"
	db "github.com/MikeMwita/fedha.git/services/app-db/db/generated"
)

var (
	ErrEmptyID = errors.New("Empty Expense ID")
)

func (h *Handler) CreateExpense(ctx context.Context, arg *db.CreateExpenseParams) (*db.Expense, error) {

	if arg == nil {
		return nil, ErrEmptyRequest
	}

	expense := &db.CreateExpenseParams{
		ExpenseID:     arg.ExpenseID,
		ExpenseTypeID: arg.ExpenseTypeID,
		Amount:        arg.Amount,
		Description:   arg.Description,
		CreatedAt:     arg.CreatedAt,
	}
	createdExpense, err := h.expenseRepo.CreateExpense(ctx, expense)
	if err != nil {
		return nil, err
	}
	return createdExpense, err
}

func (h *Handler) ListExpenses(ctx context.Context) ([]db.Expense, error) {
	expense, err := h.expenseRepo.ListExpenses(ctx)
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (h *Handler) GetExpenseByID(ctx context.Context, expenseid int32) (*db.Expense, error) {
	if expenseid == 0 {
		return nil, ErrEmptyID
	}
	expense, err := h.expenseRepo.GetExpenseByID(ctx, expenseid)
	if err != nil {
		return nil, err
	}
	return expense, nil
}

func (h *Handler) UpdateExpense(ctx context.Context, arg *db.UpdateExpenseParams) (*db.Expense, error) {
	if arg == nil {
		return nil, ErrEmptyRequest
	}
	expense := &db.UpdateExpenseParams{
		ExpenseID:     arg.ExpenseID,
		ExpenseTypeID: arg.ExpenseTypeID,
		Amount:        arg.Amount,
		Description:   arg.Description,
		CreatedAt:     arg.CreatedAt,
	}
	updatedExpense, err := h.expenseRepo.UpdateExpense(ctx, expense)
	if err != nil {
		return nil, err
	}
	return updatedExpense, nil
}

func (h *Handler) DeleteExpense(ctx context.Context, expenseid int32) error {
	if expenseid == 0 {
		return ErrEmptyRequest
	}
	err := h.expenseRepo.DeleteExpense(ctx, expenseid)
	if err != nil {
		return err
	}
	return nil
}
