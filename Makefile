postgres:
	docker run --name udemy -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createdb:
	docker exec -it udemy createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it udemy dropdb --username=root simple_bank
migrateup:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/dborowsky/simplebank/db/ Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock