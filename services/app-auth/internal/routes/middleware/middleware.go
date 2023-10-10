package middleware

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/platform/Monitoring/Prometheus"
	"github.com/gin-gonic/gin"
)

type MiddlewareManager interface {
	Auth(ctx *gin.Context)
	Cors(ctx *gin.Context)
}

type middlewareManager struct {
	Config      config.Config
	AuthUsecase adapters.AuthUseCase
	promMetrics *Prometheus.Metrics
}

func NewMiddlewareManager(cfg config.Config, authUsecase adapters.AuthUseCase, promMetrics *Prometheus.Metrics) MiddlewareManager {
	return &middlewareManager{
		Config:      cfg,
		AuthUsecase: authUsecase,
		promMetrics: promMetrics,
	}
}
