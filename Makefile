MIGRATION_DIR = "./migration"
DB_DRIVER = "postgres"
DB_STRING = "postgres://igortryhan:postgres@localhost:5432/logs_storage?sslmode=disable"

.PHONY: run docker-build docker-run migrate-create migrate-up migrate-down

run:
	go run ./cmd/logsStorage/main.go

docker-build:
	docker build -t logs-storage-service:0.0.1 -f Dockerfile .

docker-run:
	docker run --rm --name logs-storage-service -p 8081:8081 logs-storage-service:0.0.1

#make migrate-create name=your_migration_name
migrate-create:
	goose -dir ${MIGRATION_DIR} create $(name) sql

migrate-up:
	goose -dir ${MIGRATION_DIR} ${DB_DRIVER} ${DB_STRING} up

migrate-down:
	goose -dir ${MIGRATION_DIR} ${DB_DRIVER} ${DB_STRING} down
