package usecase

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core /adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type AuthUseCase struct {
	authService    adapters.AuthService
	sessionService adapters.SessionService
}

func (a AuthUseCase) Register(ctx context.Context, user dto.RegisterReq) (*dto.RegisterRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUseCase) Login(ctx context.Context, user dto.LoginInitRequest) (*dto.LoginInitResponseData, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthUseCase(as adapters.AuthService, ss adapters.SessionService) adapters.AuthUseCase {
	return &AuthUseCase{
		authService:    as,
		sessionService: ss,
	}
}
