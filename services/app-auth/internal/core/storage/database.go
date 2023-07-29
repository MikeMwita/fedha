package storage

import (
	"context"
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"google.golang.org/grpc"
)

type dbStorage struct {
	dbClient db.DbServiceClient
}

func (d dbStorage) CreateUser(ctx context.Context, in *db.RegUserReq, opts ...grpc.CallOption) (*db.RegUserRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) UpdateUser(ctx context.Context, in *db.UpdateUserReq, opts ...grpc.CallOption) (*db.UpdateUserRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) GetPagedUsers(ctx context.Context, in *db.GetPagedUsersReq, opts ...grpc.CallOption) (*db.GetPagedUsersRes, error) {
	//TODO implement me
	panic("implement me")
}

func (d dbStorage) GetUserByField(ctx context.Context, in *db.GetByfieldReq, opts ...grpc.CallOption) (*db.GetByfieldRes, error) {
	//TODO implement me
	panic("implement me")
}

func NewDbStorage(serviceCfg config.DatabaseService) (adapters.DbStorage, errror) {
	client, err := apps.NewDbStorageClient(serviceCfg)
	if err != nil {
		return nil, err
	}
	return &dbStorage{
		dbClient: client,
	}, nil
}
