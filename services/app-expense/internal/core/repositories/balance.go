package repositories

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type BalanceRepository struct {
	dbStorage ports.DbStorage
}

func (b BalanceRepository) GenerateMonthlySummary(ctx context.Context, in *db.MonthlySummaryRequest, opts ...grpc.CallOption) (*db.MonthlySummaryResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (b BalanceRepository) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewBalanceRepository(dbStorage ports.DbStorage) ports.BalanceRepo {

	return &MonthlyRepo{
		dbStorage: dbStorage,
	}
}
