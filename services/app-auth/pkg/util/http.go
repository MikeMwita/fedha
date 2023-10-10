package util

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/pkg/errors"
	"log/slog"
	"net/http"
	"time"
)

// Get request id from Gin context

func GetRequestID(c *gin.Context) string {
	return c.Writer.Header().Get("X-Request-ID")
}

// ReqIDCtxKey is a key used for the Request ID in context
type ReqIDCtxKey struct{}

// Get ctx with timeout and request id from Gin context

func GetCtxWithReqID(c *gin.Context) (context.Context, context.CancelFunc) {
	ctx, cancel := context.WithTimeout(c.Request.Context(), 15*time.Second)
	ctx = context.WithValue(ctx, ReqIDCtxKey{}, GetRequestID(c))
	return ctx, cancel
}

// Get context with request id

func GetRequestCtx(c *gin.Context) context.Context {
	return context.WithValue(c.Request.Context(), ReqIDCtxKey{}, GetRequestID(c))
}

// Get config path for local or docker

func GetConfigPath(configPath string) string {
	if configPath == "docker" {
		return "./config/config-docker"
	}
	return "./config/config-local"
}

//func ConfigureJWTCookie(cfg *config.Config, jwtToken string) *http.Cookie {
//	return &http.Cookie{
//		Name:       cfg.Cookie.Name,
//		Value:      jwtToken,
//		Path:       "/",
//		MaxAge:     int(cfg.Cookie.MaxAge.Seconds()),
//		Secure:     cfg.Cookie.Secure,
//		HttpOnly:   cfg.Cookie.HTTPOnly,
//		SameSite:   http.SameSiteNoneMode,
//	}
//}

func CreateSessionCookie(cfg *config.Config, session string) *http.Cookie {
	return &http.Cookie{
		Name:  cfg.Session.Name,
		Value: session,
		Path:  "/",
		//MaxAge:   int(cfg.Session.Expire()),
		Secure:   cfg.Cookie.Secure,
		HttpOnly: cfg.Cookie.HTTPOnly,
		SameSite: http.SameSiteNoneMode,
	}
}

// Delete session
func DeleteSessionCookie(c *gin.Context, sessionName string) {
	c.SetCookie(sessionName, "", -1, "/", "", false, true)
}

// UserCtxKey is a key used for the User object in the context
type UserCtxKey struct{}

// Get user from context

func GetUserFromCtx(ctx context.Context) (*entity.User, error) {
	user, ok := ctx.Value(UserCtxKey{}).(*entity.User)
	if !ok {
		return nil, errors.New("user not found in context")
	}

	return user, nil
}

// Get user ip address

func GetIPAddress(c *gin.Context) string {
	return c.ClientIP()
}

// ErrResponseWithLog logs the error and returns an error response for Gin context
//func ErrResponseWithLog(ctx *gin.Context, logger sloger.Logger, err error) {
//	logger.Error(
//		"ErrResponseWithLog, RequestID: %s, IPAddress: %s, Error: %s",
//		GetRequestID(ctx),
//		GetIPAddress(ctx),
//		err,
//	)
//	pkg.ErrorResponse(ctx, err)
//}

// LogResponseError logs the error for Gin context
func LogResponseError(ctx *gin.Context, logger slog.Logger, err error) {
	logger.Error(
		"LogResponseError, RequestID: %s, IPAddress: %s, Error: %s",
		GetRequestID(ctx),
		GetIPAddress(ctx),
		err,
	)
}

// Read request body and validate

func ReadRequest(ctx *gin.Context, request interface{}) error {
	if err := ctx.BindJSON(request); err != nil {
		return err
	}
	if err := validator.New().Struct(request); err != nil {
		return err
	}
	return nil
}

// Read sanitize and validate request
//func SanitizeRequest(ctx *gin.Context, request interface{}) error {
//	body, err := ioutil.ReadAll(ctx.Request.Body)
//	if err != nil {
//		return err
//	}
//	defer ctx.Request.Body.Close()
//
//	sanBody, err := sanitize.SanitizeJSON(body)
//	if err != nil {
//		return ctx.AbortWithError(http.StatusBadRequest, err)
//	}
//
//	if err = json.Unmarshal(sanBody, request); err != nil {
//		return err
//	}
//
//	if err = validator.New().Struct(request); err != nil {
//		return err
//	}
//
//	return nil
//}
