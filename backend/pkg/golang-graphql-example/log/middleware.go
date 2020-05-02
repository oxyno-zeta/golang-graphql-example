package log

import (
	"context"
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

type logContextKey struct {
	name string
}

const loggerGinCtxKey = "LoggerCtxKey"

var loggerCtxKey = &logContextKey{name: "logger"}

func Middleware(logger Logger, getRequestID func(c *gin.Context) string, getSpanID func(ctx context.Context) string) gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		// Get request
		r := c.Request

		// Create logger fields
		logFields := make(map[string]interface{})

		// Check if it is http or https
		scheme := "http"
		if r.TLS != nil {
			scheme = "https"
		}

		logFields["http_scheme"] = scheme
		logFields["http_proto"] = r.Proto
		logFields["http_method"] = r.Method

		logFields["remote_addr"] = r.RemoteAddr
		logFields["user_agent"] = r.UserAgent()
		logFields["client_ip"] = c.ClientIP()

		logFields["uri"] = fmt.Sprintf("%s://%s%s", scheme, r.Host, r.RequestURI)

		// Log request id
		logFields["request_id"] = getRequestID(c)

		// Get span id
		spanID := getSpanID(c.Request.Context())
		if spanID != "" {
			logFields["span_id"] = spanID
		}

		requestLogger := logger.WithFields(logFields)

		requestLogger.Debugln("request started")

		// Add logger to request
		c.Set(loggerGinCtxKey, requestLogger)
		c.Request = c.Request.WithContext(context.WithValue(c.Request.Context(), loggerCtxKey, requestLogger))

		// Next
		c.Next()

		// Get status
		status := c.Writer.Status()
		bytes := c.Writer.Size()

		// Create new fields
		endFields := map[string]interface{}{
			"resp_status":       status,
			"resp_bytes_length": bytes,
			"resp_elapsed_ms":   float64(time.Since(t1).Nanoseconds()) / 1000000.0,
		}

		endRequestLogger := requestLogger.WithFields(endFields)

		logFunc := endRequestLogger.Infoln

		if status >= 300 && status < 400 {
			logFunc = endRequestLogger.Warnln
		}

		if status >= 400 {
			logFunc = endRequestLogger.Errorln
		}

		logFunc("request complete")
	}
}
