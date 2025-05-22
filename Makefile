# PostgreSQL configuration (Homebrew)
PG_VERSION := 15
CONNECTION_STRING := "postgres://jamieholliday:@localhost:5432/gator"

.PHONY: db-start db-stop db-status db-restart db-help db-generate db-migrate-up db-migrate-down

db-start:
	@echo "Starting PostgreSQL service..."
	brew services start postgresql@$(PG_VERSION)

db-stop:
	@echo "Stopping PostgreSQL service..."
	brew services stop postgresql@$(PG_VERSION)

db-status:
	@echo "PostgreSQL service status:"
	brew services info postgresql@$(PG_VERSION)

db-restart:
	@echo "Restarting PostgreSQL service..."
	brew services restart postgresql@$(PG_VERSION)

db-generate:
	@echo "Generating sqlc code"
	sqlc generate

db-migrate-up:
	@echo "Running up migration"
	cd sql/schema; goose postgres $(CONNECTION_STRING) up

db-migrate-down:
	@echo "Running up migration"
	cd sql/schema; goose postgres $(CONNECTION_STRING) down
