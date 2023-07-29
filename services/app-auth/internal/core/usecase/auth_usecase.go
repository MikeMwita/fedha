package usecase

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type AuthUseCase struct {
	authService    adapters.AuthService
	sessionService adapters.SessionService
}

func (a AuthUseCase) Register(ctx context.Context, user dto.RegisterReq) (*dto.RegisterRes, error) {
	// Call the authentication service to handle user registration
	registerResponse, err := a.Register(ctx, user)
	if err != nil {
		return nil, err
	}
	return registerResponse, nil
}

func (a AuthUseCase) Login(ctx context.Context, user dto.LoginInitRequest) (*dto.LoginInitResponseData, error) {
	loginResponse, err := a.Login(ctx, user)
	if err != nil {
		// Handle any errors that may occur during the login process
		return nil, err
	}

	// Return the login response
	return loginResponse, nil
}

func NewAuthUseCase(as adapters.AuthService, ss adapters.SessionService) adapters.AuthUseCase {
	return &AuthUseCase{
		authService:    as,
		sessionService: ss,
	}
}
