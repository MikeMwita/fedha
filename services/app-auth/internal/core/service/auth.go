package service

import (
	"context"
	"errors"
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

func (a AuthService) GetUserById(c *gin.Context, userId string) {

	// Get the user from the database
	user, err := a.repo.GetUserById(c, userId)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}
	// Assign the returned object to a variable
	u := user

	// Use the variable
	c.JSON(200, u)
}

func (a AuthService) Register(ctx *gin.Context, request dto.RegisterReq) (*dto.RegisterRes, error) {
	// Get the context from the Gin request

	//ctx := c.Request.Context()
	//validate register request
	var req dto.RegisterReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
	}
	// Create a new user
	user := entity.User{
		Email:    string(request.Email),
		UserName: request.FullName,
		Hash:     request.Password,
	}

	if err := a.ValidateUser(user); err != nil {
		return nil, err
	}
	// Check if the username is already taken
	if err := a.repo.FindByUsername(ctx, request.Username); err == nil {
		return nil, ErrUsernameTaken
	}

	// Save the user
	err := a.repo.Save(ctx, user)
	if err != nil {
		return nil, err
	}

	// Return the register response data
	return &dto.RegisterRes{
		CreatedAt:         user.CreatedAt,
		Email:             user.Email,
		FullName:          user.FullName,
		PasswordChangedAt: user.PasswordChangedAt,
		Username:          user.UserName,
	}, nil

}

func (a AuthService) Login(request dto.LoginInitRequest) (*dto.LoginInitResponseData, error) {
	// Validate the login request
	if err := a.ValidateLoginRequest(request); err != nil {
		return nil, err
	}

	// Get the user from the repository
	user, err := a.repo.Login(context.Background(), dto.LoginRequest(request))
	if err != nil {
		return nil, err
	}

	// Compare the password
	//if !bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(LoginRequest.Password)) {
	//	return nil, errors.New("invalid password")
	//}

	// Generate a JWT token
	token, err := pkg.GenerateJWTToken(user.UserId)
	if err != nil {
		return nil, err
	}

	// Return the login response data
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

//func NewAuthService(repo adapters.AuthRepo) adapters.AuthService {
//	return &AuthService{repo: repo}
//}
