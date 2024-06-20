package metrics

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"

	gqlprometheus "github.com/99designs/gqlgen-contrib/prometheus"
	gqlgraphql "github.com/99designs/gqlgen/graphql"
	gormprometheus "gorm.io/plugin/prometheus"
)

type prometheusMetrics struct {
	reqCnt                *prometheus.CounterVec
	resSz                 *prometheus.SummaryVec
	reqDur                *prometheus.SummaryVec
	reqSz                 *prometheus.SummaryVec
	up                    prometheus.Gauge
	configReloadFail      prometheus.Gauge
	gormPrometheus        map[string]gorm.Plugin
	amqpConsumedMessages  *prometheus.CounterVec
	amqpPublishedMessages *prometheus.CounterVec
}

func (*prometheusMetrics) GraphqlMiddleware() gqlgraphql.HandlerExtension {
	return gqlprometheus.Tracer{}
}

func (*prometheusMetrics) PrometheusHTTPHandler() http.Handler {
	return promhttp.Handler()
}

func (impl *prometheusMetrics) UpFailedConfigReload() {
	impl.configReloadFail.Set(1)
}

func (impl *prometheusMetrics) DownFailedConfigReload() {
	impl.configReloadFail.Set(0)
}

func (impl *prometheusMetrics) IncreaseSuccessfullyAMQPConsumedMessage(queue, consumerTag, routingKey string) {
	impl.amqpConsumedMessages.WithLabelValues(queue, consumerTag, routingKey, "success").Inc()
}

func (impl *prometheusMetrics) IncreaseFailedAMQPConsumedMessage(queue, consumerTag, routingKey string) {
	impl.amqpConsumedMessages.WithLabelValues(queue, consumerTag, routingKey, "error").Inc()
}

func (impl *prometheusMetrics) IncreaseSuccessfullyAMQPPublishedMessage(exchange, routingKey string) {
	impl.amqpPublishedMessages.WithLabelValues(exchange, routingKey, "success").Inc()
}

func (impl *prometheusMetrics) IncreaseFailedAMQPPublishedMessage(exchange, routingKey string) {
	impl.amqpPublishedMessages.WithLabelValues(exchange, routingKey, "error").Inc()
}

// The gorm prometheus plugin cannot be instantiated twice because there is a loop inside that cannot be modified or stopped.
// This loop get all data from database and the loop cannot be modified in terms of the duration.
// Labels and all other options cannot be modified.
func (impl *prometheusMetrics) DatabaseMiddleware(connectionName string) gorm.Plugin {
	// Check if gorm prometheus doesn't already exist
	if impl.gormPrometheus[connectionName] == nil {
		// Create middleware
		md := gormprometheus.New(gormprometheus.Config{
			RefreshInterval: defaultPrometheusGormRefreshMetricsSecond, // refresh metrics interval (default 15 seconds)
		})
		// Apply labels
		md.Labels = map[string]string{
			"connection_name": connectionName,
		}
		// Save it
		impl.gormPrometheus[connectionName] = md
	}

	return impl.gormPrometheus[connectionName]
}

// Instrument will instrument gin routes.
func (impl *prometheusMetrics) Instrument(serverName string, routerPath bool) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		reqSz := computeApproximateRequestSize(c.Request)

		c.Next()

		status := strconv.Itoa(c.Writer.Status())
		elapsed := float64(time.Since(start)) / float64(time.Second)
		resSz := float64(c.Writer.Size())

		path := c.Request.URL.Path
		if routerPath {
			path = c.FullPath()
		}

		impl.reqDur.WithLabelValues(serverName).Observe(elapsed)
		impl.reqCnt.WithLabelValues(serverName, status, c.Request.Method, c.Request.Host, path).Inc()
		impl.reqSz.WithLabelValues(serverName).Observe(float64(reqSz))
		impl.resSz.WithLabelValues(serverName).Observe(resSz)
	}
}

// From https://github.com/DanielHeckrath/gin-prometheus/blob/master/gin_prometheus.go
func computeApproximateRequestSize(r *http.Request) int {
	s := 0
	if r.URL != nil {
		s = len(r.URL.Path)
	}

	s += len(r.Method)
	s += len(r.Proto)

	for name, values := range r.Header {
		s += len(name)
		for _, value := range values {
			s += len(value)
		}
	}

	s += len(r.Host)

	// N.B. r.Form and r.MultipartForm are assumed to be included in r.URL.

	if r.ContentLength != -1 {
		s += int(r.ContentLength)
	}

	return s
}

func (impl *prometheusMetrics) register() {
	impl.reqCnt = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "http_requests_total",
			Help: "How many HTTP requests processed, partitioned by status code and HTTP method.",
		},
		[]string{"server_name", "status_code", "method", "host", "path"},
	)
	prometheus.MustRegister(impl.reqCnt)

	impl.reqDur = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_duration_seconds",
			Help: "The HTTP request latencies in seconds.",
		},
		[]string{"server_name"},
	)
	prometheus.MustRegister(impl.reqDur)

	impl.reqSz = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_request_size_bytes",
			Help: "The HTTP request sizes in bytes.",
		},
		[]string{"server_name"},
	)
	prometheus.MustRegister(impl.reqSz)

	impl.resSz = prometheus.NewSummaryVec(
		prometheus.SummaryOpts{
			Name: "http_response_size_bytes",
			Help: "The HTTP response sizes in bytes.",
		},
		[]string{"server_name"},
	)
	prometheus.MustRegister(impl.resSz)

	impl.up = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "up",
			Help: "1 = up, 0 = down",
		},
	)
	impl.up.Set(1)
	prometheus.MustRegister(impl.up)

	impl.configReloadFail = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "config_reload_fail",
			Help: "1 = last config reload failed, 0 = last config reload was ok",
		},
	)
	prometheus.MustRegister(impl.configReloadFail)

	impl.amqpConsumedMessages = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "amqp_consumed_messages_total",
			Help: "How many AMQP messages have been consumed by queue, consumer tag, routing key and status",
		},
		[]string{"queue", "consumer_tag", "routing_key", "status"},
	)
	prometheus.MustRegister(impl.amqpConsumedMessages)

	impl.amqpPublishedMessages = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "amqp_published_messages_total",
			Help: "How many AMQP messages have been published by exchange, routing key and status",
		},
		[]string{"exchange", "routing_key", "status"},
	)
	prometheus.MustRegister(impl.amqpPublishedMessages)

	// Register gqlgen graphql prometheus metrics
	gqlprometheus.Register()
}
