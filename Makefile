DB_URL=postgres://root:secret@localhost:5432/simple_bank?sslmode=disable

postgres:
	docker run --name  postgres --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)"   -verbose up

migrateup1:
	migrate -path db/migration  -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration  -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration  -database "$(DB_URL)" -verbose down 1

migration:
	migrate create -ext sql -dir db/migration -seq init_schema
	
sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go bank/db/sqlc 

db_docs:
	dbdocs build docs/db.dbml
db_schema:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml
proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	--grpc-gateway_out=pb --grpc-gateway_opt=paths=source_relative \
    proto/*.proto

.PHONY: postgres createdb dropdb migration migrateup migratedown migrateup1 migratedown1 sqlc test server mock db_docs db_schema proto