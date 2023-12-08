// Package internal contains internal code for xo.
package symbols

import (
	"reflect"
)

// Symbols are extracted (generated) symbols from the types package.
//
//go:generate yaegi extract os/exec
//go:generate yaegi extract github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/repos/postgresql/gen/db
//go:generate yaegi extract github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/rest
//go:generate yaegi extract github.com/danicc097/openapi-go-gin-postgres-sqlc/internal/utils/structs
//go:generate yaegi extract github.com/fatih/structtag
//go:generate yaegi extract github.com/google/uuid
//go:generate yaegi extract github.com/iancoleman/strcase
//go:generate yaegi extract github.com/swaggest/jsonschema-go
//go:generate yaegi extract github.com/swaggest/openapi-go/openapi3
var Symbols map[string]map[string]reflect.Value = make(map[string]map[string]reflect.Value)
