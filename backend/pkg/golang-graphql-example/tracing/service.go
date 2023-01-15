package tracing

import (
	"context"
	"io"
	"net/http"
	"time"

	"emperror.dev/errors"
	"github.com/99designs/gqlgen-contrib/gqlopentracing"
	gqlgraphql "github.com/99designs/gqlgen/graphql"
	opentracing "github.com/opentracing/opentracing-go"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/config"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/uber/jaeger-client-go"
	jaegercfg "github.com/uber/jaeger-client-go/config"
	jaegerprom "github.com/uber/jaeger-lib/metrics/prometheus"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

type service struct {
	closer         io.Closer
	tracer         opentracing.Tracer
	cfgManager     config.Manager
	logger         log.Logger
	metricsFactory *jaegerprom.Factory
}

func (s *service) DatabaseMiddleware() gorm.Plugin {
	return gormopentracing.New(
		gormopentracing.WithCreateOpName("sql:create"),
		gormopentracing.WithUpdateOpName("sql:update"),
		gormopentracing.WithQueryOpName("sql:query"),
		gormopentracing.WithDeleteOpName("sql:delete"),
		gormopentracing.WithRowOpName("sql:row"),
		gormopentracing.WithRawOpName("sql:raw"),
	)
}

func (s *service) GraphqlMiddleware() gqlgraphql.HandlerExtension {
	return gqlopentracing.Tracer{}
}

func (s *service) GetTracer() opentracing.Tracer {
	return s.tracer
}

func (s *service) StartTrace(operationName string) Trace {
	// Start a new span from tracer
	sp := s.tracer.StartSpan(operationName)

	// Return trace object with span
	return &trace{span: sp}
}

func (s *service) StartChildTraceOrTraceFromContext(ctx context.Context, operationName string) Trace {
	// Get trace from context
	tr := GetTraceFromContext(ctx)
	// Check if it exists
	if tr != nil {
		return tr.GetChildTrace(operationName)
	}

	// Create new trace
	return s.StartTrace(operationName)
}

func (s *service) ExtractFromTextMapAndStartSpan(txtMap map[string]string, operationName string) (Trace, error) {
	// Get carrier
	carrier := opentracing.TextMapCarrier(txtMap)

	// Extract
	sctx, err := s.tracer.Extract(opentracing.TextMap, carrier)
	// Check error
	if err != nil && !errors.Is(err, opentracing.ErrSpanContextNotFound) {
		return nil, errors.WithStack(err)
	}

	// Start span
	sp := s.tracer.StartSpan(operationName, opentracing.ChildOf(sctx))

	return &trace{span: sp}, nil
}

func (s *service) ExtractFromHTTPHeaderAndStartSpan(headers http.Header, operationName string) (Trace, error) {
	// Get carrier
	carrier := opentracing.HTTPHeadersCarrier(headers)

	// Extract
	sctx, err := s.tracer.Extract(opentracing.HTTPHeaders, carrier)
	// Check error
	if err != nil && !errors.Is(err, opentracing.ErrSpanContextNotFound) {
		return nil, errors.WithStack(err)
	}

	// Start span
	sp := s.tracer.StartSpan(operationName, opentracing.ChildOf(sctx))

	return &trace{span: sp}, nil
}

func (s *service) InitializeAndReload() error {
	// Save closer
	cl := s.closer

	// Setup
	err := s.setup()
	// Check error
	if err != nil {
		return err
	}

	// Close old one
	err = cl.Close()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	return nil
}

func (s *service) setup() error {
	cfg := s.cfgManager.GetConfig()
	// Initialize configuration
	jcfg := jaegercfg.Configuration{
		ServiceName: "golang-graphql-example",
		Sampler: &jaegercfg.SamplerConfig{
			Type:  jaeger.SamplerTypeConst,
			Param: 1,
		},
	}

	// Check if configuration can be set
	if !cfg.Tracing.Enabled {
		jcfg.Disabled = true
	} else {
		// Add reporter configuration
		jcfg.Reporter = &jaegercfg.ReporterConfig{
			LogSpans:  cfg.Tracing.LogSpan,
			QueueSize: cfg.Tracing.QueueSize,
		}

		// Check if flush interval is customized
		if cfg.Tracing.FlushInterval != "" {
			// Try to parse duration for flush interval
			dur, err := time.ParseDuration(cfg.Tracing.FlushInterval)
			if err != nil {
				return errors.WithStack(err)
			}

			jcfg.Reporter.BufferFlushInterval = dur
		}

		// Check if UDP is customized
		if cfg.Tracing.UDPHost != "" {
			jcfg.Reporter.LocalAgentHostPort = cfg.Tracing.UDPHost
		}
	}

	// Initialize tracer with a logger and a metrics factory
	tracer, closer, err := jcfg.NewTracer(
		jaegercfg.Logger(s.logger.GetTracingLogger()),
		jaegercfg.Metrics(s.metricsFactory),
	)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Set the singleton opentracing.Tracer with the Jaeger tracer.
	opentracing.SetGlobalTracer(tracer)

	s.closer = closer
	s.tracer = tracer

	return nil
}

func newService(cfgManager config.Manager, logger log.Logger) (*service, error) {
	// Create prometheus metrics factory
	factory := jaegerprom.New()

	svc := &service{
		cfgManager:     cfgManager,
		logger:         logger,
		metricsFactory: factory,
	}

	// Run setup
	err := svc.setup()
	if err != nil {
		return nil, err
	}

	return svc, nil
}
