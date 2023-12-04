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

	MetadataAttributeKey = attribute.Key("metadata")
)

// GetOTelSpanName returns a span name based on the calling package and function.
func GetOTelSpanName(parentIndex int) string {
	pc, _, _, _ := runtime.Caller(parentIndex)
	funcPtr := runtime.FuncForPC(pc)
	if funcPtr == nil {
		return "UnknownFunction"
	}

	return strings.TrimPrefix(funcPtr.Name(), "github.com/danicc097/openapi-go-gin-postgres-sqlc/")
}

// MetadataAttribute allows adding metadata in a standard way to a span.
func MetadataAttribute(metadata any) attribute.KeyValue {
	p := ""
	s, err := json.Marshal(metadata)
	if err == nil {
		p = string(s)
	}

	return MetadataAttributeKey.String(p)
}
