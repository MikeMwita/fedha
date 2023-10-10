package handlers

import (
	"errors"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/util"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"net/http"
)

const (
	BearerScopes = "Bearer.Scopes"
)

type AuthHandler struct {
	AuthUC         adapters.AuthUseCase
	SessionService adapters.SessionService
	cfg            *config.Config
	logger         slog.Logger
	tracer         trace.Tracer
}

func (h AuthHandler) Register() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := h.tracer.Start(c.Request.Context(), "auth.Register")
		defer span.End()

		user := &entity.User{}
		if err := c.BindJSON(user); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			util.LogResponseError(c, h.logger, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		createdUser, err := h.AuthUC.Register(ctx, user)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			util.LogResponseError(c, h.logger, err)
			return
		}
		sess, err := h.SessionService.CreateSession(ctx, &entity.Session{
			UserID: createdUser.User.UserID,
		}, h.cfg.Session.Expire)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())
			util.LogResponseError(c, h.logger, err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		cookie := util.CreateSessionCookie(h.cfg, sess)
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		c.JSON(http.StatusCreated, createdUser)
	}
}
func (h *AuthHandler) Login() gin.HandlerFunc {
	type Login struct {
		Email    string `json:"email" db:"email" validate:"omitempty,lte=60,email"`
		Password string `json:"password,omitempty" db:"password" validate:"required,gte=6"`
	}

	return func(c *gin.Context) {
		ctx, span := h.tracer.Start(c.Request.Context(), "auth.Login")
		defer span.End()

		login := &Login{}
		if err := util.ReadRequest(c, login); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusBadRequest})
			return
		}

		userWithToken, err := h.AuthUC.Login(ctx, &entity.User{
			Email:    login.Email,
			Password: login.Password,
		})
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			c.JSON(http.StatusUnauthorized, gin.H{"message": "unauthorized", "status_code": http.StatusUnauthorized})
			return
		}

		sess, err := h.SessionService.CreateSession(ctx, &entity.Session{
			UserID: userWithToken.User.UserID,
		}, h.cfg.Session.Expire)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "status_code": http.StatusInternalServerError})
			return
		}

		cookie := util.CreateSessionCookie(h.cfg, sess)
		c.SetCookie(cookie.Name, cookie.Value, cookie.MaxAge, cookie.Path, cookie.Domain, cookie.Secure, cookie.HttpOnly)
		c.JSON(http.StatusOK, gin.H{"message": "login success", "status_code": http.StatusOK, "data": userWithToken})
	}
}

func (h AuthHandler) Logout() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := h.tracer.Start(c.Request.Context(), "auth.Logout")
		defer span.End()

		cookie, err := c.Cookie("session-id")
		if err != nil {
			if errors.Is(err, http.ErrNoCookie) {
				span.RecordError(err)
				span.SetStatus(codes.Error, err.Error())

				c.JSON(http.StatusUnauthorized, gin.H{"message": err.Error() + " /logout", "data": nil, "status_code": http.StatusUnauthorized})
				return
			}

			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusInternalServerError})
			return
		}

		if err := h.SessionService.DeleteByID(ctx, cookie); err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusInternalServerError})
			return
		}

		util.DeleteSessionCookie(c, h.cfg.Session.Name)

		c.JSON(http.StatusOK, gin.H{"message": "success", "status_code": http.StatusOK})
	}
}

func (h AuthHandler) GetUserByID() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, span := h.tracer.Start(c.Request.Context(), "auth.GetUserByID")
		defer span.End()

		uID, err := uuid.Parse(c.Param("user_id"))
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			c.JSON(http.StatusBadRequest, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusBadRequest})
			return
		}

		user, err := h.AuthUC.GetByID(ctx, uID)
		if err != nil {
			span.RecordError(err)
			span.SetStatus(codes.Error, err.Error())

			c.JSON(http.StatusNotFound, gin.H{"message": err.Error(), "data": nil, "status_code": http.StatusNotFound})
			return
		}

		c.JSON(http.StatusOK, gin.H{"message": "user found", "data": user, "status_code": http.StatusOK})
	}
}

func NewHandler(authUC adapters.AuthUseCase, SessionService service.SessionService, cfg *config.Config, logger slog.Logger, tracer trace.Tracer) adapters.AuthHandler {
	return &AuthHandler{
		AuthUC:         authUC,
		SessionService: SessionService,
		cfg:            cfg,
		logger:         logger,
		tracer:         tracer,
	}
}
