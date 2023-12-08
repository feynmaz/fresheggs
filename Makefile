API_VERSION = v1

.PHONY: generate-api
generate-api:
	oapi-codegen -package="$(API_VERSION)" -generate=types,chi-server -o internal/controller/http/$(API_VERSION)/codegen.go schemas/$(API_VERSION)/schema.yaml
	go mod tidy
