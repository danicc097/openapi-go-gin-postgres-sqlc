package tracing

import (
	"runtime"
	"strings"

	"go.opentelemetry.io/otel/attribute"
)

// Filterable with user-id="..."
// In frontend we would have a combination of user-id and random string
// to each navigation to correlate user interaction, fetch and document load traces
// from each open instance.
const UserIDAttribute = attribute.Key("user-id")

func GetOTelSpanName(parentIndex int) string {
	pc, _, _, _ := runtime.Caller(parentIndex)
	funcPtr := runtime.FuncForPC(pc)
	if funcPtr == nil {
		return "UnknownFunction"
	}

	return strings.TrimPrefix(funcPtr.Name(), "github.com/danicc097/openapi-go-gin-postgres-sqlc/")
}
