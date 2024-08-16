postgres:
	docker run --name postgrs12 -p 5433:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine
createdb:
	docker exec -it postgres12 createdb --username=root  --owner=root simple_bank
dropdb:
	docker exec -it postgres12 dropdb simple_bank
migrateup: 
	migrate -path db/migration -database "postgresql://mrd:mIc%3C%3E9cAA%23x%7C7a%25IsqYo%7B%3ELxCHkP@go-bank.c18ms0iyacu7.ap-south-1.rds.amazonaws.com:5432/go_bank" -verbose up
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
proto:
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
    --go-grpc_out=pb --go-grpc_opt=paths=source_relative \
    proto/*.proto


.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown sqlc test server proto                            