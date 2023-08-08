package services

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type BalanceService struct {
	balanceRepo ports.BalanceRepo
}

func (b BalanceService) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	totalIncome := b.balanceRepo.GetTotalIncome(ctx, in.Dates)
	totalExpense := b.balanceRepo.GetTotalExpense(ctx, in.Dates)
	remainingBalance := totalIncome.(float64) - totalExpense.(float64)

	return &db.RemainingBalanceResponse{
		RemainingBalance: remainingBalance,
	}, nil

}

func NewBalanceService(balanceRepo ports.BalanceRepo) ports.BalanceService {
	return &BalanceService{
		balanceRepo: balanceRepo,
	}
}
