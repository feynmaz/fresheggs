.PHONY: openapi format

# Format code with ngofumpt (https://github.com/mvdan/gofumpt).
format:
	gofumpt -w ./internal

# Run golangci-lint (https://github.com/golangci/golangci-lint).
lint:
	golangci-lint run ./internal/...

# Generate OpenAPI v2 spec (https://github.com/swaggo/swag).
openapi:
	swag init -g api.go --parseDependency --parseInternal --dir ./internal/api --output ./openapi

# Run unit tests.
test:
	go test -v -cover ./internal/...
	
# Generate OpenAPI spec and build app.
build: format openapi
	CGO_ENABLED=0 cd ./cmd && go build -o server

# Build and run app with config file provided.
# Example:
# `make run args="-c /home/user/template_config.yaml"`
run: build
	cd ./cmd && ./server $(args)

docker-image:

compose: