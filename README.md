# SCS Status Page OpenAPI Spec

See `openapi.yaml` file for spec.

## Generated Code

See `./pkg/api` for Go module containing generated code.

In order to regenerate code, `oapi-codegen` is used.

```
$ go install github.com/deepmap/oapi-codegen/cmd/oapi-codegen@latest
```

For further infos see [upstream documentation](https://github.com/deepmap/oapi-codegen).

See `Makefile` on how to use oapi-codegen to regenerate code.
