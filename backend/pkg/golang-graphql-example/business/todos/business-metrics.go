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
			ID: "todo-number",
			Collector: prometheus.NewGauge(
				prometheus.GaugeOpts{
					Name: "todo_number",
					Help: "How many todo are present",
				},
			),
			//nolint:mnd
			Interval: 20 * time.Second,
			//nolint:mnd
			Timeout:      5 * time.Second,
			InitialFetch: func(_ context.Context, _ prometheus.Collector) error { return nil },
			Update: func(ctx context.Context, collector prometheus.Collector) error {
				c, err := s.dao.CountTodo(ctx, nil)
				if err != nil {
					return err
				}

				//nolint: forcetypeassert
				collector.(prometheus.Gauge).Set(float64(c))

				return nil
			},
		},
	}
}
