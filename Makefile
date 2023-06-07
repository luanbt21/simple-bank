db:
	docker compose up -d

createdb:
	docker exec -it simplebank-db-1 createdb --username root simple_bank

dropdb:
	docker exec -it simplebank-db-1 dropdb --username root simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run .

mock:
	go generate ./...

.PHONY: db createdb migrateup migratedown sqlc mock
