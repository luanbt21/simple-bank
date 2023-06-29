db:
	docker compose up -d

createdb:
	docker exec -it simplebank-db-1 createdb --username root simple_bank

dropdb:
	docker exec -it simplebank-db-1 dropdb --username root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migrateup1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up 1

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down 1

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run .

mock:
	go generate ./...

.PHONY: db createdb migrateup migratedown migrateup1 migratedown1 sqlc mock
