postgres:
	docker run --name  postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

createdb:
	docker exec -it postgres createdb --username=root --owner=root simple_bank

dropdb:
	docker exec -it postgres dropdb simple_bank

db_up:
	migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

db_up1:
	migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1



db_down:
	migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

db_down1:
	migrate -path db/migrate -database "postgres://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1


migrate:
	migrate create -ext sql -dir db/migrate -seq init_schema
	
sqlc:
	sqlc generate

server:
	go run main.go

test:
	go test -v -cover ./...

mock:
	mockgen -package mockdb -destination db/mock/store.go bank/db/sqlc Store

.PHONY: postgres createdb dropdb migrate db_up db_down sqlc test server mock db_down1