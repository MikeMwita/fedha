package ports

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type BalanceRepo interface {
	GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error)
	GetTotalIncome(ctx context.Context, dates []*timestamppb.Timestamp) interface{}
	GetTotalExpense(ctx context.Context, dates []*timestamppb.Timestamp) interface{}
}

type BalanceService interface {
	GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error)
}
