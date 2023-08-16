package handlers

import (
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/ports"
	"github.com/labstack/gommon/log"
)

type Handler struct {
	expense.UnimplementedExpenseServiceServer
	expenseService         ports.ExpenseService
	balanceService         ports.BalanceService
	incomeService          ports.IncomeService
	monthly_summaryService ports.MonthlyService
}

func (h Handler) mustEmbedUnimplementedExpenseServiceServer() {

	log.Error("not implemented")
}

func NewHandler(expenseService ports.ExpenseService, balanceService ports.BalanceService, incomeService ports.IncomeService, monthly_summaryService ports.MonthlyService) *Handler {
	return &Handler{
		expenseService:         expenseService,
		balanceService:         balanceService,
		incomeService:          incomeService,
		monthly_summaryService: monthly_summaryService,
	}
}
