package app_auth

import (
	"github.com/gin-gonic/gin"
)

func main() {

	// Create the repo instance
	repo := repo.NewRepo()

	// Create the auth service instance
	authService := auth.NewAuthService(repo)

	// Create the handlers instance
	handlers := handlers.NewHandlers(authService)

	// Create the gin engine
	r := gin.Default()

	// Register the handlers
	handlers.Register(r)

	// Run the server
	r.Run(":5000")
}
