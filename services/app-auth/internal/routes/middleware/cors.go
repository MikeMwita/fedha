package middleware

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/models"
	"github.com/gin-gonic/gin"
)

func (m middlewareManager) Cors(ctx *gin.Context) {
	var allowedOrigins = []string{
		"http://localhost",
		"https://www.fedha.co.ke",
		"https://auth.fedha.co.ke",
	}

	origin := ctx.GetHeader("Origin")
	isAllowed := false

	for _, allowedOrigin := range allowedOrigins {
		if origin == allowedOrigin {
			isAllowed = true
			break
		}
	}

	if isAllowed {
		ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept, Authorization")
		ctx.Header("Access-Control-Allow-Credentials", "true")
		ctx.Header("Access-Control-Allow-Origin", origin)

		// handle preflight request
		if ctx.Request.Method == "OPTIONS" {
			ctx.AbortWithStatus(204)
			return
		}

		ctx.Next()

	} else {
		ctx.AbortWithStatus(403)
		ctx.JSON(403, models.HttpResponse{
			Success: false,
			Message: "Forbidden",
			Data:    nil,
			Errors:  []string{"oops! you are not allowed to access this resource"},
		})
		return
	}
}
