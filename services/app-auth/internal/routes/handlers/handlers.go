package handlers

import (
	"context"


	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	AuthUC         adapters.AuthUseCase
	SessionService service.DefaultSessionService
}

func (h Handler) CreateExpenseType(c *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (h Handler) Register(c *gin.Context) {
	// Validate the register request
	var data dto.RegisterRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Register the user
	registerResponse, err := h.AuthUC.Register(context.Background(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the register response data
	c.JSON(http.StatusOK, gin.H{"data": registerResponse})
}

func (h Handler) UserLogout(c *gin.Context) {
	// Get the user ID from the session
	sessionID := c.GetHeader("Authorization")

	// Clear the session
	err := h.SessionService.ClearSession(context.Background(), sessionID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the success message
	c.JSON(http.StatusOK, gin.H{"message": "User logged out successfully"})
}

func (h Handler) RefreshToken(c *gin.Context) {
	// Get the refresh token from the request
	var refreshToken string
	if err := c.ShouldBindJSON(&refreshToken); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Refresh the token
	refreshTokenResponse, err := h.AuthUC.RefreshToken(context.Background(), RefreshTokenRequest.RefreshToken)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the refresh token response
	c.JSON(http.StatusOK, gin.H{"data": refreshTokenResponse})
}

func (h Handler) Login(c *gin.Context) {
	// Validate the login request
	var data dto.LoginRequest
	if err := c.ShouldBindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Login the user
	loginResponse, err := h.AuthUC.Login(context.Background(), data)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Return the login response data
	c.JSON(http.StatusOK, gin.H{"data": loginResponse})
}

func (h Handler) GetUserById(c *gin.Context, userId string) {
	// Get the user from the database
	_, err := h.AuthUC.GetUserById(c, userId)
	if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return the user data
	c.JSON(http.StatusOK, gin.H{"data": user})
}

func NewHandler(authUC adapters.AuthUseCase) ServerInterface {
	return &Handler{
		AuthUC: authUC,
	}
}
