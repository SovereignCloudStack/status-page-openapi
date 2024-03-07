GENERATOR=npx openapi-generator-cli
GO_GENERATOR=oapi-codegen
SPEC=openapi.yaml

all: validate generate-all

validate:
	${GENERATOR} validate -i ${SPEC}

generate-all: generate-angular-client generate-go-server

generate-angular-client:
	${GENERATOR} generate --generator-key status-page-api

generate-go-server:
	${GO_GENERATOR} --package api --generate types,server,spec openapi.yaml > pkg/api/api.gen.go

generate-docs:
	${GENERATOR} generate --generator-key status-page-api-doc
