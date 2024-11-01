MIGRATE_PATH = ./scripts/migrations
DATABASE_URL = ${DB_SOURCE}

include app.env
export $(shell sed 's/=.*//' app.env)

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
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" up


migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" down


migrate-drop:
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" drop

.PHONY: format vet dockerup dockerdown migrate-up migrate-down migrate-drop