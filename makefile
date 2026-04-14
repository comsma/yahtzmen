## db/init: create database and run all migrations via yahtzmen init
db/init:
	go run ./cmd/yahtzmen init

## db/migrations/new name=<name>: create a new goose migration file
db/migrations/new:
	@echo "Creating migration files for ${name}..."
	goose -dir ./migrations create ${name} sql

dev/dependencies/up:
	podman kube play deployment/dev/database.yml

dev/dependencies/down:
	podman kube down deployment/dev/database.yml

db/seed:
	go run ./cmd/yahtzmen seed

dev/run:
	go run ./cmd/yahtzmen serve

## test: run all tests with Podman as the container provider
test:
	DOCKER_HOST=unix://$$(podman machine inspect --format '{{.ConnectionInfo.PodmanSocket.Path}}') \
	TESTCONTAINERS_RYUK_DISABLED=true \
	go test ./... -v -timeout 120s

test/web:
	DOCKER_HOST=unix://$$(podman machine inspect --format '{{.ConnectionInfo.PodmanSocket.Path}}') \
	TESTCONTAINERS_RYUK_DISABLED=true \
	go test ./internal/web/... -v -timeout 120s

## sqlc/generate: regenerate models from queries
sqlc/generate:
	sqlc generate

## build: build the yahtzmen binary
build:
	go build -o bin/yahtzmen ./cmd/yahtzmen
