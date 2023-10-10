package handlers

import (
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters/mock"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/pkg/convert"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"log/slog"
	"net/http"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"
)

func TestAuthHandler_Register(t *testing.T) {
	t.Parallel()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockAuthUC := mock.NewMockAuthUseCase(ctrl)
	mockSessionService := mock.NewMockSessionService(ctrl)

	cfg := &config.Config{
		Session: config.Session{
			Expire: 10,
		},
		Logger: config.Logger{
			Development: true,
		},
	}

	h := AuthHandler{
		AuthUC:         mockAuthUC,
		SessionService: mockSessionService,
		cfg:            cfg,
	}

	user := &entity.User{
		UserName: "UserName",
		Email:    "email@gmail.com",
		Password: "123456",
	}
	buf, err := convert.AnyToBytesBuffer(user)
	require.NoError(t, err)
	require.NotNil(t, buf)
	require.Nil(t, err)

	// Create a new *gin.Context for testing.
	ginContext, _ := gin.CreateTestContext(httptest.NewRecorder())
	ginContext.Request = httptest.NewRequest(http.MethodPost, "/api/v1/auth/register", strings.NewReader(buf.String()))

	handlerFunc := h.Register()
	handlerFunc(ginContext)

	userUID := uuid.New()
	userWithToken := &entity.UserWithToken{
		User: &entity.User{
			UserID: userUID,
		},
	}
	sess := &entity.Session{
		UserID: userUID,
	}
	session := "session"

	mockAuthUC.EXPECT().Register(gomock.Any(), gomock.Eq(user)).Return(userWithToken, nil)
	mockSessionService.EXPECT().CreateSession(gomock.Any(), gomock.Eq(sess), 10).Return(session, nil)

}

func TestAuthHandler_GetUserByID(t *testing.T) {
	type fields struct {
		AuthUC         adapters.AuthUseCase
		SessionService service.SessionService
		cfg            *config.Config
		logger         slog.Logger
	}
	var tests []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AuthHandler{
				AuthUC:         tt.fields.AuthUC,
				SessionService: tt.fields.SessionService,
				cfg:            tt.fields.cfg,
				logger:         tt.fields.logger,
			}
			if got := h.GetUserByID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAuthHandler_Logout(t *testing.T) {
	type fields struct {
		AuthUC         adapters.AuthUseCase
		SessionService adapters.SessionService
		cfg            *config.Config
		logger         slog.Logger
	}
	var tests []struct {
		name   string
		fields fields
		want   gin.HandlerFunc
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := AuthHandler{
				AuthUC:         tt.fields.AuthUC,
				SessionService: tt.fields.SessionService,
				cfg:            tt.fields.cfg,
				logger:         tt.fields.logger,
			}
			if got := h.Logout(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Logout() = %v, want %v", got, tt.want)
			}
		})
	}
}
