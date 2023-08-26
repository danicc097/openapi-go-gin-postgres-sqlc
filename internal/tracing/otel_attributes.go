package tracing

import (
	"encoding/json"
	"runtime"
	"strings"

	"go.opentelemetry.io/otel/attribute"
)

// Filterable with <attribute>="...".
const (
	UserIDAttributeKey = attribute.Key("user-id") // for outermost layer only
	ParamsAttributeKey = attribute.Key("params")
)

func GetOTelSpanName(parentIndex int) string {
	pc, _, _, _ := runtime.Caller(parentIndex)
	funcPtr := runtime.FuncForPC(pc)
	if funcPtr == nil {
		return "UnknownFunction"
	}

	return strings.TrimPrefix(funcPtr.Name(), "github.com/danicc097/openapi-go-gin-postgres-sqlc/")
}

func ParamsAttribute(params any) attribute.KeyValue {
	p := ""
	s, err := json.Marshal(params)
	if err == nil {
		p = string(s)
	}

	return ParamsAttributeKey.String(p)
}
