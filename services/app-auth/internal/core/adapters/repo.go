package adapters

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/entity"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/service"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"time"
)

//type AuthRepo interface {
//	Register(ctx context.Context, registerReq dto.RegisterRequest) (*dto.RegisterResponseData, error)
//	Login(ctx context.Context, loginReq dto.LoginRequest) (*dto.LoginResponseData, error)
//	//Login(ctx context.Context,loginReq dto.LoginInitRequest) (*dto.LoginInitResponseData, error)
//	GetUserByID(ctx context.Context, in *db.GetUserByIDRequest, opts ...grpc.CallOption) (*db.RegUserRes, error)
//	RefreshToken(c *gin.Context, data dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
//	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
//	UserLogout(c *gin.Context)
//	SetAccessToken(ctx context.Context, key string, value string, expiration time.Duration) error
//	GetAccessToken(ctx context.Context, value string) (string, error)
//	DeleteAccessToken(ctx context.Context, key string) error
//	FindByUsername(ctx context.Context, username string) (*db.RegUserRes, error)
//	SaveUser(ctx context.Context, in *db.SaveUserRequest, opts ...grpc.CallOption) (*db.User, error)
//	DeleteSession(ctx context.Context, record *service.SessionRecord) error
//	InvalidateSession() error
//	VerifyRefreshToken(token string) error
//}

//type AuthRepository interface {
//	Register(ctx context.Context, user *entity.User) (*entity.UserWithToken, error)
//	Login(ctx context.Context, user *entity.User) (*entity.UserWithToken, error)
//	Update(ctx context.Context, user *entity.User) (*entity.User, error)
//	Delete(ctx context.Context, userID uuid.UUID) error
//	GetByID(ctx context.Context, userID uuid.UUID) (*entity.User, error)
//	FindByEmail(ctx context.Context, user *entity.User) (*entity.User, error)
//}

type AuthRepository interface {
	Register(ctx context.Context, registerReq dto.RegisterRequest) (*dto.RegisterResponseData, error)
	Login(ctx context.Context, loginReq dto.LoginRequest) (*dto.LoginResponseData, error)
	UpdateUser(ctx context.Context, user entity.User) (*entity.User, error)
	GetUserByID(ctx context.Context, in *db.GetUserByIDRequest, opts ...grpc.CallOption) (*db.RegUserRes, error)
	FindByEmail(ctx context.Context, user *entity.User) (*entity.User, error)

	RefreshToken(c *gin.Context, data dto.RefreshTokenRequest) (*dto.RefreshTokenResponse, error)
	UserLogout(c *gin.Context)
	SetAccessToken(ctx context.Context, key string, value string, expiration time.Duration) error
	GetAccessToken(ctx context.Context, value string) (string, error)
	DeleteAccessToken(ctx context.Context, key string) error
	FindByUsername(ctx context.Context, username string) (*db.RegUserRes, error)
	SaveUser(ctx context.Context, in *db.SaveUserRequest, opts ...grpc.CallOption) (*db.User, error)
	DeleteSession(ctx context.Context, record *service.SessionRecord) error
	InvalidateSession() error
	VerifyRefreshToken(token string) error
	Delete(ctx context.Context, id uuid.UUID) error
	GetByID(ctx context.Context, id uuid.UUID) (*db.RegUserRes, error)
	Update(ctx context.Context, userReq *db.UpdateUserReq) (*db.UpdateUserRes, error)
}
