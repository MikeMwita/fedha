package adapters

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

type AuthRepo interface {
	Register(ctx context.Context, data dto.RegisterReq) (*dto.RegisterRes, error)
	Login(ctx context.Context, data dto.LoginInitRequest) (*dto.LoginInitResponseData, error)
}

type AuthRedisRepo interface {
	SaveAccessToken(ctx context.Context, userID, accessToken string) error
	GetAccessToken(ctx context.Context, userID string) (string, error)
	DeleteAccessToken(ctx context.Context, userID string) error
}
