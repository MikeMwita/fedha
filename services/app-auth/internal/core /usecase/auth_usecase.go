package usecase

import "github.com/MikeMwita/fedha.git/services/app-auth/internal/core /adapters"

type AuthUseCase struct {
	authService adapters.AuthService
}

func NewAuthUseCase(as adapters.AuthService, ss adapters.SessionService) *adapters.AuthUseCase {
	return &AuthUseCase{
		authService:    as,
		sessionService: ss,
	}
}
