package ports

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"google.golang.org/grpc"
)

type MonthlyRepo interface {
	GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error)
}

type MonthlyService interface {
	GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error)
}
