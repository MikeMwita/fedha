package middleware

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/routes/handlers"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"net/http"
	"strings"
)

////token middleware
//
//func tokenAuthMiddleware(next http.Handler) http.Handler {
//	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
//		// Get the token from the request header
//		token := r.Header.Get("Authorization")
//
//		// Check if the token is valid
//		if token == "" {
//			w.WriteHeader(http.StatusUnauthorized)
//			return
//		}
//
//		// The token is valid, so continue with the request
//		next.ServeHTTP(w, r)
//	})
//}

const (
	AuthorisationHeader       = "Authorization"
	AuthorizationBearerPrefix = "Bearer"
	UserUUIDKey               = "UserUUID"
)

func unauthorisedError() dto.DefaultRes[any] {
	return dto.DefaultRes[any]{
		Message: "request failed: unauthorised",
		Error:   "request not authorised: missing a valid token",
		Code:    http.StatusUnauthorized,
		Data:    nil,
	}
}

func (m *MiddlewareManager) Auth(ctx *gin.Context) {
	if ctx.IsAborted() {
		return
	}

	if _, exists := ctx.Get(handlers.BearerScopes); exists {
		// generate default error when unauthorised
		unauthorisedResponse := unauthorisedError()

		// get token from header
		bearerToken := ctx.GetHeader(AuthorisationHeader)
		bearerToken = strings.TrimPrefix(bearerToken, AuthorizationBearerPrefix)
		bearerToken = strings.TrimSpace(bearerToken)

		if bearerToken == "" {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorisedResponse)
			return
		}
		// validate token
		userId, err := m.AuthUsecase.VerifyAccessToken(bearerToken)
		if err != nil {
			log.Error(err)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, unauthorisedResponse)
			return
		}

		ctx.Set(UserUUIDKey, userId)
	}
	ctx.Next()
}
