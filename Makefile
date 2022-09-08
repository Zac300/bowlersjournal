postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root gqlgen_sqlc_example_db

dropdb:
	docker exec -it postgres12 dropdb gqlgen_sqlc_example_db

migrateup:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gqlgen_sqlc_example_db?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5432/gqlgen_sqlc_example_db?sslmode=disable" -verbose down

sqlc:
	sqlc generate

gqlgen:
	gqlgen

test:
	go clean -testcache && go test -v -cover ./...

server:
	go run main.go

mock:
	mockgen -destination db/mock/store.go -package mockdb github.com/Zac300/simplebank/db/sqlc Store

.PHONNY: postgres createdb dropdb migrateup migratedown sqlc gqlgen test server mock