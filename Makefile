DB_URL=postgresql://root:postgres@localhost:5432/simple_cat?sslmode=disable

postgres: 
	docker run --name postgres17 --network bank-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=postgres -d postgres:17-alpine

createdb:
	docker exec -it postgres17 createdb --username=root --owner=root simple_cat

dropdb:
	docker exec -it postgres17 dropdb simple_cat

migrateup:
	migrate --path db/migration --database "$(DB_URL)" --verbose up

migrateup1:
	migrate --path db/migration --database "$(DB_URL)" --verbose up 1

migratedown:
	migrate --path db/migration --database "$(DB_URL)" --verbose down

migratedown1:
	migrate --path db/migration --database "$(DB_URL)" --verbose down 1	

sqlc:
	sqlc generate

.PHONY: postgres createdb dropdb migrateup migrateup1 migratedown migratedown1 sqlc