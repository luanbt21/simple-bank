createdb:
	docker exec -it postgres createdb --username appuser simple_bank

dropdb:
	docker exec -it postgres dropdb --username appuser simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://appuser:strongpasswordapp@localhost:5432/simple_bank?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://appuser:strongpasswordapp@localhost:5432/simple_bank?sslmode=disable" -verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

.PHONY: createdb migrateup migratedown sqlc
