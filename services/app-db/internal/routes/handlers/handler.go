package handlers

import (
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/ports"
)

type Handler struct {
	db.UnimplementedDbServiceServer
	userRepo    ports.UserRepo
	expenseRepo ports.ExpenseRepo
}

func NewHandler(userRepo ports.UserRepo, expenseRepo ports.ExpenseRepo) *Handler {
	return &Handler{
		userRepo:    userRepo,
		expenseRepo: expenseRepo,
	}
}
