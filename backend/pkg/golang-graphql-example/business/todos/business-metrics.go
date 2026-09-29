package todos

import (
	"context"
	"time"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
)

func (s *service) GetBusinessMetricDefinitions() []*metrics.BusinessMetricDefinition {
	return []*metrics.BusinessMetricDefinition{
		{
			ID: "todo-count",
			Collector: prometheus.NewGauge(
				prometheus.GaugeOpts{
					Name: "todo_count",
					Help: "How many todo count",
				},
			),
			Interval:     20 * time.Second,
			Timeout:      5 * time.Second,
			InitialFetch: func(ctx context.Context, collector prometheus.Collector) error { return nil },
			Update: func(ctx context.Context, collector prometheus.Collector) error {
				c, err := s.dao.CountTodo(ctx, nil)
				if err != nil {
					return err
				}

				// nolint: forcetypeassert
				collector.(prometheus.Gauge).Set(float64(c))

				return nil
			},
		},
	}
}
