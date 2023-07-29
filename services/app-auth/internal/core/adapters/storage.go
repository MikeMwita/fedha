package adapters

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"google.golang.org/grpc"
)

type DbStorage interface {
	//methods for interacting with the database
	CreateUser(ctx context.Context, in *db.RegUserReq, opts ...grpc.CallOption) (*db.RegUserRes, error)
	UpdateUser(ctx context.Context, in *db.UpdateUserReq, opts ...grpc.CallOption) (*db.UpdateUserRes, error)
	GetPagedUsers(ctx context.Context, in *db.GetPagedUsersReq, opts ...grpc.CallOption) (*db.GetPagedUsersRes, error)
	GetUserByField(ctx context.Context, in *db.GetByfieldReq, opts ...grpc.CallOption) (*db.GetByfieldRes, error)
}

type CacheStorage interface {
}
