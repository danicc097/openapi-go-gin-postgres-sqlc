
# openapi-go-gin-postgres-sqlc

[![Go Report Card](https://goreportcard.com/badge/github.com/danicc097/openapi-go-gin-postgres-sqlc)](https://goreportcard.com/report/github.com/danicc097/openapi-go-gin-postgres-sqlc)
[![GoDoc](https://pkg.go.dev/badge/github.com/danicc097/openapi-go-gin-postgres-sqlc)](https://pkg.go.dev/github.com/danicc097/openapi-go-gin-postgres-sqlc)

API-first and Database-first approach with OpenAPI v3 and sqlc codegen.
Featuring an overhaul of the [Go Gin
server](https://github.com/OpenAPITools/openapi-generator/blob/master/docs/generators/go-gin-server.md)
generator templates and a sensible post-generation tool that allows you to use cleanly
structured, easily extendable code by smartly merging nodes
from your modified and generated files' abstract syntax trees.


## OpenAPI schema magic fields

- **Struct tags** with `x-go-custom-tag` in schema fields get appended as is. Example (gin-specific):
```YAML
x-go-custom-tag: binding:"required[,customValidator]" [key:val ]
# Special case for ``format: date-time`` fields
# form data only:
x-go-custom-tag: time_format:"2006-01-02"
# the rest require custom unmarshalling if time is not RFC3339:
# see https://github.com/gin-gonic/gin/issues/1193
# there are some quirks to take into account as well:
# see https://segmentfault.com/a/1190000022264001
x-go-custom-tag: binding:"required"
```

Any custom field with an `x-*` name pattern in the OpenAPI spec will be available in
`vendorExtensions`.

