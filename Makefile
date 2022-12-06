dsn=mysql://root:root@tcp(localhost:3306)/reading-tracker?parseTime=true

.PHONY: migrate-create
migrate-create:
	@echo "Creating migration..."
	@migrate create -ext sql -dir migrations -seq $(name)
	@echo "Migration created."

.PHONY: migrate-up
migrate-up:
	@echo "Running migrations..."
	@migrate -path migrations -database $(dsn) up
	@echo "Migrations ran."

.PHONY: migrate-down
migrate-down:
	@echo "Rolling back migrations..."
	@migrate -path migrations -database $(dsn) down
	@echo "Migrations rolled back."