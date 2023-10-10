package apps

//
//var (
//	_ adapters.DbStorage = (db.DbServiceClient)(nil)
//)
//
//func NewDBServiceClient(config config.Database) (db.DbServiceClient, error) {
//	conn, err := grpc.Dial(config.Host+":"+config.Port, grpc.WithTransportCredentials(insecure.NewCredentials()))
//	if err != nil {
//		return nil, err
//	}
//
//	log.Info("connected to app-db")
//	client := db.NewDbServiceClient(conn)
//	return client, nil
//}
