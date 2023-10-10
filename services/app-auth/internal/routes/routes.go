package routes

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/routes/handlers"
	"github.com/gin-gonic/gin"
)

type Routes struct {
	*gin.Engine
	authhandler handlers.AuthHandler
}

func NewRoutes(engine *gin.Engine, authhandler handlers.AuthHandler) *Routes {
	r := &Routes{
		Engine:      engine,
		authhandler: authhandler,
	}

	r.setupRoutes()
	return r
}

func (r *Routes) setupRoutes() {
	auth := r.Group("/api/v1/auth")
	auth.POST("/login", r.authhandler.Login())
	auth.POST("/register", r.authhandler.Register())
	auth.POST("/logout", r.authhandler.Logout())
}
