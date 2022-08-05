package tracing

import (
	"context"
	"net/http"

	gqlgraphql "github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"gorm.io/gorm"
)

// Service Tracing service.
//
//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing Service
type Service interface {
	// Reload service.
	Reload() error
	// Get opentracing tracer.
	GetTracer() opentracing.Tracer
	// Http Gin HttpMiddleware to add trace per request.
	HTTPMiddleware(getRequestID func(ctx context.Context) string) gin.HandlerFunc
	// Graphql Middleware.
	GraphqlMiddleware() gqlgraphql.HandlerExtension
	// Get database middleware.
	DatabaseMiddleware() gorm.Plugin
	// StartSpan will return a new trace created from scratch.
	StartTrace(operationName string) Trace
	// StartChildTraceOrTraceFromContext will return a child trace if a trace is found inside
	// the context or a new trace with the operation name.
	StartChildTraceOrTraceFromContext(ctx context.Context, operationName string) Trace
	// ExtractFromTextMapAndStartSpan will extract trace from textmap and start a new one from this one.
	ExtractFromTextMapAndStartSpan(txtMap map[string]string, operationName string) (Trace, error)
	// ExtractFromTextMapAndStartSpan will extract trace from http headers and start a new one from this one.
	ExtractFromHTTPHeaderAndStartSpan(headers http.Header, operationName string) (Trace, error)
}

// Trace structure.
//
//go:generate mockgen -destination=./mocks/mock_Trace.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing Trace
type Trace interface {
	// Add tag to trace.
	SetTag(key string, value interface{})
	// Add tags to trace.
	SetTags(tags map[string]interface{})
	// MarkAsError will mark trace as in error.
	MarkAsError()
	// Get a child trace.
	GetChildTrace(operationName string) Trace
	// End the trace.
	Finish()
	// Get the trace ID.
	GetTraceID() string
	// InjectInHTTPHeader will inject span in http header for forwarding.
	InjectInHTTPHeader(header http.Header) error
	// InjectInTextMap will inject span in text map for forwarding.
	InjectInTextMap(textMap map[string]string) error
}

func New(cfgManager config.Manager, logger log.Logger) (Service, error) {
	return newService(cfgManager, logger)
}
