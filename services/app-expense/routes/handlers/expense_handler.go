package handlers

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"google.golang.org/grpc"
)

func (h *Handler) CreateExpense(ctx context.Context, in *expense.ExpenseRequest, opts ...grpc.CallOption) (*expense.ExpenseResponse, error) {

}

func (h *Handler) GetExpense(ctx context.Context, in *expense.GetExpenseRequest, opts ...grpc.CallOption) (*expense.ExpenseResponse, error) {
}
func (h *Handler) UpdateExpense(ctx context.Context, in *expense.UpdateExpenseRequest, opts ...grpc.CallOption) (*expense.ExpenseResponse, error) {
}
func (h *Handler) DeleteExpense(ctx context.Context, in *expense.DeleteExpenseRequest, opts ...grpc.CallOption) (*expense.DeleteExpenseResponse, error) {
}
