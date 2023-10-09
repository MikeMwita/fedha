package cmd

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/repository"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/storage"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/usecase"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/routes/server"
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/util"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
)

func Execute() {
	//var logger slog.Logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	//logger := log.New()

	logger.Info("Starting Authentication microservice...")

	confPath := util.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(confPath)
	if err != nil {
		logger.Error("Loading and parsing config failed")
		log.Fatalf("GetConfig: %v", err)
	}

	dbService := config.Database{
		Port: cfg.Database.Host,
		Host: cfg.Database.Port,
	}

	dbStorage, err := storage.NewDbStorage(dbService)
	if err != nil {
		logger.Error("Error creating database storage")
		os.Exit(1)
	}

	cacheStorage := repository.NewAuthRedisRepository("")
	authRepo := repository.NewAuthRepo(dbStorage, cacheStorage)

	authService := service.NewAuthService(authRepo)

	//authUsecase := usecase.NewAuthUsecase(authService, nil)

	authUsecase := usecase.NewAuthUsecase(cfg, authService, authRepo, nil)

	//sessionService := service.NewSessionService(authRepo)
	handler := server.NewServer(authUsecase, cfg)

	//handler := server.NewServer(authUsecase)

	serviceAddress := ":" + os.Getenv("PORT")

	srv := &http.Server{
		Addr:    serviceAddress,
		Handler: handler.Router,
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			logger.Error("Failed to start serve")
		}
	}()

	<-quit
	logger.Info("Shutting down...")

	srv.Shutdown(context.Background())
	logger.Info("Server stopped")
}
