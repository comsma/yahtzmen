## dev/migrations/new name=$1: create a new database migration
db/migrations/new:
	@echo "Creating migration files for ${name}..."
	migrate create -seq -ext=.sql -dir=./migrations ${name}
db/migrations/up:
	@echo "Running up migrations..."
	migrate -path ./migrations -database "${DB_DSN}" up
db/migrations/down:
	@echo "Running up migrations..."
	migrate -path ./migrations -database "${DB_DSN}" down

dev/dependencies/up:
	podman kube play deployment/dev/database.yml

dev/tls:
	go run /opt/homebrew/opt/go/libexec/src/crypto/tls/generate_cert.go --rsa-bits=2048 --host=localhost

gen/api:
	@echo "Generating API..."
	go generate ./...
	openapi-typescript ./openapi.yaml -o ./web/src/api/schema.d.ts


