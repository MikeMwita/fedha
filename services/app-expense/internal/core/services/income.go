package services

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type IncomeService struct {
	incomeRepository ports.IncomeRepository
}

func (i IncomeService) CreateIncome(ctx context.Context, in *db.CreateIncomeRequest, opts ...grpc.CallOption) (*db.CreateIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i IncomeService) GetIncome(ctx context.Context, in *db.GetIncomeRequest, opts ...grpc.CallOption) (*db.GetIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i IncomeService) UpdateIncome(ctx context.Context, in *db.UpdateIncomeRequest, opts ...grpc.CallOption) (*db.UpdateIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (i IncomeService) DeleteIncome(ctx context.Context, in *db.DeleteIncomeRequest, opts ...grpc.CallOption) (*db.DeleteIncomeResponse, error) {
	//TODO implement me
	panic("implement me")
}

func NewIncomeService(incomeRepository ports.IncomeRepository) ports.IncomeService {
	return &IncomeService{incomeRepository: incomeRepository}
}
