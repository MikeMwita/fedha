package repository

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
)

type ExpenseRepositoryImpl struct {
	expenseStorage adapters.ExpenseStorage
}

func (e ExpenseRepositoryImpl) GetExpense(ctx context.Context, id string) (*expense.GetExpenseRequest, error) {
	//TODO implement me
	panic("implement me")
}

func NewExpenseRepositoryImpl(expenseStorage adapters.ExpenseStorage) adapters.ExpenseStorage {
	return &ExpenseRepositoryImpl{expenseStorage: expenseStorage}
}
