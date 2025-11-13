package tracing

import (
	"context"
	"fmt"
	"net/http"
	"reflect"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"

	oteltrace "go.opentelemetry.io/otel/trace"
)

func GetTraceFromContext(ctx context.Context) Trace {
	sp := oteltrace.SpanFromContext(ctx)
	if sp == nil {
		return nil
	}

	return &trace{
		span: sp,
	}
}

func GetChildTraceFromContext(ctx context.Context, operationName string) (context.Context, Trace) {
	// Get trace from context
	pTrace := GetTraceFromContext(ctx)
	// Check if parent trace is nil
	if pTrace == nil {
		return ctx, nil
	}

	return pTrace.GetChildTrace(ctx, operationName)
}

func GetTraceIDFromContext(ctx context.Context) string {
	tr := GetTraceFromContext(ctx)
	if tr != nil {
		return tr.GetTraceID()
	}

	return ""
}

func SetTraceToContext(ctx context.Context, t Trace) context.Context {
	return oteltrace.ContextWithSpan(ctx, t.(*trace).span) //nolint: forcetypeassert // Ignored
}

func InjectInHTTPHeader(ctx context.Context, header http.Header) {
	// Inject
	otel.GetTextMapPropagator().Inject(ctx, propagation.HeaderCarrier(header))
}

func InjectInTextMap(ctx context.Context, textMap map[string]string) {
	// Inject
	otel.GetTextMapPropagator().Inject(ctx, propagation.MapCarrier(textMap))
}

func ExtractFromTextMapAndStartSpan(
	ctx context.Context,
	txtMap map[string]string,
	operationName string,
) (context.Context, Trace) {
	// Get carrier
	carrier := propagation.MapCarrier(txtMap)

	// Extract and set in context
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	// Create trace
	tr := &trace{span: oteltrace.SpanFromContext(ctx)}

	return tr.GetChildTrace(ctx, operationName)
}

func ExtractFromHTTPHeaderAndStartSpan(
	ctx context.Context,
	headers http.Header,
	operationName string,
) (context.Context, Trace) {
	// Get carrier
	carrier := propagation.HeaderCarrier(headers)

	// Extract and set in context
	ctx = otel.GetTextMapPropagator().Extract(ctx, carrier)

	// Create trace
	tr := &trace{span: oteltrace.SpanFromContext(ctx)}

	return tr.GetChildTrace(ctx, operationName)
}

func manageGenericAttribute(key string, value any) *attribute.KeyValue {
	// Create res
	var res attribute.KeyValue

	// Parse
	aValue := reflect.ValueOf(value)
	switch aValue.Kind() { //nolint:exhaustive // That's why "default" is present...
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		res = attribute.Int64(key, aValue.Int())
	case reflect.Float32, reflect.Float64:
		res = attribute.Float64(key, aValue.Float())
	case reflect.Bool:
		res = attribute.Bool(key, aValue.Bool())
	case reflect.String:
		res = attribute.String(key, aValue.String())
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		res = attribute.Int64(key, int64(aValue.Uint())) //nolint:gosec // Overflow conversion shouldn't arrive here
	default:
		res = attribute.String(key, fmt.Sprintf("%v", value))
	}

	return &res
}
