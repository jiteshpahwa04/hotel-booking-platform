MIGRATIONS_FOLDER=db/migrations
DB_URL=root:my-password@tcp(localhost:3306)/auth

# Create a new migration file - gmake migrate-create name="migration_name" (make instead of gmake on Windows)
migrate-create:
	goose -dir $(MIGRATIONS_FOLDER) create $(name) sql

migrate-up:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) up

migrate-down:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) down

migrate-status:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) status

migrate-reset:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) reset

migreate-redo:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) redo

migrate-to:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) up-to $(version)

migrate-down-to:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) down-to $(version)

migrate-force:
	goose -dir $(MIGRATIONS_FOLDER) mysql $(DB_URL) force $(version)

migrate-help:
	goose -h