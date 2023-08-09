package repositories

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type BalanceRepository struct {
	dbStorage ports.DbStorage
}

func (b BalanceRepository) GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error) {

	monthlySummary, err := b.dbStorage.GenerateMonthlySummary(ctx, in)
	if err != nil {
		return nil, err
	}
	return monthlySummary, nil

}

func (b BalanceRepository) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	remainingBalance, err := b.dbStorage.GetRemainingBalance(ctx, in)
	if err != nil {
		return nil, err
	}
	return remainingBalance, nil

}

func NewBalanceRepository(dbStorage ports.DbStorage) ports.BalanceRepo {

	return &MonthlyRepo{
		dbStorage: dbStorage,
	}
}
