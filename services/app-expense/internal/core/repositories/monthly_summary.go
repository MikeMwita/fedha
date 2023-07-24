package repositories

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type MonthlyRepo struct {
	dbStorage ports.DbStorage
}

func (m MonthlyRepo) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewMonthlyRepo(dbStorage ports.DbStorage) ports.MonthlyRepo {

	return &BalanceRepository{
		dbStorage: dbStorage,
	}
}
