package repository

import (
	"context"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/dto"
)

//repo implements storage

type authRepo struct {
	dbStorage    adapters.DbStorage
	cacheStorage adapters.CacheStorage
}

func (a authRepo) Register(ctx context.Context, data dto.RegisterReq) (*dto.RegisterRes, error) {
	//TODO implement me
	panic("implement me")
}

func (a authRepo) Login(ctx context.Context, data dto.LoginInitRequest) (*dto.LoginInitResponseData, error) {
	//TODO implement me
	panic("implement me")
}

func NewAuthRepo(dbStorage adapters.DbStorage, cacheStorage adapters.CacheStorage) adapters.AuthRepo {

	return &authRepo{
		dbStorage:    dbStorage,
		cacheStorage: cacheStorage,
	}
}
