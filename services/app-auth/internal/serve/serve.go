package serve

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/sanity-io/litter"
	"gorm.io/gorm"
	"log/slog"
	"net/http"
	"net/http/pprof"
	"os"
	"os/signal"
	"syscall"
	"time"
)

const (
	maxHeaderBytes = 1 << 20
	ctxTimeout     = 5
	certFile       = "ssl/Server.crt"
	keyFile        = "ssl/Server.pem"
)

type Server struct {
	authUsecase adapters.AuthUseCase

	r   *gin.Engine
	cfg *config.Config
	db  *gorm.DB
	//redisClient *redis.Client
	logger slog.Logger
}

func NewServe(authUsecase adapters.AuthUseCase, cfg *config.Config, db *gorm.DB, redisClient *redis.Client, logger slog.Logger) *Server {
	r := gin.Default()

	return &Server{
		r:   r,
		cfg: cfg,
		db:  db,
		//redisClient: redisClient,
		logger:      logger,
		authUsecase: authUsecase,
	}

}

func (s *Server) Run() error {
	if s.cfg.Server.SSL {
		server := &http.Server{
			Addr:           s.cfg.Server.Port,
			ReadTimeout:    time.Second * s.cfg.Server.ReadTimeout,
			WriteTimeout:   time.Second * s.cfg.Server.WriteTimeout,
			MaxHeaderBytes: maxHeaderBytes,
		}

		go func() {
			slog.Info("port", s.cfg.Server.Port, "Server is listening on PORT: %s", s.cfg.Server.Port)
			if err := server.ListenAndServeTLS(certFile, keyFile); err != nil {
				slog.Error("Error starting TLS Server")
			}
		}()

		go func() {
			s.logger.Info("Starting Debug Server on PORT: %s", s.cfg.Server.PprofPort)
			if err := http.ListenAndServe(s.cfg.Server.PprofPort, http.HandlerFunc(pprof.Index)); err != nil {
				s.logger.Error("Error PPROF ListenAndServe: %v", err)
			}
		}()

		quit := make(chan os.Signal, 1)
		signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

		<-quit

		ctx, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
		defer shutdown()

		s.logger.Info("Server Exited Properly")
		return server.Shutdown(ctx)
	}

	go func() {
		s.logger.Info("Server is listening on PORT: %s", s.cfg.Server.Port)
		if err := s.r.Run(s.cfg.Server.Port); err != nil {
			s.logger.Error("Error starting Server: %v", err)
		}
	}()

	go func() {
		s.logger.Info("Starting Debug Server on PORT: %s", s.cfg.Server.PprofPort)
		if err := http.ListenAndServe(s.cfg.Server.PprofPort, http.HandlerFunc(pprof.Index)); err != nil {
			s.logger.Error("Error PPROF ListenAndServe: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	<-quit

	_, shutdown := context.WithTimeout(context.Background(), ctxTimeout*time.Second)
	defer shutdown()
	s.logger.Info("Server Exited Properly")
	return nil
	//return serve.Shutdown(ctx)
}

// HandleRequestLogs is a middleware that logs incoming HTTP requests
func (s *Server) HandleRequestLogs() gin.HandlerFunc {
	return func(c *gin.Context) {
		startTime := time.Now()
		c.Next()
		endTime := time.Now()
		latency := endTime.Sub(startTime)

		s.logger.Info("Request: %s %s - %v", c.Request.Method, c.Request.URL.Path, latency)
	}
}

// HandleRecovery recovers from any panics in the request processing chain

func (s *Server) HandleRecovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				s.logger.Error("Panic recovered: %v", r)
				litter.Dump(c.Request)
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
