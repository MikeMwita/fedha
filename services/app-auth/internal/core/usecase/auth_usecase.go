package usecase

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/gin-gonic/gin"
)

type AuthUsecase struct {
	authService    adapters.AuthService
	sessionService adapters.SessionService
}

func (a AuthUsecase) Login(c *gin.Context, data dto.LoginRequest) (*dto.LoginResponseData, error) {
	res, err := a.authService.Login(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a AuthUsecase) Register(c *gin.Context, data dto.RegisterRequest) (*dto.RegisterResponseData, error) {
	res, err := a.authService.Register(c, data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a AuthUsecase) VerifyAccessToken(token string) (interface{}, interface{}) {
	//TODO implement me
	panic("implement me")
}

func (a AuthUsecase) RefreshToken(c *gin.Context, data dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	res, err := a.authService.RefreshToken(data)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a AuthUsecase) UpdateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	res, err := a.authService.UpdateUser(user)
	if err != nil {
		return nil, err
	}
	return res, nil
}

func (a AuthUsecase) UserLogout(c *gin.Context) {

}

func (a AuthUsecase) GetUserById(c *gin.Context, userId string) (string, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthUsecase(as adapters.AuthService, ss adapters.SessionService) adapters.AuthUseCase {
	return &AuthUsecase{
		authService:    as,
		sessionService: ss}
}
