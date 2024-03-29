postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=rahasia123 -d postgres:alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:rahasia123@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:rahasia123@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:rahasia123@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

migrate1: #if occur error dirty version 1, use this!!!
	migrate -path db/migration -database "postgresql://root:rahasia123@localhost:5432/simple_bank?sslmode=disable" force 1

sqlc:
	sqlc generate

test:
	go test -v ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Federico191/Simple_Bank/db/sqlc Store
.PHONY: postgres createdb dropdb migrateup migratedown migrate1 sqlc server mock