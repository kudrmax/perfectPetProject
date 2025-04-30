OPENAPI_FILE := openapi/openapi.yaml
OPENAPI_IN_ONE_FILE := openapi/openapi.gen.yaml
OAPI_CONFIG := configs/oapi-codegen.yaml
OAPI_PKG := github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen


.PHONY: generate

bundle:
	redocly bundle $(OPENAPI_FILE) --output $(OPENAPI_IN_ONE_FILE)

codegen: bundle
	GOFLAGS=-mod=mod go run $(OAPI_PKG) --config $(OAPI_CONFIG) $(OPENAPI_IN_ONE_FILE)
