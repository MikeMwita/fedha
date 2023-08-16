package handlers

import (
	"context"
	"errors"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"google.golang.org/grpc"
)

//func (h Handler) CreateTrade(ctx context.Context, request *exchange.CreateTradeRequest) (*exchange.CreateTradeResponse, error) {
//	if request == nil {
//		return nil, ErrEmptyRequest
//	}
//
//	tradeRequest := &db.CreateTradeRequest{
//		Trade: &db.Trade{
//			AccountId:    request.Trade.AccountId,
//			TradeType:    request.Trade.TradeType,
//			TradeStatus:  request.Trade.TradeStatus,
//			FromCurrency: request.Trade.FromCurrency,
//			ToCurrency:   request.Trade.ToCurrency,
//			FromAmount:   request.Trade.FromAmount,
//		},
//	}
//
//	tradeResponse, err := h.tradeService.CreateTrade(ctx, tradeRequest)
//	if err != nil {
//		return nil, err
//	}
//
//	tradeResponseResult := &exchange.CreateTradeResponse{
//		TradeId: tradeResponse.TradeId,
//	}
//
//	return tradeResponseResult, nil
//}
var (
	ErrEmptyRequest = errors.New("empty request")
	ErrAmountLessThanZero = errors.New("amount less than zero")
)
(ctx context.Context, request *expense.CreateIncomeRequest) (*expense.CreateIncomeResponse, error)
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
			Date:   request.Date.AsTime(),
		},
	}
	createdIncome,err := h.incomeService.CreateIncome(ctx, incomeRequest)
	if err != nil {
		return nil, err
	}
	return createdIncome,nil
}

func (h *Handler) GetIncome(ctx context.Context, in *expense.GetIncomeRequest, opts ...grpc.CallOption) (*expense.GetIncomeResponse, error) {

}

func (h *Handler) UpdateIncome(ctx context.Context, in *expense.UpdateIncomeRequest, opts ...grpc.CallOption) (*expense.UpdateIncomeResponse, error) {

}

func (h *Handler) DeleteIncome(ctx context.Context, in *expense.DeleteIncomeRequest, opts ...grpc.CallOption) (*expense.DeleteIncomeResponse, error) {

}
