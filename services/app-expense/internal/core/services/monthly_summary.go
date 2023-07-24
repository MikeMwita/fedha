package services

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type MonthlyRepo struct {
	monthly ports.IncomeRepository
}

func (m MonthlyRepo) GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewMonthlyRepo(monthly ports.IncomeRepository) ports.MonthlyService {
	return &MonthlyRepo{
		monthly: monthly,
	}
}
