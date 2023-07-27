package storage

import (
	"github.com/MikeMwita/fedha-go-gen.grpc/generated_rpc_code/github.com/MikeMwita/fedha-go-gen.grpc/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core /adapters"
)

type dbStorage struct {
	dbClient db.DbServiceClient
}


func NewDbStorage(serviceCfg config.DatabaseService)(adapters.DbStorage,errror)  {
	client,err:=apps.NewDbStorageClient(serviceCfg)
	if err!=nil{
		return &dbStorage{dbClient: client,}
	},nil
}