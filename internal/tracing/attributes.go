package tracing

import (
	"go.opentelemetry.io/otel/attribute"
)

// Filterable with user-id="..."
// In frontend we would have something unique (and not personally identifiable)
// to each navigation to correlate user interaction, fetch and document load traces
const UserIDAttribute = attribute.Key("user-id")
