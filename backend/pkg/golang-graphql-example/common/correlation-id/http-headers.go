package correlationid

import (
	"context"
	"net/http"
)

const (
	requestIDHeader     = "X-Request-Id"
	correlationIDHeader = "X-Correlation-Id"
)

func SetInHeaders(ctx context.Context, h http.Header) {
	// Get correlation id from context
	correlationID := GetFromContext(ctx)

	// Check if it exists
	if correlationID != "" {
		// Add request id (to be backward compatible)
		h.Set(requestIDHeader, correlationID)
		// Add correlation id
		h.Set(correlationIDHeader, correlationID)
	}
}
