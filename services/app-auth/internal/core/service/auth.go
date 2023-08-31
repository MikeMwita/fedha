package service

import (
	"context"
	"errors"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg"
	"github.com/gin-gonic/gin"
)

var (
	ErrUsernameTaken = errors.New("username already taken")
	ErrEmptyUsername = errors.New("username cannot be empty")
	ErrEmptyPassword = errors.New("password cannot be empty")
)

type AuthService struct {
	repo   adapters.AuthRepo
	config config.Jwt
}

func (a AuthService) GetUserById(c *gin.Context, id string) (*entity.User, error) {
	req := &db.GetUserByIDRequest{
		UserId: id,
	}
	_, err := a.repo.GetUserByID(c, req)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return nil, err
	}
	return nil, err
}

func (a AuthService) Register(ctx *gin.Context, request dto.RegisterRequest) (*dto.RegisterResponseData, error) {
	var req dto.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
	}
	user := entity.User{
		Email:    string(request.Email),
		UserName: request.FullName,
		Hash:     request.Password,
	}
	if err := a.ValidateUser(user); err != nil {
		return nil, err
	}
	if err, _ := a.repo.FindByUsername(ctx, request.Username); err == nil {
		return nil, ErrUsernameTaken
	}

	return nil, nil
}

func (a AuthService) Login(request dto.LoginRequest) (*dto.LoginResponseData, error) {
	if err := a.ValidateLoginRequest(dto.LoginInitRequest(request)); err != nil {
		return nil, err
	}
	user, err := a.repo.Login(context.Background(), dto.LoginRequest(request))
	if err != nil {
		return nil, err
	}
	_, err = pkg.GenerateJWTToken(user.UserId)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseData{}, nil
}
func (a AuthService) ValidateLoginRequest(request dto.LoginInitRequest) error {
	if request.Username == "" {
		return ErrEmptyUsername
	}

	if request.Password == "" {
		return ErrEmptyPassword
	}

	return nil
}

func (a AuthService) RefreshToken(request dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	if err := a.validateRefreshTokenRequest(request); err != nil {
		return nil, err
	}
	if err := a.repo.VerifyRefreshToken(request.RefreshToken); err != nil {
		return nil, err
	}
	user := entity.User{
		UserId: request.RefreshToken,
	}
	token, err := pkg.GenerateJWTToken(user)
	if err != nil {
		return nil, err
	}

	return &dto.RefreshTokenResponse{
		AccessToken: token,
	}, nil
}

func (a AuthService) validateRefreshTokenRequest(request dto.RefreshTokenRequest) error {
	if request.RefreshToken == "" {
		return errors.New("refresh token cannot be empty")
	}
	return nil
}

func (a AuthService) UpdateUser(user entity.User) (*entity.User, error) {
	if err := a.ValidateUser(user); err != nil {
		return nil, err
	}
	err, _ := a.repo.UpdateUser(context.Background(), user)
	if err != nil {
		return nil, errors.New("User not updated")
	}
	return &user, nil
}

func (a AuthService) ValidateUser(user entity.User) error {
	if user.UserId == "" {
		return errors.New("user ID cannot be empty")
	}
	if user.UserName == "" {
		return errors.New("username cannot be empty")
	}
	if user.Email == "" {
		return errors.New("email cannot be empty")
	}

	if user.PhoneNumber == "" {
		return errors.New("phone number cannot be empty")
	}

	return nil
}

func (a AuthService) UserLogout(userUUID string) error {
	err := a.repo.DeleteAccessToken(context.Background(), userUUID)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthService(repo adapters.AuthRepo) adapters.AuthService {
	return &AuthService{repo: repo}
}
