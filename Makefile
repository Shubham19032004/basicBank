DB_URL=postgres://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name  postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migration -path db/migration  -database "$(DB_URL)" -verbose up

migrateup1:
	migration -path db/migration  -database "$(DB_URL)" -verbose up 1

migratedown:
	migration -path db/migration  -database "$(DB_URL)" -verbose down

migratedown1:
	migration -path db/migration  -database "$(DB_URL)" -verbose down 1

migration:
	migration create -ext sql -dir db/migration -seq init_schema
	
sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go bank/db/sqlc Store

.PHONY: postgres createdb dropdb migration migrateup migratedown migrateup1 migratedown1 sqlc test server mock 