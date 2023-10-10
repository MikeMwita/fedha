package adapters

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/gin-gonic/gin"
)

type AuthService interface {
	Register(ctx *gin.Context, request dto.RegisterRequest) (*dto.RegisterResponseData, error)
	Login(request dto.LoginInitRequest) (*dto.LoginInitResponseData, error)
	RefreshToken(request dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
	UpdateUser(user entity.User) (*entity.User, error)
	UserLogout(userUUID string) error
	GetUserById(c *gin.Context, userId string)
}
