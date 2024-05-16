postgres:
	docker run --name postgrs12 -p 5433:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
createdb:
	docker exec -it postgres12 createdb --username=root  --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up
migrateup1: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up 1
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down
migratedown1:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down 1
sqlc:
	sqlc generate
test: 
	go test -v -cover ./...
server:
	go run main.go
	


.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown sqlc test server 