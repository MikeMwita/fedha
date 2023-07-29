package server

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-db/config"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/repository"
	"github.com/MikeMwita/fedha.git/services/app-db/internal/core/storage"
	"github.com/jackc/pgx/v4/pgxpool"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"net"
)

type Server struct {
	cfg config.Config
}

func (s *Server) Run() {
	log.Info("GRPC Server INITIALIZING")

	// Create the PostgreSQL pool
	pool, err := pgxpool.Connect(context.Background(), s.cfg.Postgres.ConnectionString)
	if err != nil {
		log.Fatal("Unable to connect to database:", err)
		return
	}
	defer pool.Close()

	// Create the user storage using the PostgreSQL pool
	userStorage := storage.NewUserStorage(pool)

	// Create the user repository using the user storage
	userRepo := repository.NewUserRepository(userStorage)

	// Create the expense repository using the user storage
	expenseRepo := repository.NewExpenseRepository(userStorage)

	// Create a handler
	grpcHandler := handlers.NewHandler(userRepo, expenseRepo)

	// Running the server
	list, err := net.Listen("tcp", ":"+s.cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer()
	// Register the handler with the gRPC server
	// TODO: Implement gRPC server registration here using grpcServer
	// For example: pb.RegisterYourServiceServer(grpcServer, grpcHandler)

	if err := grpcServer.Serve(list); err != nil {
		log.Fatal("Failed to serve:", err)
	}
}
