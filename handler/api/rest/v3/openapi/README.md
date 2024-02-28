[oapi-codegen](https://github.com/deepmap/oapi-codegen) is used.

```sh
go install github.com/deepmap/oapi-codegen/v2/cmd/oapi-codegen@latest

cd handler/api/v3/openapi

oapi-codegen -config config.yaml spec.yaml
```