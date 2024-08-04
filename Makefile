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

test:
	go test -v -cover -short ./...

sqlc:
	sqlc generate

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/mustafayilmazdev/musarchive/db/sqlc Store
	mockgen -package mockdb -destination worker/mock/distributor.go github.com/mustafayilmazdev/musarchive/worker TaskDistributor

locale:
	musale --json=locales/assets/en.json --output=locales/localekeys.go -p=localization

redis:
	docker run --name redis -p 6379:6379 -d redis:7.4-alpine

.PHONY: postgres migrateup1 migrateup migratedown migratedown1 new_migration test sqlc server mock locale redis
