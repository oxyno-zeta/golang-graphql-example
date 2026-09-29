package metrics

import (
	"context"
	"fmt"
	"time"

	"emperror.dev/errors"
	"github.com/prometheus/client_golang/prometheus"
	"golang.org/x/sync/errgroup"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

const (
	businessMetricsGeneralInitialFetchTraceName    = "metrics.business-metrics.general-initial-fetch"
	businessMetricsDefinitionInitialFetchTraceName = "metrics.business-metrics.definition-initial-fetch"
	businessMetricsDefinitionUpdateTraceName       = "metrics.business-metrics.definition-update"
	businessMetricsTraceIDTagKey                   = "business-metrics.id"
	businessMetricsLogIDTagKey                     = "business-metric-id"
	defaultTimeout                                 = 20 * time.Second
	defaultInternal                                = 30 * time.Second
	metricSuccessStatus                            = "success"
	metricErrorStatus                              = "error"
)

type BusinessMetricDefinition struct {
	Collector    prometheus.Collector
	InitialFetch func(ctx context.Context, collector prometheus.Collector) error
	Update       func(ctx context.Context, collector prometheus.Collector) error
	ID           string
	Interval     time.Duration
	Timeout      time.Duration
}

func (impl *prometheusMetrics) AddBusinessMetricDefinition(input *BusinessMetricDefinition) {
	// Add default timeout if not present
	if input.Timeout.Nanoseconds() == 0 {
		input.Timeout = defaultTimeout
	}

	// Add default interval if not present
	if input.Interval.Nanoseconds() == 0 {
		input.Interval = defaultInternal
	}

	impl.businessMetricsDefinitions = append(impl.businessMetricsDefinitions, input)
}

func (impl *prometheusMetrics) RegisterBusinessMetricDefinitions() error {
	// Loop over definitions
	for _, def := range impl.businessMetricsDefinitions {
		// Register
		err := prometheus.Register(def.Collector)
		// Check error
		if err != nil {
			// Wrap to inform which error it is
			return errors.Wrap(err, fmt.Sprintf("collector %s error", def.ID))
		}
	}

	// Default
	return nil
}

func (impl *prometheusMetrics) InitialFetchBusinessMetricDefinitions(ctx context.Context) (err error) {
	// Start trace
	ctx, trace := impl.tracingSvc.StartTrace(ctx, businessMetricsGeneralInitialFetchTraceName)

	// Defer close
	defer func() {
		if err != nil {
			trace.MarkAsError()
			trace.AddError(err)
		}

		// Finish
		trace.Finish()
	}()

	// Create error group in order to run all initial fetch in parallel
	group, groupCtx := errgroup.WithContext(ctx)
	// Loop over definitions
	for _, def := range impl.businessMetricsDefinitions {
		// Start routine
		group.Go(func() error {
			// Call init
			err := impl.internalInitialFetchBusinessMetricDefinitions(groupCtx, def)
			// Check error
			if err != nil {
				return errors.Wrap(err, fmt.Sprintf("collector %s error", def.ID))
			}

			// Default
			return nil
		})
	}

	return group.Wait()
}

func (impl *prometheusMetrics) StartUpdaterBusinessMetricDefinitions(mainLogger log.Logger) {
	// Loop over definitions
	for _, def := range impl.businessMetricsDefinitions {
		// Create new context with logger
		ctx := log.SetLoggerToContext(context.Background(), mainLogger)

		// Start updater in goroutine
		go impl.internalStartUpdaterBusinessMetricDefinitions(ctx, def)
	}
}

func (impl *prometheusMetrics) internalStartUpdaterBusinessMetricDefinitions(ctx context.Context, def *BusinessMetricDefinition) {
	// Create function to perform the update
	f := func() {
		// Create context
		nCtx, cancel := impl.createContextForBusinessMetricDefinition(ctx, def)

		// Defer
		defer cancel()

		// Start trace
		nCtx, trace := impl.tracingSvc.StartTrace(nCtx, businessMetricsDefinitionUpdateTraceName)
		// Add metadata
		trace.SetTag(businessMetricsTraceIDTagKey, def.ID)

		// Initialize metric status
		metricStatus := metricSuccessStatus

		// Perform update
		err := def.Update(nCtx, def.Collector)
		// Check error
		if err != nil {
			logger := log.GetLoggerFromContext(nCtx)

			// Mark trace as error
			trace.MarkAsError()
			trace.AddError(err)

			// Change status
			metricStatus = metricErrorStatus

			// log error
			logger.Error(err)
		}

		// Increase counter
		impl.businessMetricUpdateManaged.WithLabelValues(def.ID, metricStatus).Inc()

		// Finish
		trace.Finish()
	}

	// Infinite loop
	for {
		// Start
		f()

		// Wait
		time.Sleep(def.Interval)
	}
}

func (impl *prometheusMetrics) internalInitialFetchBusinessMetricDefinitions(ctx context.Context, def *BusinessMetricDefinition) (err error) {
	// Create context
	nCtx, cancel := impl.createContextForBusinessMetricDefinition(ctx, def)

	// Defer
	defer cancel()

	// Start trace
	ctx, trace := impl.tracingSvc.StartTrace(ctx, businessMetricsDefinitionInitialFetchTraceName)
	// Add metadata
	trace.SetTag(businessMetricsTraceIDTagKey, def.ID)

	// Defer close
	defer func() {
		if err != nil {
			trace.MarkAsError()
			trace.AddError(err)
		}

		// Finish
		trace.Finish()
	}()

	// Call init
	return def.InitialFetch(nCtx, def.Collector)
}

func (impl *prometheusMetrics) createContextForBusinessMetricDefinition(
	ctx context.Context,
	def *BusinessMetricDefinition,
) (context.Context, context.CancelFunc) {
	// Get logger from context
	logger := log.GetLoggerFromContext(ctx)

	// Add information to logger
	logger = logger.WithField(businessMetricsLogIDTagKey, def.ID)

	// Add back logger to context
	resCtx := log.SetLoggerToContext(ctx, logger)

	// Add timeout
	return context.WithTimeout(resCtx, def.Timeout)
}
