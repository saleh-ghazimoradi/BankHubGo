MIGRATE_PATH = ./scripts/migrations
DATABASE_URL = ${DB_SOURCE}


format:
	@echo "apply go fmt to the project"
	go fmt ./...

vet:
	@echo "check errors by vet"
	go vet ./...

dockerup:
	docker compose --env-file app.env up -d

dockerdown:
	docker compose --env-file app.env down

migrate-up:
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" up


migrate-down:
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" down


migrate-drop:
	migrate -path $(MIGRATE_PATH) -database "$(DATABASE_URL)" drop

http:
	go run . http


.PHONY: format vet dockerup dockerdown migrate-up migrate-down migrate-drop http