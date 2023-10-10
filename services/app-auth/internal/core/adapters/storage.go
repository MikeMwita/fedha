package adapters

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"

	"google.golang.org/grpc"
	"time"
)

type DbStorage interface {
	RegisterUser(ctx context.Context, in *db.RegUserReq, opts ...grpc.CallOption) (*db.RegUserRes, error)
	UpdateUser(ctx context.Context, in *db.UpdateUserReq, opts ...grpc.CallOption) (*db.UpdateUserRes, error)
	GetPagedUsers(ctx context.Context, in *db.GetPagedUsersReq, opts ...grpc.CallOption) (*db.GetPagedUsersRes, error)
	GetUserByField(ctx context.Context, in *db.GetByfieldReq, opts ...grpc.CallOption) (*db.GetByfieldRes, error)
	GetUserByUsername(ctx context.Context, in *db.GetUserByUsernameRequest, opts ...grpc.CallOption) (*db.RegUserRes, error)
	GetUserByID(ctx context.Context, in *db.GetUserByIDRequest, opts ...grpc.CallOption) (*db.RegUserRes, error)
	SaveUser(ctx context.Context, in *db.SaveUserRequest, opts ...grpc.CallOption) (*db.User, error)
	Update(ctx context.Context, i *db.UpdateUserReq) (*db.UpdateUserRes, error)
	GetByID(ctx context.Context, in *db.GetUserByIDRequest, opts ...grpc.CallOption) (*db.RegUserRes, error)
}

type CacheStorage interface {
	SetAccessToken(ctx context.Context, key string, value string, expiration time.Duration) error
	GetAccessToken(ctx context.Context, value string) (string, error)
	DeleteAccessToken(ctx context.Context, key string) error
	DeleteSession(ctx context.Context, id string) error
}
