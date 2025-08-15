package tracing

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

const TraceIDHeaderName = "X-Trace-ID"

func (*service) HTTPMiddlewareList(
	getRequestID func(ctx context.Context) string,
) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		otelgin.Middleware(serviceName),
		func(c *gin.Context) {
			// Get trace
			t := GetTraceFromContext(c.Request.Context())
			// Add attributes
			t.SetTags(map[string]any{
				"http.host":       c.Request.Host,
				"http.request_id": getRequestID(c.Request.Context()),
			})

			// Get trace id
			traceID := t.GetTraceID()
			// Check if it exists
			if traceID != "" {
				// Add trace id into result headers
				c.Header(TraceIDHeaderName, t.GetTraceID())
			}
		},
	}
}
