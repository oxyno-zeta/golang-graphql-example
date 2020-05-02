package middlewares

import (
	"net/http"

	"github.com/gin-gonic/gin"
	uuid "github.com/gofrs/uuid"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
)

const RequestIDHeader = "X-Request-Id"
const RequestIDContextKey = "RequestID"

func RequestID(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request id from request
		requestID := c.Request.Header.Get(RequestIDHeader)

		// Check if request id exists
		if requestID == "" {
			// Generate uuid
			uuid, err := uuid.NewV4()
			if err != nil {
				// Log error
				logger.Errorln(err)
				// Send response
				c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

				return
			}
			// Save it in variable
			requestID = uuid.String()
		}

		// Store it in context
		c.Set(RequestIDContextKey, requestID)

		// Put it on header
		c.Writer.Header().Set(RequestIDHeader, requestID)

		// Next
		c.Next()
	}
}
