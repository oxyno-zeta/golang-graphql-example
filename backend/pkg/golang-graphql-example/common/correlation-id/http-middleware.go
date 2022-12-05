package correlationid

import (
	"github.com/gin-gonic/gin"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/utils"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

const correlationIDContextKey = "correlationID"

func HTTPMiddleware(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get correlation id from request
		correlationID := c.Request.Header.Get(correlationIDHeader)

		// Check if correlation id header have been set
		if correlationID == "" {
			// Get request id header
			correlationID = c.Request.Header.Get(requestIDHeader)
		}

		// Check if correlation id exists
		if correlationID == "" {
			// Generate uuid
			uuid, err := Generate()
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
		SetInGin(c, correlationID)
		// Update request with new context
		c.Request = c.Request.WithContext(
			SetInContext(c.Request.Context(), correlationID),
		)

		// Put it on header
		c.Writer.Header().Set(requestIDHeader, correlationID)

		// Next
		c.Next()
	}
}

func GetFromGin(c *gin.Context) string {
	correlationIDObj, correlationIDExists := c.Get(correlationIDContextKey)
	if correlationIDExists {
		// return request id
		return correlationIDObj.(string) //nolint: forcetypeassert // Ignored
	}

	return ""
}

func SetInGin(c *gin.Context, correlationID string) {
	c.Set(correlationIDContextKey, correlationID)
}
