API_VERSION = v1

.PHONY: generate-api
generate-api:
	oapi-codegen -package="$(API_VERSION)" -generate types -o internal/ports/http/$(API_VERSION)/openapi_types.gen.go api/openapi/$(API_VERSION)/product.yaml
	oapi-codegen -package="$(API_VERSION)" -generate chi-server -o internal/ports/http/$(API_VERSION)/openapi_api.gen.go api/openapi/$(API_VERSION)/product.yaml
	go mod tidy

