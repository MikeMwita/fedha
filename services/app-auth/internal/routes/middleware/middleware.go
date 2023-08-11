package middleware

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
)

type MiddlewareManager struct {
	Config      config.Config
	AuthUsecase adapters.AuthUseCase
}

func NewManager(cfg config.Config, authUseCase adapters.AuthUseCase) *MiddlewareManager {
	return &MiddlewareManager{
		Config:      cfg,
		AuthUsecase: authUseCase,
	}
}
