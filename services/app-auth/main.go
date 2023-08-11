package app_auth

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/repository"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/storage"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/usecase"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/routes/server"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
	"os/signal"
)

func main() {
	// Load config
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Loading config failed", err)
	}
	dbService := config.DatabaseService{
		Port: cfg.Database.Port,
		Host: cfg.Database.Host,
	}

	dbStorage, err := storage.NewDbStorage(dbService)
	if err != nil {
		log.Error(err)
		os.Exit(1)
	}
	cacheStorage := repository.NewAuthRedisRepository("")

	authRepo := repository.NewAuthRepo(dbStorage, cacheStorage)

	// services
	authService := service.NewAuthService(authRepo)

	authUsecase := usecase.NewAuthUsecase(authService, nil)
	// server config
	handler := server.NewServer(authUsecase, *cfg)

	serviceAddress := ":" + os.Getenv("PORT")
	srv := &http.Server{
		Addr:    serviceAddress,
		Handler: handler,
	}

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	go func() {
		err := srv.ListenAndServe()
		if err != nil {
			log.Fatal(err)
		}
	}()

	<-quit
	log.Println("Shutting down...")
	srv.Shutdown(context.Background())
	log.Println("Server stopped")
}
