package repository

import "github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"

type ExpenseRepositoryImpl struct {
	expenseStorage adapters.ExpenseStorage
}

func NewExpenseRepositoryImpl(expenseStorage adapters.ExpenseStorage) adapters.ExpenseStorage {
	return &ExpenseRepositoryImpl{expenseStorage: expenseStorage}
}
