package tracing

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/opentracing-contrib/go-gin/ginhttp"
	opentracing "github.com/opentracing/opentracing-go"
)

func (s *service) Middleware(getRequestID func(ctx context.Context) string) gin.HandlerFunc {
	// Add more metadata to span
	opt := ginhttp.MWSpanObserver(func(span opentracing.Span, r *http.Request) {
		// Add request host
		span.SetTag("http.request_host", r.Host)
		// Add request id
		span.SetTag("http.request_id", getRequestID(r.Context()))
		// Add request path
		span.SetTag("http.request_path", r.URL.Path)
	})

	return ginhttp.Middleware(s.tracer, opt)
}
