package server

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/routes/handlers"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/routes/middleware"
	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api"

type Config struct {
	AuthUsecase adapters.AuthUseCase
}

func NewServer(authUseCase adapters.AuthUseCase, cfg config.Config) *gin.Engine {
	r := gin.Default()

	middlewareManager := middleware.NewManager(cfg, authUseCase)

	middlewares := []handlers.MiddlewareFunc{
		middlewareManager.Auth,
	}

	options := handlers.GinServerOptions{
		BaseURL:     BaseUrl,
		Middlewares: middlewares,
	}

	// Map handlers
	h := handlers.NewHandler(authUseCase)
	handlers.RegisterHandlersWithOptions(r, h, options)

	return r
}
