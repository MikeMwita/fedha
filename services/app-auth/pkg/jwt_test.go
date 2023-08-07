package pkg

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/golang-jwt/jwt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestGenerateToken(t *testing.T) {
	user := entity.User{
		UserName: "testuser",
		UserId:   "123",
	}
	tokenString, err := GenerateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)
}

func TestValidateToken(t *testing.T) {
	user := entity.User{
		UserName: "testuser",
		UserId:   "123",
	}
	tokenString, err := GenerateToken(user)
	assert.NoError(t, err)
	assert.NotEmpty(t, tokenString)

	userID, username, err := validateToken(tokenString)
	assert.NoError(t, err)
	assert.Equal(t, user.UserId, userID)
	assert.Equal(t, user.UserName, username)

	t.Run("Invalid token", func(t *testing.T) {
		userID, username, err := validateToken("invalid-token")
		assert.Error(t, err)
		assert.Empty(t, userID)
		assert.Empty(t, username)
	})

	t.Run("Unexpected signing method", func(t *testing.T) {
		token := jwt.New(jwt.SigningMethodHS256)
		token.Claims = jwt.MapClaims{
			"userId": user.UserId,
			"sub":    user.UserName,
			"exp":    time.Now().Add(24 * time.Hour).Unix(),
		}
		tokenString, _ := token.SignedString([]byte("invalid-key"))
		userID, username, err := validateToken(tokenString)
		assert.Error(t, err)
		assert.Empty(t, userID)
		assert.Empty(t, username)
	})
}

func TestMain(m *testing.M) {
	os.Setenv("JWT_SECRET_KEY", "secret-key")
	code := m.Run()
	os.Exit(code)
}
