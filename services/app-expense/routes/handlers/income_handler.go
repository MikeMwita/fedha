package handlers

import (
	"context"
	"errors"
	"fmt"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"google.golang.org/grpc"
)

var (
	ErrEmptyRequest       = errors.New("empty request")
	ErrAmountLessThanZero = errors.New("amount less than zero")
	ErrEmptyIncomeID      = errors.New("empty income ID")
	ErrDeleteIncome       = errors.New("error deleting income")
)

func (h *Handler) CreateIncome(ctx context.Context, request *expense.CreateIncomeRequest, opts ...grpc.CallOption) (*expense.CreateIncomeResponse, error) {
	if request == nil {
		return nil, ErrEmptyRequest
	}
	if request.Amount <= 0 {
		return nil, ErrAmountLessThanZero
	}
	incomeRequest := &db.CreateIncomeRequest{
		Income: &db.Income{
			Amount: request.Amount,
		},
	}
	createdIncome, err := h.incomeService.CreateIncome(ctx, incomeRequest)
	if err != nil {
		return nil, err
	}
	response := &expense.CreateIncomeResponse{
		IncomeId: createdIncome.GetIncomeId(),
	}

	return response, nil
}

func (h *Handler) GetIncome(ctx context.Context, in *expense.GetIncomeRequest, opts ...grpc.CallOption) (*expense.GetIncomeResponse, error) {
	if in == nil {
		return nil, ErrEmptyRequest
	}
	income, err := h.incomeService.GetIncome(ctx, &db.GetIncomeRequest{
		IncomeId: in.GetIncomeId(),
	})
	if err != nil {
		return nil, err
	}
	response := &expense.GetIncomeResponse{
		Income: &expense.Income{
			Amount: income.Amount,
		},
	}

	return response, nil
}

func (h *Handler) UpdateIncome(ctx context.Context, in *expense.UpdateIncomeRequest, opts ...grpc.CallOption) (*expense.UpdateIncomeResponse, error) {
	if in == nil {
		return nil, ErrEmptyRequest
	}
	updateRequest := &db.UpdateIncomeRequest{
		IncomeId: in.GetIncomeId(),
		NewIncome: &db.Income{
			Amount: in.GetAmount(),
		},
	}
	updatedIncome, err := h.incomeService.UpdateIncome(ctx, updateRequest)
	if err != nil {
		return nil, err
	}
	response := &expense.UpdateIncomeResponse{
		IncomeId: updatedIncome.GetIncomeId(),
	}

	return response, nil
}

func (h *Handler) DeleteIncome(ctx context.Context, in *expense.DeleteIncomeRequest, opts ...grpc.CallOption) (*expense.DeleteIncomeResponse, error) {
	incomeID := in.GetIncomeId()
	if incomeID == "" {
		return nil, ErrEmptyIncomeID
	}
	deleteRequest := &db.DeleteIncomeRequest{IncomeId: incomeID}

	err, _ := h.incomeService.DeleteIncome(ctx, deleteRequest)
	if err != nil {
		return nil, ErrDeleteIncome
	}
	return &expense.DeleteIncomeResponse{
		Message: fmt.Sprintf("income %s deleted successfully", incomeID),
	}, nil
}
