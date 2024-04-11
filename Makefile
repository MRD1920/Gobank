postgres:
	docker run --name postgrs12 -p 5433:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
createdb:
	docker exec -it postgres12 createdb --username=root  --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
migrateup: 
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migration -database "postgresql://root:secret@localhost:5433/simple_bank?sslmode=disable" -verbose down
sqlc:
	sqlc generate
test: 
	go test -v -cover ./...
	


.PHONY: postgres createdb dropdb sqlc test