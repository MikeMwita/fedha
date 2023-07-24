# migrate --path db/migration -database "postgresql://root:secret@localhost:5432/expenseapp?sslmode=disable" -verbose up




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

.PHONY:postgres createdb dropdb migrateup migratedown



#docker exec -it postgres13  createdb --username=root --owner=root fedhaapp
#docker exec -it postgres13 psql -U root fedhaapp


