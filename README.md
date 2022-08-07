# openapi-go-gin-postgres-sqlc

API-first and Database-first approach with OpenAPI v3 and sqlc codegen.


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
