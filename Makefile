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

.PHONY: format vet dockerup dockerdown