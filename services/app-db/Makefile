
postgres:
	docker run --name pesaapp -p 5432:5432 -e POSTGRES_USER=mike -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it pesaapp  createdb --username=mike  --owner=mike fedhausers

login:
	docker exec -it pesaapp  psql -U mike

sqlc-generate:
	sqlc generate



.PHONY:postgres createdb login sqlc-generate