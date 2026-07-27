.PHONY: run
run:
	@go run ./cmd/service/main.go

.PHONY: build
build:
	@go build -o ./temp/bin/service ./cmd/service/main.go

.PHONY: dev
dev: build
	@./temp/bin/service

.PHONY: docker/up
docker/up:
	@docker compose -f ./deploy/compose.local.yaml up -d $(service)

.PHONY: docker/stop
docker/stop:
	@docker compose -f ./deploy/compose.local.yaml stop

.PHONY: docker/down
docker/down:
	@docker compose -f ./deploy/compose.local.yaml down
