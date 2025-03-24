package tracing

import (
	"context"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
)

func (*service) HTTPMiddlewareList(
	getRequestID func(ctx context.Context) string,
) []gin.HandlerFunc {
	return []gin.HandlerFunc{
		otelgin.Middleware(serviceName),
		func(c *gin.Context) {
			// Get trace
			t := GetTraceFromContext(c.Request.Context())
			// Add attributes
			t.SetTags(map[string]interface{}{
				"http.host":       c.Request.Host,
				"http.request_id": getRequestID(c.Request.Context()),
			})
		},
	}
}
