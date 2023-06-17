package tracing

import (
	"go.opentelemetry.io/otel/attribute"
)

// Filterable with user-id="..."
// In frontend we would have a combination of user-id and random string
// to each navigation to correlate user interaction, fetch and document load traces
// from each open instance.
const UserIDAttribute = attribute.Key("user-id")
