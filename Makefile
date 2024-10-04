API_VERSION = v1

.PHONY: generate-api
generate-api:
	oapi-codegen -package="$(API_VERSION)" -generate types -o internal/api/$(API_VERSION)/openapi_types.gen.go openapi/$(API_VERSION)/docs.yaml
	oapi-codegen -package="$(API_VERSION)" -generate chi-server -o internal/api/$(API_VERSION)/openapi_api.gen.go openapi/$(API_VERSION)/docs.yaml
	go mod tidy
