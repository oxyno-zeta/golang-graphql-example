package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Client Client metrics interface
type Client interface {
	Instrument(serverName string) gin.HandlerFunc
	GetPrometheusHTTPHandler() http.Handler
}

// NewMetricsClient will generate a new Client
func NewMetricsClient() Client {
	ctx := &prometheusMetrics{}
	ctx.register()

	return ctx
}
