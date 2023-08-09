package adapters

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
)

type ExpenseStorage interface {
	GetExpense(ctx context.Context, id string) (*expense.GetExpenseRequest, error)
}

type ExpenseRepository interface{}

type ExpenseService interface{}
