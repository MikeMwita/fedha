package model

import (
	"google.golang.org/genproto/googleapis/type/money"
)

type Expense struct {
	ExpenseID   int64        `json:"expenseID"`
	UserID      int64        `json:"userID"`
	Date        string       `json:"date"`
	Recipient   string       `json:"recipient"`
	Description string       `json:"description"`
	Category    string       `json:"category"`
	Value       *money.Money `json:"value"`
	Version     int          `json:"version"`
}

type Balance struct {
	BalanceID   int64        `json:"balanceID"`
	UserID      int64        `json:"userID"`
	Amount      *money.Money `json:"amount"`
	Description string       `json:"description"`
	Timestamp   string       `json:"timestamp"`
	Version     int          `json:"version"`
}

type Income struct {
	IncomeID    int64        `json:"balanceID"`
	UserID      int64        `json:"userID"`
	Amount      *money.Money `json:"amount"`
	Description string       `json:"description"`
	Timestamp   string       `json:"timestamp"`
	Version     int          `json:"version"`
}
