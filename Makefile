DB_URL=postgresql://root:secret@localhost:5432/musarchive?sslmode=disable

postgres:
	docker run --name blog_postgres -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres

migrateup1:
	migrate -path db/migration  -database "$(DB_URL)" -verbose up 1

migrateup:
	migrate -path db/migration  -database "$(DB_URL)" -verbose up 

migratedown:
	migrate -path db/migration  -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration  -database "$(DB_URL)" -verbose down 1

new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

sqlc:
	sqlc generate

.PHONY: postgres migrateup1 migrateup migratedown migratedown1 new_migration sqlc
