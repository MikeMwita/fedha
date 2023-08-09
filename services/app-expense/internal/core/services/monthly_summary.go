package services

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type MonthlyService struct {
	monthly ports.IncomeRepository
}

func (m MonthlyService) GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewMonthlyService(monthly ports.IncomeRepository) ports.MonthlyService {
	return &MonthlyService{
		monthly: monthly,
	}
}
