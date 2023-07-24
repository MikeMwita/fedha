package services

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type BalanceService struct {
	balanceRepo ports.BalanceRepo
}

func (b BalanceService) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewBalanceService(balanceRepo ports.BalanceRepo) ports.BalanceService {
	return &BalanceService{
		balanceRepo: balanceRepo,
	}
}
