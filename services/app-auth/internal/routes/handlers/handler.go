package handlers

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	service2 "github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	AuthService    service2.AuthService
	SessionService service2.DefaultSessionService
}

func Register(c *gin.Context) {
	var registerReq dto.RegisterReq

	if err := c.ShouldBindJSON(&registerReq); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := h.AuthService.Register(c.Request.Context(), registerReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	session, err := h.SessionService.CreateSession(c.Request.Context(), user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"session_id": session.ID,
	})
}

//
//func (h Handler) Register(ctx context.Context, user dto.RegisterReq) (*dto.RegisterRes, error) {
//	//var registerRequest *dto.RegisterReq
//	registerRequest := &dto.RegisterReq{}
//	if err := ctx.ShouldBindJSON(&registerRequest); err != nil {
//		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
//		return nil, err
//	}
//	user, err := h.AuthService.Register(ctx, registerRequest)
//	if err != nil {
//		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return nil, err
//	}
//
//	session, err := h.SessionService.CreateSession(ctx, user.ID)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return nil, err
//	}
//	return *dto.RegisterRes{
//		SessionID: session.Id,
//	}, nil
//
//}

func (h Handler) Login(ctx context.Context, user dto.LoginInitRequest) (*dto.LoginInitResponseData, error) {
	var loginRequest *dto.LoginInitRequest

}

func NewHandler(authService service2.AuthService, sessionService service2.DefaultSessionService) adapters.AuthUseCase {
	return &Handler{
		AuthService:    authService,
		SessionService: sessionService,
	}
}
