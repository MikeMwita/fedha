package repositories

import "github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"

type IncomeRepository struct {
	dbStorage ports.DbStorage
}

func NewIncomeRepository(dbStorage ports.DbStorage) ports.IncomeRepository {

	return &ExpenseRepository{
		dbStorage: dbStorage,
	}
}
