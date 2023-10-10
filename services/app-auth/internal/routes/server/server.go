package server

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/gin-gonic/gin"
)

const BaseUrl = "/api"

type Server struct {
	Router      *gin.Engine
	cfg         *config.Config
	authUseCase adapters.AuthUseCase
}

func NewServer(authUseCase adapters.AuthUseCase, cfg *config.Config) *Server {
	r := gin.Default()
	server := &Server{
		Router:      r,
		cfg:         cfg,
		authUseCase: authUseCase,
	}

	// Define your routes here
	authGroup := r.Group(BaseUrl + "/auth")
	authGroup.POST("/login", server.login)
	authGroup.POST("/register", server.register)
	authGroup.POST("/logout", server.logout)

	return server
}

func (serve *Server) login(c *gin.Context) {
	// Your login logic here
}

func (serve *Server) register(c *gin.Context) {
	// Your registration logic here
}

func (serve *Server) logout(c *gin.Context) {
	// Your logout logic here
}

func (serve *Server) Run() error {
	return nil
}
