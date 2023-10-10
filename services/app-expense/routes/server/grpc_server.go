package server

import (
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/expense"
	"github.com/MikeMwita/fedha.git/services/app-expense/config"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/repositories"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/services"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/core/storage"
	"github.com/MikeMwita/fedha.git/services/app-expense/internal/platform"
	"github.com/MikeMwita/fedha.git/services/app-expense/routes/handlers"
	"github.com/labstack/gommon/log"
	"google.golang.org/grpc"
	"net"
	"os"
	"os/signal"
	"syscall"
)

type Server struct {
	cfg config.Config
}

func (s *Server) Run() {
	log.Infof("DATABASE GRPC serve initialising")

	// create database client
	postgresClient, err := platform.NewDBServiceClient(s.cfg.DB)
	if err != nil {
		log.Panic("error:", err)
	}

	// Create User storage
	dbStorage := storage.NewDbStorage(postgresClient)

	// create  repository
	balanceRepo:=repositories.NewBalanceRepository(dbStorage)
	expenseRepo:=repositories.NewExpenseRepository(dbStorage)
	incomeRepo:=repositories.NewIncomeRepository(dbStorage)
	monthly_summaryRepo:=repositories.NewMonthlyRepo(dbStorage)


	balanceservice := services.NewExpenseService(balanceRepo)
	expenseService := services.NewExpenseService(expenseRepo)
	incomeService := services.NewIncomeService(incomeRepo)
	monthly_summaryService := services.NewMonthlyService(monthly_summaryRepo)



	// create handler
	grpcHandler := handlers.NewHandler(balanceservice,expenseService,incomeService,monthly_summaryService

	// run serve
	lis, err := net.Listen("tcp", ":"+s.cfg.Server.Port)
	if err != nil {
		log.Fatal(err)
		return
	}

	grpcServer := grpc.NewServer()
	expense.RegisterExpenseServiceServer(grpcServer, grpcHandler)

	// Run the serve

	// wait for interrupt signal to gracefully shutdown the serve with
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			log.Fatal("cannot start apps serve:", err)
			return
		}
	}()

	<-quit

}

func NewServer(cfg config.Config) *Server {
	return &Server{
		cfg: cfg,
	}
}
