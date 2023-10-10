package apps

import (
	"github.com/MikeMwita/fedha-go-gen.grpc/sdk/go-proto-gen/db"
	"github.com/MikeMwita/fedha.git/services/app-auth/config"
	"github.com/MikeMwita/fedha.git/services/app-auth/internal/core/adapters"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	_ adapters.DbStorage = (db.DbServiceClient)(nil)
)

func NewDBServiceClient(config config.Database) (db.DbServiceClient, error) {
	conn, err := grpc.Dial(config.Host+":"+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	log.Info("connected to app-db")
	client := db.NewDbServiceClient(conn)
	return client, nil
}
