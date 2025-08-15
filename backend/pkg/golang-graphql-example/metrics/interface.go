package metrics

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	gqlgraphql "github.com/99designs/gqlgen/graphql"
)

// Avoid adding a big number because getting metrics get a lock on gorm.
const defaultPrometheusGormRefreshMetricsSecond = 15

// Service Service metrics interface.
//
//go:generate mockgen -destination=./mocks/mock_Service.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics Service
type Service interface {
	// Instrument web server.
	Instrument(serverName string, routerPath bool) gin.HandlerFunc
	// Get prometheus handler for http expose.
	PrometheusHTTPHandler() http.Handler
	// Get database middleware.
	DatabaseMiddleware(connectionName string) gorm.Plugin
	// Get graphql middleware.
	GraphqlMiddleware() gqlgraphql.HandlerExtension
	// IncreaseSuccessfullyAMQPConsumedMessage will increase counter of successfully AMQP consumed message.
	IncreaseSuccessfullyAMQPConsumedMessage(queue, consumerTag, routingKey string)
	// IncreaseFailedAMQPConsumedMessage will increase counter of failed AMQP consumed message.
	IncreaseFailedAMQPConsumedMessage(queue, consumerTag, routingKey string)
	// IncreaseSuccessfullyAMQPPublishedMessage will increase counter of successfully AMQP published message.
	IncreaseSuccessfullyAMQPPublishedMessage(exchange, routingKey string)
	// IncreaseFailedAMQPPublishedMessage will increase counter of failed AMQP published message.
	IncreaseFailedAMQPPublishedMessage(exchange, routingKey string)
	// UpFailedConfigReload will raise the failed configuration reload gauge.
	UpFailedConfigReload()
	// DownFailedConfigReload will down the failed configuration reload gauge.
	DownFailedConfigReload()
}

// NewService will generate a new Service.
func NewService() Service {
	ctx := &prometheusMetrics{
		gormPrometheus: map[string]gorm.Plugin{},
	}
	// Register
	ctx.register()

	return ctx
}
