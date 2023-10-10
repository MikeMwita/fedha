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
	repo   adapters.AuthRepository
	config config.Jwt
}

func (a AuthService) GetUserById(c *gin.Context, userId string) {
	req := &db.GetUserByIDRequest{
		UserId: userId,
	}
	user, err := a.repo.GetUserByID(c, req)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	u := user
	c.JSON(200, u)
}

func (a AuthService) Register(ctx *gin.Context, request dto.RegisterRequest) (*dto.RegisterResponseData, error) {
	//ctx := c.Request.Context()
	var req dto.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
	}
	user := entity.User{
		FullName: request.FullName,
		UserName: request.FullName,
		Hash:     request.Password,
	}
	if err := a.ValidateUser(user); err != nil {
		return nil, err
	}
	// Check if the username is already taken
	if err, _ := a.repo.FindByUsername(ctx, request.Username); err == nil {
		return nil, ErrUsernameTaken
	}

	return &dto.RegisterResponseData{
		CreatedAt:         user.CreatedAt,
		PasswordChangedAt: user.PasswordChangedAt,
	}, nil

}

func (a AuthService) Login(request dto.LoginInitRequest) (*dto.LoginInitResponseData, error) {
	if err := a.ValidateLoginRequest(request); err != nil {
		return nil, err
	}

	// Get the user from the repository
	user, err := a.repo.Login(context.Background(), dto.LoginRequest(request))
	if err != nil {
		return nil, err
	}

	token, err := pkg.GenerateJWTToken(user.UserId)
	if err != nil {
		return nil, err
	}

	return &dto.LoginInitResponseData{
		AccessToken: &token,
	}, nil

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
	// Validate the refresh token request
	if err := a.validateRefreshTokenRequest(request); err != nil {
		return nil, err
	}

	// Verify the refresh token
	if err := a.repo.VerifyRefreshToken(request.RefreshToken); err != nil {
		return nil, err
	}

	// Generate a new JWT token
	user := entity.User{
		UserId: request.RefreshToken,
	}
	token, err := pkg.GenerateJWTToken(user)
	if err != nil {
		return nil, err
	}

	// Return the refresh token response data
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
	// Validate the user
	if err := a.ValidateUser(user); err != nil {
		return nil, err
	}

	// Update the user in the database
	err, _ := a.repo.UpdateUser(context.Background(), user)
	if err != nil {
		return nil, errors.New("User not updated")
	}

	// Return the updated user
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
	// Delete the user's JWT token
	err := a.repo.DeleteAccessToken(context.Background(), userUUID)
	if err != nil {
		return err
	}

	return nil
}

func NewAuthService(repo adapters.AuthRepository) adapters.AuthService {
	return &AuthService{repo: repo}
}
