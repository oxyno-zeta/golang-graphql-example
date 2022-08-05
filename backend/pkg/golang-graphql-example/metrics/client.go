package metrics

import (
	"net/http"

	gqlgraphql "github.com/99designs/gqlgen/graphql"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// Avoid adding a big number because getting metrics get a lock on gorm.
const defaultPrometheusGormRefreshMetricsSecond = 15

// Client Client metrics interface.
//
//go:generate mockgen -destination=./mocks/mock_Client.go -package=mocks github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/metrics Client
type Client interface {
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
}

// NewMetricsClient will generate a new Client.
func NewMetricsClient() Client {
	ctx := &prometheusMetrics{
		gormPrometheus: map[string]gorm.Plugin{},
	}
	// Register
	ctx.register()

	return ctx
}
