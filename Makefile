
postgres:
	docker  run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root expenseapp


dropdb:
	docker exec -it postgres14 dropdb expenseapp

migrateup:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/expenseapp?sslmode=disable" -verbose up
migratedown:
	migrate --path db/migration -database "postgresql://root:secret@localhost:5432/expenseapp?sslmode=disable" -verbose down


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

# ==============================================================================
# Docker support

docker-run:
	 docker run --name fedhaapp -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@172.20.0.2:5432/fedhaapi?sslmode=disable" fedhaapp

main-docker-run:
	 docker run --name fedhaapp --network fedha-net  -p 8080:8080 -e GIN_MODE=release -e DB_SOURCE="postgresql://root:secret@ fedha-postgres-1:5432/fedhaapi?sslmode=disable" fedhaapp

#docker build :
#       docker build -t fedhaapp .
#docker run :
#	   docker run -it -p 8080:8080 fedhaapp

create-network:
	docker network create fedha-net
connect-net:
	docker network connect fedha-net fedha-postgres-1


#docker run -d -p 8080:8080 fedha-api

.PHONY:postgres createdb dropdb migrateup migratedown redis proto