package services

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"google.golang.org/grpc"
)

type IncomeService struct {
	incomeRepository ports.IncomeRepository
}

func (i IncomeService) CreateIncome(ctx context.Context, in *db.CreateIncomeRequest, opts ...grpc.CallOption) (*db.CreateIncomeResponse, error) {
	incomeId, err := i.incomeRepository.CreateIncome(ctx, in)
	if err != nil {
		return nil, err
	}

	return &db.CreateIncomeResponse{
		IncomeId: incomeId.String(),
	}, nil
}

func (i IncomeService) GetIncome(ctx context.Context, in *db.GetIncomeRequest, opts ...grpc.CallOption) (*db.GetIncomeResponse, error) {
	_, err := i.incomeRepository.GetIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return &db.GetIncomeResponse{
		//Income: income,
	}, nil
}

func (i IncomeService) UpdateIncome(ctx context.Context, in *db.UpdateIncomeRequest, opts ...grpc.CallOption) (*db.UpdateIncomeResponse, error) {
	_, err := i.incomeRepository.GetIncome(ctx, &db.GetIncomeRequest{IncomeId: in.IncomeId})
	if err != nil {
		return nil, err
	}
	_, err = i.incomeRepository.UpdateIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return &db.UpdateIncomeResponse{}, nil
}

func (i IncomeService) DeleteIncome(ctx context.Context, in *db.DeleteIncomeRequest, opts ...grpc.CallOption) (*db.DeleteIncomeResponse, error) {
	_, err := i.incomeRepository.GetIncome(ctx, &db.GetIncomeRequest{IncomeId: in.IncomeId})
	if err != nil {
		return nil, err
	}
	_, err = i.incomeRepository.DeleteIncome(ctx, in)
	if err != nil {
		return nil, err
	}
	return &db.DeleteIncomeResponse{}, nil
}

func NewIncomeService(incomeRepository ports.IncomeRepository) ports.IncomeService {
	return &IncomeService{incomeRepository: incomeRepository}
}
