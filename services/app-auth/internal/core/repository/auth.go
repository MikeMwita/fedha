package repository

import (
	"context"
	"errors"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"math/rand"
	"time"
)

type authRepo struct {
	dbStorage    adapters.DbStorage
	cacheStorage adapters.CacheStorage
}

func (a authRepo) GetAccessToken(ctx context.Context, value interface{}) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) Save(ctx interface{}, user entity.User) error {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) VerifyRefreshToken(token string) error {
	//TODO implement me
	panic("implement me")
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
	user, err := a.dbStorage.GetUserByUsername(ctx, loginReq.Username)
	if err != nil {
		return nil, errors.New("invalid username")
	}
	// Compare the password
	//if !bcrypt.CompareHashAndPassword([]byte(user.Hash), []byte(LoginRequest.Password)) {
	//	return nil, errors.New("invalid password")
	//}
	return user, nil

}

func (a authRepo) GetUserById(c *gin.Context, userId string) {
	user, err := a.dbStorage.GetUserByID(c, userId)
	if err != nil {
		c.JSON(404, gin.H{"error": "user not found"})
		return
	}

	c.JSON(200, user)
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

	// Invalidate the user's session
	c.SetCookie("session_id", "", -1, "/", "", false, true)

	// Return a success message
	c.JSON(200, gin.H{"message": "Successfully logged out"})
}
func (a authRepo) SetAccessToken(ctx context.Context, key string, value interface{}, expiration time.Duration) error {
	// Validate the key
	if key == "" {
		return errors.New("key cannot be empty")
	}

	// Validate the value
	if value == nil {
		return errors.New("value cannot be nil")
	}

	// Set the expiration time
	if expiration <= 0 {
		expiration = time.Hour * 24
	}

	// Set the token in the cache
	err := a.cacheStorage.SetAccessToken(ctx, key, value, expiration)
	if err != nil {
		return err
	}

	return nil
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

func (a authRepo) FindByUsername(ctx interface{}, username interface{}) interface{} {
	// Get the username
	usernameStr, ok := username.(string)
	if !ok {
		return nil
	}

	// Find the user by username
	user, err := a.dbStorage.GetUserByUsername(ctx, usernameStr)
	if err != nil {
		return nil
	}

	// Return the user
	return user
}

func (a authRepo) SaveUser(ctx interface{}, user *entity.User) interface{} {
	// Validate the user
	if user.UserId == "" {
		return nil
	}

	// Hash the password
	hashedPassword, err := hashPassword(user.Password)
	if err != nil {
		return nil
	}

	// Set the hashed password on the user
	user.PasswordHash = hashedPassword

	// Save the user in the database
	_, err = a.dbStorage.SaveUser(ctx, user)
	if err != nil {
		return nil
	}

	// Return the user
	return user
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
