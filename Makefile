DB_URL=postgres://postgres:postgres@localhost:5432/butterfly?sslmode=disable

run:
	go run main.go

new_migrate:
	migrate create -ext sql -dir db/migration -seq $(name)

migrate_up:
	migrate -path db/migration -database "${DB_URL}" -verbose up

migrate_down:
	migrate -path db/migration -database "${DB_URL}" -verbose down

.PHONY: run new_migrate migrate_up migrate_down