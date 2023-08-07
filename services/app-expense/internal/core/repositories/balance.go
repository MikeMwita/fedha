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

	monthlySummary, err := b.dbStorage.GenerateMonthlySummary(ctx, in)
	if err != nil {
		return nil, err
	}
	return monthlySummary, nil

	////get total monthly expenses
	//totalExpenses, err := b.dbStorage.GetIncome(ctx, in.Month)
	//if err != nil {
	//	return nil, err
	//}
	//
	////total monthly income
	//totalIncome, err := b.dbStorage.GetIncome(ctx, in.Year, in.Month)
	//if err != nil {
	//	return nil, err
	//}
	////calcluating remaining  balance
	//
	//remainingBalance := totalIncome - totalExpenses
	//
	//monthlySummary := &db.MonthlySummaryResponse{
	//	Year:             in.Year,
	//	Month:            in.Month,
	//	TotalExpenses:    totalExpenses,
	//	TotalIncome:      totalIncome,
	//	RemainingBalance: remainingBalance,
	//}
	//return monthlySummary, nil
}

func (b BalanceRepository) GetRemainingBalance(ctx context.Context, in *db.RemainingBalanceRequest, opts ...grpc.CallOption) (*db.RemainingBalanceResponse, error) {
	remainingBalance, err := b.dbStorage.GetRemainingBalance(ctx, in)
	if err != nil {
		return nil, err
	}
	return remainingBalance, nil

	////total expenses
	//totalExpenses, err := b.dbStorage.GetRemainingBalance(ctx, in.Dates)
	//if err != nil {
	//	return nil, err
	//}
	//// Get the total income for the dates.
	//totalIncome, err := b.dbStorage.GetIncome(ctx, in.Dates)
	//if err != nil {
	//	return nil, err
	//}
	//// Calculate the remaining balance.
	//remainingBalance := totalIncome - totalExpenses
	//
	//// Create a new `RemainingBalanceResponse` object.
	//remainingBalanceResponse := &db.RemainingBalanceResponse{
	//	RemainingBalance: remainingBalance,
	//}
	//return remainingBalanceResponse, nil

}

func NewBalanceRepository(dbStorage ports.DbStorage) ports.BalanceRepo {

	return &MonthlyRepo{
		dbStorage: dbStorage,
	}
}
