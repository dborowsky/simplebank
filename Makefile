postgres:
	docker run --name simplebank -p 6060:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:14-alpine
createdb:
	docker exec -it simplebank createdb --username=root --owner=root simple_bank
dropdb:
	docker exec -it simplebank dropdb --username=root simple_bank
migrateup:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:6060/simple_bank?sslmode=disable" -verbose up
migrateup1:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:6060/simple_bank?sslmode=disable" -verbose up 1	
migratedown:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:6060/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migrate -database "postgresql://root:secret@localhost:6060/simple_bank?sslmode=disable" -verbose down 1	
sqlc:
	sqlc generate
test:
	go test -v -cover ./...
server:
	go run main.go
mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/dborowsky/simplebank/db/ Store
.PHONY: postgres createdb dropdb migrateup migratedown sqlc server mock