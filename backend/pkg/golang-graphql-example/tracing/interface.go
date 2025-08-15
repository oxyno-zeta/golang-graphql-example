package tracing

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	gqlgraphql "github.com/99designs/gqlgen/graphql"
	oteltrace "go.opentelemetry.io/otel/trace"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

// Service Tracing service.
//
//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing Service
type Service interface {
	// InitializeAndReload service.
	InitializeAndReload() error
	// Http Gin HttpMiddleware list to add trace per request.
	HTTPMiddlewareList(getRequestID func(ctx context.Context) string) []gin.HandlerFunc
	// Graphql Middleware.
	GraphqlMiddleware() gqlgraphql.HandlerExtension
	// Get database middleware.
	DatabaseMiddleware() gorm.Plugin
	// StartSpan will return a new trace created.
	StartTrace(
		ctx context.Context,
		operationName string,
		opts ...oteltrace.SpanStartOption,
	) (context.Context, Trace)
	// Close is used on application shutdown.
	Close() error
}

// Trace structure.
//
//go:generate mockgen -destination=./mocks/mock_Trace.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/tracing Trace
type Trace interface {
	// Add tag to trace.
	SetTag(key string, value any)
	// Add tags to trace.
	SetTags(tags map[string]any)
	// MarkAsError will mark trace as in error.
	MarkAsError()
	// Get a child trace.
	GetChildTrace(ctx context.Context, operationName string) (context.Context, Trace)
	// End the trace.
	Finish()
	// Get the trace ID.
	GetTraceID() string
	// InjectInHTTPHeader will inject span in http header for forwarding.
	// @deprecated: Use global method
	InjectInHTTPHeader(header http.Header)
	// InjectInTextMap will inject span in text map for forwarding.
	// @deprecated: Use global method
	InjectInTextMap(textMap map[string]string)
}

func New(cfgManager config.Manager, logger log.Logger) Service {
	return &service{
		cfgManager: cfgManager,
		logger:     logger,
	}
}
