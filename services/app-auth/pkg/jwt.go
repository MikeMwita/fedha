package pkg

import (
	"errors"
	"fmt"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/golang-jwt/jwt"
	"os"
	"time"
)

var jwtKey = []byte(os.Getenv("JWT_SECRET_KEY"))

type authClaims struct {
	jwt.StandardClaims
	UserID string `json:"userId"`
}

func GenerateJWTToken(user entity.User) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, authClaims{
		StandardClaims: jwt.StandardClaims{
			Subject:   user.UserName,
			ExpiresAt: expiresAt,
		},
		UserID: user.UserId,
	})

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// validateToken parses and validates the JWT token and returns the user ID and username if valid.
func ValidateToken(tokenString string) (string, string, error) {
	var claims authClaims
	token, err := jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtKey, nil
	})
	if err != nil {
		return "", "", err
	}

	if !token.Valid {
		return "", "", errors.New("invalid token")
	}

	userID := claims.UserID
	username := claims.Subject
	return userID, username, nil
}
