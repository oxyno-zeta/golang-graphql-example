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

type UncorrelatedTraceOutput struct {
	ChildTrace          Trace
	ChildContext        context.Context //nolint:containedctx // Won't do a 4 output function
	UncorrelatedTrace   Trace
	UncorrelatedContext context.Context //nolint:containedctx // Won't do a 4 output function
}

type TraceEventOption = oteltrace.EventOption

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
	// StartUncorrelatedChildTrace will create a child trace on the parent trace and create a new trace from scratch and reference this one in the
	// new child trace created.
	// This is allowing to create non correlated trace and spread trace load across multiple traces.
	StartUncorrelatedChildTrace(
		ctx context.Context,
		parentTrace Trace,
		childTraceName, uncorrelatedTraceName string,
	) *UncorrelatedTraceOutput
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
	// AddError will add error message in event trace.
	// ! Warning: This won't mark trace as error
	AddError(err error, opts ...TraceEventOption)
	// AddAndMarkError will add error message in event trace and mark trace as in error.
	AddAndMarkError(err error, opts ...TraceEventOption)
	// AddEvent will add an event in trace.
	AddEvent(eventName string, opts ...TraceEventOption)
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
