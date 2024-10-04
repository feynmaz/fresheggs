# Format code with ngofumpt (https://github.com/mvdan/gofumpt).
format:
	gofumpt -w ./internal

# Run golangci-lint (https://github.com/golangci/golangci-lint).
lint:
	golangci-lint run ./internal/...
