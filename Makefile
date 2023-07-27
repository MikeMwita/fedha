
postgres:
	docker  run --name postgres14 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres14 createdb --username=root --owner=root expenseapp


dropdb:
	docker exec -it postgres14 dropdb expenseapp

migrateup:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/expenseapp?sslmode=disable" -verbose up
migratedown:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/expenseapp?sslmode=disable" -verbose down

#proto:
#	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
#	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
#	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
#    protos/*.proto


proto:
	mkdir -p docs/protos/google/pb
	mkdir -p docs/protos/google/pb_gateway
	protoc --proto_path=docs/protos/google/proto \
		   --go_out=docs/protos/google/pb --go_opt=paths=source_relative \
		   --go-grpc_out=docs/protos/google/pb --go-grpc_opt=paths=source_relative \
		   --grpc-gateway_out=docs/protos/google/pb_gateway --grpc-gateway_opt=paths=source_relative \
		   docs/protos/google/proto/*.proto



#docker exec -it postgres13  createdb --username=root --owner=root fedhaapp
#docker exec -it postgres13 psql -U root fedhaapp


redis:
	docker run --name redis -p 6379:6379 -d redis:7.0.12-alpine



.PHONY:postgres createdb dropdb migrateup migratedown redis proto