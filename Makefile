.PHONY: generate-api
generate-api:
	oapi-codegen -package="v1" -generate=types -o v1/types.go schemas/v1/schema.yaml
	go mod tidy
