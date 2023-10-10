package cmd

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/repository"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/storage"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/usecase"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/routes/server"
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/tracing"
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/util"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/trace"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
)

func Execute() {

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))

	logger.Info("Starting Authentication microservice...")

	confPath := util.GetConfigPath(os.Getenv("config"))
	cfg, err := config.GetConfig(confPath)
	if err != nil {
		logger.Error("Loading and parsing config failed")
		log.Fatalf("GetConfig: %v", err)
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Setting up the tracer
	var _ trace.Tracer
	if cfg.IsTracingEnabled {
		traceProvider := tracing.SetupTracer(ctx, cfg)
		otel.SetTracerProvider(traceProvider)
		_ = traceProvider.Tracer("your-app-here")

		// Configuring   OTel exporter
		_, err := otlptracehttp.New(
			ctx,
			otlptracehttp.WithEndpoint("http://localhost:5000/api/traces"),
			otlptracehttp.WithInsecure(),
		)
		if err != nil {
			logger.Error("Failed to create OpenTelemetry exporter")
			log.Fatalf("otlptracehttp.New: %v", err)
		}

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

	authUsecase := usecase.NewAuthUsecase(cfg, authService, authRepo, nil)

	handler := server.NewServer(authUsecase, cfg)

	serviceAddress := ":" + os.Getenv("PORT")

	srv := &http.Server{
		Addr:    serviceAddress,
		Handler: handler.Router,
	}

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
