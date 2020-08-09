// +build integration

package server

import (
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics"
)

// Generate metrics instance
var metricsCtx = metrics.NewMetricsClient()
