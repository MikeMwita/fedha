package adapters

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/gin-gonic/gin"
	"time"
)

type AuthRepo interface {
	Register(ctx context.Context, registerReq dto.RegisterRequest) (*dto.RegisterResponseData, error)
	Login(ctx context.Context, loginReq dto.LoginRequest) (*dto.LoginResponseData, error)
	GetUserById(c *gin.Context, userId string) (string, error)
	RefreshToken(c *gin.Context, data dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
	UserLogout(c *gin.Context)
	SetAccessToken(ctx context.Context, key string, value interface{}, expiration time.Duration) error
	GetAccessToken(ctx context.Context, value interface{}) error
	DeleteAccessToken(ctx context.Context, key string) error
	FindByUsername(ctx interface{}, username interface{}) interface{}
	Save(ctx interface{}, user entity.User) error
	DeleteSession(ctx context.Context, record *service.SessionRecord) error
	InvalidateSession() error
	VerifyRefreshToken(token string) error
}
