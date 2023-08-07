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
	//create new income record
	income := &db.CreateIncomeRequest{
		Amount: in.Amount,
		Date:   in.Date,
	}
	// Save the income record to the database.
	_, err := e.dbStorage.CreateIncome(ctx, income)
	if err != nil {
		return nil, err
	}

	// Create a new `CreateIncomeResponse` object.
	createIncomeResponse := &db.CreateIncomeResponse{
		//IncomeId: incomeid,

	}

	return createIncomeResponse, nil
}

func (e ExpenseRepository) GetIncome(ctx context.Context, in *db.GetIncomeRequest, opts ...grpc.CallOption) (*db.GetIncomeResponse, error) {
	income, err := e.dbStorage.GetIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return income, nil
}

func (e ExpenseRepository) UpdateIncome(ctx context.Context, in *db.UpdateIncomeRequest, opts ...grpc.CallOption) (*db.UpdateIncomeResponse, error) {
	incomeUpdate, err := e.dbStorage.UpdateIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return incomeUpdate, nil
}

func (e ExpenseRepository) DeleteIncome(ctx context.Context, in *db.DeleteIncomeRequest, opts ...grpc.CallOption) (*db.DeleteIncomeResponse, error) {
	incomeDelete, err := e.dbStorage.DeleteIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return incomeDelete, nil
}

func (e ExpenseRepository) CreateExpense(ctx context.Context, in *db.ExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	// Create a new expense record.
	expense := &db.ExpenseRequest{
		Title:    in.Title,
		Amount:   in.Amount,
		Category: in.Category,
		Date:     in.Date,
	}

	// Save the expense record to the database.
	_, err := e.dbStorage.CreateExpense(ctx, expense)
	if err != nil {
		return nil, err
	}

	// Create a new `ExpenseResponse` object.

	expenseResponse := &db.ExpenseResponse{
		//ExpenseId: expenseId,
		Title:    expense.Title,
		Amount:   expense.Amount,
		Category: expense.Category,
		Date:     expense.Date,
	}

	return expenseResponse, nil
}

func (e ExpenseRepository) GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	// Get the expense record from the database.
	expense, err := e.dbStorage.GetExpense(ctx, in)
	if err != nil {
		return nil, err
	}

	// Create a new `ExpenseResponse` object.
	expenseResponse := &db.ExpenseResponse{
		ExpenseId: expense.ExpenseId,
		Title:     expense.Title,
		Amount:    expense.Amount,
		Category:  expense.Category,
		Date:      expense.Date,
	}

	return expenseResponse, nil
}

func (e ExpenseRepository) UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error) {
	expenseUpdate, err := e.dbStorage.UpdateExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return expenseUpdate, nil
}

func (e ExpenseRepository) DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error) {
	expenseDelete, err := e.dbStorage.DeleteExpense(ctx, in)
	if err != nil {
		return nil, err
	}
	return expenseDelete, nil
}

func NewExpenseRepository(dbStorage ports.DbStorage) ports.ExpenseRepository {

	return &ExpenseRepository{
		dbStorage: dbStorage,
	}
}
