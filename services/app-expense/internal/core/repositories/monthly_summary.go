package repositories

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type MonthlyRepo struct {
	dbStorage ports.DbStorage
}

func (m MonthlyRepo) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	remainingBalance, err := m.dbStorage.GetRemainingBalance(ctx, in)
	if err != nil {
		return nil, err
	}
	return remainingBalance, nil
}
func (m MonthlyRepo) GetTotalIncome(ctx context.Context, dates []*timestamppb.Timestamp) interface{} {
	//TODO implement me
	panic("implement me")
}

func (m MonthlyRepo) GetTotalExpense(ctx context.Context, dates []*timestamppb.Timestamp) interface{} {
	//TODO implement me
	panic("implement me")
}

func NewMonthlyRepo(dbStorage ports.DbStorage) ports.MonthlyRepo {

	return &BalanceRepository{
		dbStorage: dbStorage,
	}
}
