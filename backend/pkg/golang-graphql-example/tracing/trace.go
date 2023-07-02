package tracing

import (
	"context"
	"net/http"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	oteltrace "go.opentelemetry.io/otel/trace"
)

type trace struct {
	span oteltrace.Span
}

func (t *trace) MarkAsError() {
	t.SetTag("error", true)
}

func (t *trace) SetTag(key string, value interface{}) {
	t.span.SetAttributes(*manageGenericAttribute(key, value))
}

func (t *trace) SetTags(tags map[string]interface{}) {
	for k, v := range tags {
		t.SetTag(k, v)
	}
}

func (*trace) GetChildTrace(ctx context.Context, operationName string) (context.Context, Trace) {
	ctx, sp := getTracerFromTraceProvider(otel.GetTracerProvider()).Start(ctx, operationName)

	return ctx, &trace{span: sp}
}

func (t *trace) InjectInHTTPHeader(header http.Header) {
	// Create fake context
	ctx := oteltrace.ContextWithSpan(context.TODO(), t.span)

	// Inject
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(header))
}

func (t *trace) InjectInTextMap(textMap map[string]string) {
	// Create fake context
	ctx := oteltrace.ContextWithSpan(context.TODO(), t.span)

	// Inject
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(textMap))
}

func (t *trace) Finish() {
	t.span.End()
}

func (t *trace) GetTraceID() string {
	if t.span.SpanContext().HasTraceID() {
		return t.span.SpanContext().TraceID().String()
	}

	return ""
}
