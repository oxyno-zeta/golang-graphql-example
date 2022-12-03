package middlewares

import (
	"github.com/gin-gonic/gin"
	correlationid "github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/correlation-id"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/utils"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

const requestIDHeader = "X-Request-Id"
const correlationIDContextKey = "correlationID"

func CorrelationID(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request id from request
		correlationID := c.Request.Header.Get(requestIDHeader)

		// Check if request id exists
		if correlationID == "" {
			// Generate uuid
			uuid, err := correlationid.Generate()
			// Check error
			if err != nil {
				// Log error
				logger.Error(err)
				// Send response
				utils.AnswerWithError(c, err)

				return
			}
			// Save it in variable
			correlationID = uuid
		}

		// Store it in context
		SetCorrelationIDInGin(c, correlationID)
		// Update request with new context
		c.Request = c.Request.WithContext(
			correlationid.SetInContext(c.Request.Context(), correlationID),
		)

		// Put it on header
		c.Writer.Header().Set(requestIDHeader, correlationID)

		// Next
		c.Next()
	}
}

func GetCorrelationIDFromGin(c *gin.Context) string {
	correlationIDObj, correlationIDExists := c.Get(correlationIDContextKey)
	if correlationIDExists {
		// return request id
		return correlationIDObj.(string) //nolint: forcetypeassert // Ignored
	}

	return ""
}

func SetCorrelationIDInGin(c *gin.Context, correlationID string) {
	c.Set(correlationIDContextKey, correlationID)
}
