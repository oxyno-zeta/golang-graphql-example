package tracing

import (
	"context"

	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

// Service Tracing service
type Service interface {
	Reload() error
	GetTracer() opentracing.Tracer
	Middleware(getRequestID func(ctx context.Context) string) gin.HandlerFunc
}

// Trace structure
type Trace interface {
	SetTag(key string, value interface{})
	GetChildTrace(operationName string) Trace
	Finish()
	GetTraceID() string
}

func New(cfgManager config.Manager, logger log.Logger) (Service, error) {
	return newService(cfgManager, logger)
}
