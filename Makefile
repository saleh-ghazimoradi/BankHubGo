
include app.env

MIGRATE_PATH = ./scripts/migrations
DATABASE_URL = ${DB_SOURCE}

format:
	@echo "apply go fmt to the project"
	go fmt ./...

vet:
	@echo "check errors by vet"
	go vet ./...

dockerup:
	@echo "docker up"
	docker compose --env-file app.env up -d

dockerdown:
	@echo "docker down"
	docker compose --env-file app.env down

migrate-up:
	@echo "migrate up"
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" up


migrate-down:
	@echo "migrate down"
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" down


migrate-drop:
	@echo "migrate drop"
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" drop

.PHONY: format vet dockerup dockerdown migrate-up migrate-down migrate-drop