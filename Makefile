DB_URL=postgresql://root:abc123@localhost:5432/simple_bank?sslmode=disable
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=abc123 -d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=root --owner=root simple_bank

dockerstart:
	docker start postgres12

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

#migrateup1:
#migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

#migratedown1:
#migrate -path db/migration -database "$(DB_URL)" -verbose down 1

#generate need normal user's permission, not root
sqlc:
	sqlc generate

test:
#-v --> verbose(more details) -cover --> coverage(unit test cover how much code) ./... -> ./all package
	go test -v -cover ./...

server:
	go run main.go

.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server 