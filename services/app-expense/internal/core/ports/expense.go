package ports

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"google.golang.org/grpc"
)

type ExpenseRepository interface {
	CreateExpense(ctx context.Context, in *db.ExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error)
}

type ExpenseService interface {
	CreateExpense(ctx context.Context, in *db.ExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	GetExpense(ctx context.Context, in *db.GetExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	UpdateExpense(ctx context.Context, in *db.UpdateExpenseRequest, opts ...grpc.CallOption) (*db.ExpenseResponse, error)
	DeleteExpense(ctx context.Context, in *db.DeleteExpenseRequest, opts ...grpc.CallOption) (*db.DeleteExpenseResponse, error)
}
