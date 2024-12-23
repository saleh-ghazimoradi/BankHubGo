ifneq (,$(wildcard ./app.env))
	include app.env
	export $(shell sed 's/=.*//' app.env)
endif

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

migrate-create:
	@echo "Creating migration files for ${name}..."
	migrate create -seq -ext=.sql -dir=./scripts/migrations ${name}

migrate-up:
	@echo "Running up migrations..."
	migrate -path ${MIGRATE_PATH} -database "${DATABASE_URL}" up

migrate-down:
	@echo "Rolling back migrations..."
	@if [ -z "$(n)" ]; then \
		migrate -path ${MIGRATE_PATH} -database "${DATABASE_URL}" down 1; \
	else \
		migrate -path ${MIGRATE_PATH} -database "${DATABASE_URL}" down $(n); \
	fi


migrate-drop:
	@echo "Dropping all migrations..."
	migrate -path ${MIGRATE_PATH} -database "${DATABASE_URL}" drop -f

http:
	go run . http

test:
	go test -v -cover ./...


.PHONY: format vet dockerup dockerdown migrate-up migrate-down migrate-drop http test