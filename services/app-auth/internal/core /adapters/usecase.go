package adapters

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type AuthUseCase interface {
	Register(ctx context.Context, user dto.RegisterReq) (*dto.RegisterRes, error)
	Login(ctx context.Context, user dto.LoginInitRequest) (*dto.LoginInitResponseData, error)
}
