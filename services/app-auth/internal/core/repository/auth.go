package repository

import (
	"context"
	"errors"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/validation"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"math/rand"
	"time"
)

type authRepo struct {
	dbStorage    adapters.DbStorage
	cacheStorage adapters.CacheStorage
}

func (a authRepo) SetAccessToken(ctx context.Context, key string, value string, expiration time.Duration) error {
	// Validate the key
	if key == "" {
		return errors.New("key cannot be empty")
	}
	if len(value) == 0 {
		return errors.New("value cannot be nil")
	}
	err := a.cacheStorage.SetAccessToken(ctx, key, value, expiration)
	if err != nil {
		return err
	}

	return nil
}

func (a authRepo) GetAccessToken(ctx context.Context, value string) (string, error) {
	token, err := a.cacheStorage.GetAccessToken(ctx, value)
	if err != nil {
		return "", err
	}
	tokenExpiration := time.Now().Add(24 * time.Hour)
	if token != "" && time.Now().After(tokenExpiration) {
		// delete expired token
		err = a.cacheStorage.DeleteAccessToken(ctx, value)
		if err != nil {
			return "", err
		}
		token = ""
	}
	return token, nil
}

func (a authRepo) FindByUsername(ctx context.Context, username string) (*db.RegUserRes, error) {
	//req object
	request := &db.GetUserByUsernameRequest{
		Username: username,
	}

	user, err := a.dbStorage.GetUserByUsername(ctx, request)
	if err != nil {
		return user, err
	}

	// Return the user
	return user, nil
}

func (a authRepo) SaveUser(ctx context.Context, in *db.SaveUserRequest, opts ...grpc.CallOption) (*db.User, error) {
	// Validate  input request
	if in == nil || in.User == nil {
		return nil, errors.New("invalid request")
	}
	// Hash the password
	hashedPassword, err := hashPassword(in.User.Hash)
	if err != nil {
		return nil, err
	}
	// Set hashed password on the user object
	in.User.Hash = hashedPassword

	user, err := a.dbStorage.SaveUser(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return user, err
}

func (a authRepo) VerifyRefreshToken(token string) error {
	return errors.New("not implemented")
}

func (a authRepo) GetUserByID(ctx context.Context, in *db.GetUserByIDRequest, opts ...grpc.CallOption) (*db.RegUserRes, error) {
	// Validate input request
	if in == nil || in.UserId == "" {
		return nil, errors.New("invalid request")
	}
	user, err := a.dbStorage.GetUserByID(ctx, in, opts...)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (a authRepo) InvalidateSession() error {
	return nil
}

func (a authRepo) DeleteSession(ctx context.Context, record *service.SessionRecord) error {
	err := a.cacheStorage.DeleteSession(ctx, record.Id)
	if err != nil {
		return err
	}
	return nil
}

func (a authRepo) Register(ctx context.Context, registerReq dto.RegisterRequest) (*dto.RegisterResponseData, error) {
	user := db.RegUserReq{

		FullName:     registerReq.FullName,
		Email:        string(registerReq.Email),
		PasswordHash: registerReq.Password,
		Username:     registerReq.Username,
	}

	// Hash the password
	hashedPassword, err := hashPassword(registerReq.Password)
	if err != nil {
		return nil, err
	}
	// Set the hashed password on the user
	user.PasswordHash = hashedPassword

	// Register the user in the database
	_, err = a.dbStorage.RegisterUser(ctx, &user)
	if err != nil {
		return nil, err
	}

	// Return the register response data
	return &dto.RegisterResponseData{}, nil

}

func hashPassword(password string) (string, error) {
	// Generate a hash of the password
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	// Return the hash
	return string(hash), nil
}

func (a authRepo) Login(ctx context.Context, loginReq dto.LoginRequest) (*dto.LoginResponseData, error) {
	// Find the user by username
	_, err := a.dbStorage.GetUserByUsername(ctx, &db.GetUserByUsernameRequest{Username: loginReq.Username})
	if err != nil {
		return nil, errors.New("invalid username")
	}

	//loginReq := dto.LoginRequest{
	//	Username: loginReq.Username,
	//	Password: loginReq.Password,
	//}
	err = validation.ValidateLogin(ctx, dto.LoginInitRequest(loginReq))
	if err != nil {
		return nil, err
	}

	// Return the login response data
	return &dto.LoginResponseData{
		//User: user,
	}, nil
}

func (a authRepo) UpdateUser(ctx context.Context, user entity.User) (*entity.User, error) {
	// Validate the user
	if user.UserId == "" {
		return nil, errors.New("user ID cannot be empty")
	}
	userReq := &db.UpdateUserReq{
		UserId:   user.UserId,
		Email:    user.Email,
		UserName: user.UserName,
	}
	// Update the user in the database
	updatedUser, err := a.dbStorage.UpdateUser(ctx, userReq)
	if err != nil {
		return nil, err
	}

	// Return the updated user
	return &entity.User{
		UserId:   updatedUser.UserId,
		Email:    updatedUser.Email,
		UserName: updatedUser.UserName,
	}, nil
}

func (a authRepo) UserLogout(c *gin.Context) {
	// Get the user ID from the session
	userID := c.MustGet("userID").(string)

	// Delete the user's JWT token from the cache
	err := a.cacheStorage.DeleteAccessToken(c, userID)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	// Return a success message
	c.JSON(200, gin.H{"message": "Successfully logged out"})
}

func (a authRepo) DeleteAccessToken(ctx context.Context, key string) error {
	// Validate the key
	if key == "" {
		return errors.New("key cannot be empty")
	}

	// Delete the token from the cache
	err := a.cacheStorage.DeleteAccessToken(ctx, key)
	if err != nil {
		return err
	}

	return nil
}

func (a authRepo) RefreshToken(c *gin.Context, data dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error) {
	// Validate the refresh token
	if data.RefreshToken == "" {
		return nil, errors.New("refresh token cannot be empty")
	}

	// Return the refresh token response
	accessToken := "accessToken"
	return &dto.RefreshTokenResponse{
		AccessToken: accessToken,
		//ExpiresIn:  time.Now().Add(time.Hour * 24),

	}, nil
}

func generateRandomString(length int) string {
	// Generate a random string of the specified length.
	b := make([]byte, length)
	for i := range b {
		b[i] = byte(rand.Intn(256))
	}

	// Return the random string.
	return string(b)
}
func NewAuthRepo(dbStorage adapters.DbStorage, cacheStorage adapters.CacheStorage) adapters.AuthRepo {
	return &authRepo{
		dbStorage:    dbStorage,
		cacheStorage: cacheStorage,
	}

}
