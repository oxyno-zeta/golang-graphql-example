package middlewares

import (
	"context"

	"github.com/gin-gonic/gin"
	uuid "github.com/gofrs/uuid"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/utils"
	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/log"
	"github.com/pkg/errors"
)

type contextKey struct {
	name string
}

var reqCtxKey = &contextKey{name: "request-id"}

const requestIDHeader = "X-Request-Id"
const requestIDContextKey = "RequestID"

func RequestID(logger log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get request id from request
		requestID := c.Request.Header.Get(requestIDHeader)

		// Check if request id exists
		if requestID == "" {
			// Generate uuid
			uuid, err := uuid.NewV4()
			// Check error
			if err != nil {
				// Add stack trace to error
				err2 := errors.WithStack(err)
				// Log error
				logger.Error(err2)
				// Send response
				utils.AnswerWithError(c, err2)

				return
			}
			// Save it in variable
			requestID = uuid.String()
		}

		// Store it in context
		SetRequestIDInGin(c, requestID)
		// Update request with new context
		c.Request = c.Request.WithContext(SetRequestIDInContext(c.Request.Context(), requestID))

		// Put it on header
		c.Writer.Header().Set(requestIDHeader, requestID)

		// Next
		c.Next()
	}
}

func GetRequestIDFromGin(c *gin.Context) string {
	requestIDObj, requestIDExists := c.Get(requestIDContextKey)
	if requestIDExists {
		// return request id
		return requestIDObj.(string) // nolint: forcetypeassert // Ignored
	}

	return ""
}

func GetRequestIDFromContext(ctx context.Context) string {
	requestIDObj := ctx.Value(reqCtxKey)
	if requestIDObj != nil {
		return requestIDObj.(string) // nolint: forcetypeassert // Ignored
	}

	return ""
}

func SetRequestIDInContext(ctx context.Context, requestID string) context.Context {
	return context.WithValue(ctx, reqCtxKey, requestID)
}

func SetRequestIDInGin(c *gin.Context, requestID string) {
	c.Set(requestIDContextKey, requestID)
}
