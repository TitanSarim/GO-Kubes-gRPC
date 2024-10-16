# Makefile

DB_URL=postgresql://postgres:12345@localhost:5432/simple_bank?sslmode=disable
MIGRATE_PATH=db/migration 

run:
	@echo "Running the Go project..."
	DB_URL=$(DB_URL) go run main.go

# Define a migrate target for running migrations (up)
migrate-up:
	@echo "Running all migrations up..."
	migrate -path db/migration -database "$(DB_URL)" -verbose up

# Define a migrate target for rolling back all migrations (down)
migrate-down:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

# Define a target to migrate up by 1 step
migrate-up1:
	@echo "Running 1 migration up..."
	migrate -path db/migration -database "$(DB_URL)" -verbose up

# Define a target to migrate down by 1 step
migrate-down1:
	@echo "Rolling back 1 migration down..."
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

# Create a new migration file (requires a name as an argument: make new-migration name=<migration_name>)
new-migration:
	@echo "Creating new migration named '$(name)'..."
	migrate create -ext sql -dir db/migration -seq $(name)

test:
	go test -v -cover -short ./...

.PHONY: run migrate-up migrate-down migrate-up1 migrate-down1 new-migration test
