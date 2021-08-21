postgres:
	docker run --name postgres13 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it postgres13 createdb --username=root --owner=root fibonacci
createinitmigrate:
	migrate create -ext sql -dir api-multi/db/pg/migration -seq init_schema
dropdb:
	docker exec -it postgres13 dropdb simple_bank
migrateup:
	migrate -path api-multi/db/pg/migration -database "postgresql://root:secret@localhost:5432/fibonacci?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fibonacci?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fibonacci?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/fibonacci?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run api-multi/main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/thanhftu/simple_bank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup sqlc server createinitmigrate mock