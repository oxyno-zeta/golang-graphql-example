package tracing

import (
	"context"
	"net/url"
	"time"

	"emperror.dev/errors"
	gqlgraphql "github.com/99designs/gqlgen/graphql"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/version"
	"github.com/ravilushqa/otelgqlgen"
	b3propagator "go.opentelemetry.io/contrib/propagators/b3"
	jaegerpropagator "go.opentelemetry.io/contrib/propagators/jaeger"
	otpropagator "go.opentelemetry.io/contrib/propagators/ot"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/sdk/resource"
	tracesdk "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.17.0"
	oteltrace "go.opentelemetry.io/otel/trace"
	"gorm.io/gorm"
	gormtracing "gorm.io/plugin/opentelemetry/tracing"
)

const serviceName = "golang-graphql-example"
const tracerName = serviceName

type service struct {
	tracerProvider *tracesdk.TracerProvider
	tracer         oteltrace.Tracer
	cfgManager     config.Manager
	logger         log.Logger
}

func (*service) DatabaseMiddleware() gorm.Plugin {
	return gormtracing.NewPlugin(gormtracing.WithoutMetrics())
}

func (*service) GraphqlMiddleware() gqlgraphql.HandlerExtension {
	return otelgqlgen.Middleware()
}

func (s *service) GetTracer() oteltrace.Tracer {
	return s.tracer
}

func (s *service) StartTrace(ctx context.Context, operationName string) (context.Context, Trace) {
	// Start a new span from tracer
	ctx, sp := s.tracer.Start(ctx, operationName)

	// Return trace object with span
	return ctx, &trace{span: sp}
}

func (s *service) Close() error {
	// Check if tracer provider exists
	if s.tracerProvider != nil {
		// Shutdown
		return errors.WithStack(s.tracerProvider.Shutdown(context.Background()))
	}

	// Default
	return nil
}

func (s *service) InitializeAndReload() error {
	// Get configuration
	cfg := s.cfgManager.GetConfig()

	// Build option array
	tracerOpts := []tracesdk.TracerProviderOption{}

	// Check if enabled
	if cfg.Tracing != nil && cfg.Tracing.Enabled {
		// Init exporter
		var exp tracesdk.SpanExporter
		// Switch on type
		switch cfg.Tracing.Type {
		case config.TracingOtelHTTPType:
			ur, err := url.Parse(cfg.Tracing.OtelHTTP.ServerURL)
			// Check error
			if err != nil {
				return errors.WithStack(err)
			}

			// Create options array
			opts := []otlptracehttp.Option{
				otlptracehttp.WithEndpoint(ur.Host),
				otlptracehttp.WithURLPath(ur.RawPath),
			}
			// Check if http is asked
			if ur.Scheme == "http" {
				opts = append(opts, otlptracehttp.WithInsecure())
			}
			// Check if timeout is set
			if cfg.Tracing.OtelHTTP.TimeoutString != "" {
				dur, err2 := time.ParseDuration(cfg.Tracing.OtelHTTP.TimeoutString)
				// Check error
				if err2 != nil {
					return errors.WithStack(err2)
				}
				// Save
				opts = append(opts, otlptracehttp.WithTimeout(dur))
			}
			// Check if headers are set
			if cfg.Tracing.OtelHTTP.Headers != nil {
				opts = append(opts, otlptracehttp.WithHeaders(cfg.Tracing.OtelHTTP.Headers))
			}

			// Create client
			client := otlptracehttp.NewClient(opts...)
			// Create exporter
			exporter, err := otlptrace.New(context.TODO(), client)
			// Check error
			if err != nil {
				return errors.WithStack(err)
			}
			// Save
			exp = exporter
		default:
			return errors.New("Tracing type not supported")
		}

		// Create batch params
		batchOpts := []tracesdk.BatchSpanProcessorOption{}
		// Check if max queue size is defined
		if cfg.Tracing.MaxQueueSize != 0 {
			batchOpts = append(batchOpts, tracesdk.WithMaxQueueSize(cfg.Tracing.MaxQueueSize))
		}
		// Check if max batch size is defined
		if cfg.Tracing.MaxBatchSize != 0 {
			batchOpts = append(batchOpts, tracesdk.WithMaxExportBatchSize(cfg.Tracing.MaxBatchSize))
		}

		// Exporter
		tracerOpts = append(tracerOpts, tracesdk.WithBatcher(exp, batchOpts...))
	}

	// Prepare attributes
	attributes := []attribute.KeyValue{
		semconv.ServiceName(serviceName),
		semconv.ServiceVersion(version.GetVersion().Version),
	}
	// Check if fixed tags exists
	if cfg.Tracing != nil && cfg.Tracing.FixedTags != nil {
		for k, v := range cfg.Tracing.FixedTags {
			attributes = append(attributes, *manageGenericAttribute(k, v))
		}
	}

	// Create resource with attributes
	res, err := resource.New(
		context.Background(),
		resource.WithFromEnv(),   // pull attributes from OTEL_RESOURCE_ATTRIBUTES and OTEL_SERVICE_NAME environment variables
		resource.WithProcess(),   // This option configures a set of Detectors that discover process information
		resource.WithOS(),        // This option configures a set of Detectors that discover OS information
		resource.WithContainer(), // This option configures a set of Detectors that discover container information
		resource.WithHost(),      // This option configures a set of Detectors that discover host information
		resource.WithAttributes(attributes...),
	)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	var sampler tracesdk.Sampler
	// Create sampling configuration
	switch cfg.Tracing.SamplerType {
	// Check if there is always on sampler is defined or is not defined
	case "", config.TracingSamplerAlwaysOn:
		sampler = tracesdk.AlwaysSample()
	// Check if there is always off sampler is defined
	case config.TracingSamplerAlwaysOff:
		sampler = tracesdk.NeverSample()
	// Check if there is ratio sampler is defined
	case config.TracingSamplerRatio:
		// Check if there is a ratio configuration
		if cfg.Tracing.SamplerCfg != nil && cfg.Tracing.SamplerCfg.RatioCfg != nil {
			sampler = tracesdk.TraceIDRatioBased(cfg.Tracing.SamplerCfg.RatioCfg.Ratio)
		}
	}

	// Add resources and sampling to options
	tracerOpts = append(tracerOpts,
		// Record information about this application in a Resource.
		tracesdk.WithResource(res),
		// Add sampler
		tracesdk.WithSampler(sampler),
	)

	// Create tracer provider
	tp := tracesdk.NewTracerProvider(
		tracerOpts...,
	)

	// Save tracer provider
	s.tracerProvider = tp

	// Save tracer provider
	otel.SetTracerProvider(tp)
	// Save propagator
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		// Jaeger propagator
		jaegerpropagator.Jaeger{},
		// B3 propagators
		b3propagator.New(b3propagator.WithInjectEncoding(b3propagator.B3SingleHeader)),
		b3propagator.New(b3propagator.WithInjectEncoding(b3propagator.B3MultipleHeader)),
		// OpenTracing propagator
		otpropagator.OT{},
		// Otel propagator
		propagation.Baggage{},
		propagation.TraceContext{},
	))

	// Create tracer
	tr := getTracerFromTraceProvider(tp)

	// Save it
	s.tracer = tr

	return nil
}

func getTracerFromTraceProvider(tp oteltrace.TracerProvider) oteltrace.Tracer {
	return tp.Tracer(tracerName)
}
