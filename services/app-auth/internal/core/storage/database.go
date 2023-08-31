package storage

import (
	"context"
	"errors"
	"fmt"
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/platform/apps"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"time"
)

var (
	ErrDbDown = errors.New("database service down")
)

type dbStorage struct {
	dbClient db.DbServiceClient
}

func (d dbStorage) RegisterUser(ctx context.Context, in *db.RegUserReq, opts ...grpc.CallOption) (*db.RegUserRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	err := Retry(ctx, func() error {
		_, err := d.dbClient.RegisterUser(ctx, in, opts...)
		if err != nil {
			if grpcStatus, ok := status.FromError(err); ok && grpcStatus.Code() == codes.Unavailable {
				return errors.New("database unavailable")
			}
			return err
		}
		return nil
	})
	if err != nil {
		return nil, fmt.Errorf("failed to register user: %v", err)
	}
	return nil, err
}

func Retry(ctx context.Context, fn func() error) error {
	var err error
	for i := 0; i < 3; i++ {
		err = fn()
		if err == nil {
			return nil
		}
		log.Printf("Failed to register user: %v", err)
		time.Sleep(time.Second)
	}
	return err
}

func (d dbStorage) UpdateUser(ctx context.Context, in *db.UpdateUserReq, opts ...grpc.CallOption) (*db.UpdateUserRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.UpdateUser(ctx, in, opts...)
}

func (d dbStorage) GetPagedUsers(ctx context.Context, in *db.GetPagedUsersReq, opts ...grpc.CallOption) (*db.GetPagedUsersRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.GetPagedUsers(ctx, in, opts...)
}

func (d dbStorage) GetUserByField(ctx context.Context, in *db.GetByfieldReq, opts ...grpc.CallOption) (*db.GetByfieldRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.GetUserByField(ctx, in, opts...)
}

func (d dbStorage) GetUserByUsername(ctx context.Context, in *db.GetUserByUsernameRequest, opts ...grpc.CallOption) (*db.RegUserRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.GetUserByUsername(ctx, in, opts...)
}

func (d dbStorage) GetUserByID(ctx context.Context, in *db.GetUserByIDRequest, opts ...grpc.CallOption) (*db.RegUserRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.GetUserByID(ctx, in, opts...)
}

func (d dbStorage) SaveUser(ctx context.Context, in *db.SaveUserRequest, opts ...grpc.CallOption) (*db.User, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.SaveUser(ctx, in, opts...)
}
func (d dbStorage) FindByUsername(ctx context.Context, username string) (*db.RegUserRes, error) {
	if d.dbClient == nil {
		return nil, ErrDbDown
	}
	return d.dbClient.GetUserByUsername(ctx, &db.GetUserByUsernameRequest{Username: username})
}

func NewDbStorage(serviceCfg config.DatabaseService) (adapters.DbStorage, error) {
	client, err := apps.NewDBServiceClient(serviceCfg)
	if err != nil {
		return nil, err
	}
	return &dbStorage{
		dbClient: client,
	}, nil
}
