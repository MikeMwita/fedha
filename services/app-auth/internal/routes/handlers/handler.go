package handlers

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core /adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core /service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type Handler struct {
	AuthService    service.AuthService
	SessionService service.DefaultSessionService
}

func (h Handler) Register(ctx context.Context, user dto.RegisterReq) (*dto.RegisterRes, error) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Login(ctx context.Context, user dto.LoginInitRequest) (*dto.LoginInitResponseData, error) {
	//TODO implement me
	panic("implement me")
}

func NewHandler(authService service.AuthService, sessionService service.DefaultSessionService) adapters.AuthUseCase {
	return &Handler{
		AuthService:    authService,
		SessionService: sessionService,
	}
}
