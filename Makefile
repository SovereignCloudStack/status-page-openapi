GENERATOR=npx openapi-generator-cli
GO_GENERATOR=oapi-codegen
SPEC=openapi.yaml

GO_DIR=pkg/api

.DEFAULT_GOAL := all

${GO_DIR}:
	@mkdir -p $@

${GO_DIR}/server: ${GO_DIR}
	@mkdir -p $@

.PHONY: all validate generate-all generate-angular-client generate-go-server

all: validate generate-all

validate:
	${GENERATOR} validate -i ${SPEC}

generate-all: generate-angular-client generate-go-server

generate-angular-client:
	${GENERATOR} generate --generator-key status-page-api-angular-client

generate-go-server: ${GO_DIR}/server
	${GO_GENERATOR} --package server --generate types,server,spec openapi.yaml > pkg/api/server/server.gen.go
	go mod tidy

generate-docs:
	${GENERATOR} generate --generator-key status-page-api-doc
