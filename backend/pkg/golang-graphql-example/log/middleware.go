package log

import (
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	gqlgraphql "github.com/99designs/gqlgen/graphql"

	"github.com/oxyno-zeta/golang-graphql-example/pkg/golang-graphql-example/common/utils"
)

type logContextKey struct {
	name string
}

const loggerGinCtxKey = "LoggerCtxKey"

var loggerCtxKey = &logContextKey{name: "logger"}

const nsToMs = 1000000.0

func GetLoggerFromContext(ctx context.Context) Logger {
	logger, _ := ctx.Value(loggerCtxKey).(Logger)

	return logger
}

func GetLoggerFromGin(c *gin.Context) Logger {
	val, _ := c.Get(loggerGinCtxKey)
	logger, _ := val.(Logger)

	return logger
}

func SetLoggerToContext(ctx context.Context, logger Logger) context.Context {
	return context.WithValue(ctx, loggerCtxKey, logger)
}

func SetLoggerToGin(c *gin.Context, logger Logger) {
	c.Set(loggerGinCtxKey, logger)
}

func Middleware(
	logger Logger,
	getCorrelationID func(c *gin.Context) string,
	getTraceID func(ctx context.Context) string,
) gin.HandlerFunc {
	return func(c *gin.Context) {
		t1 := time.Now()
		// Get request
		r := c.Request

		// Create logger fields
		logFields := make(map[string]any)

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

		logFields["uri"] = utils.GetRequestURL(c.Request)

		// Log correlation id
		logFields["correlation_id"] = getCorrelationID(c)

		// Get trace id
		traceID := getTraceID(c.Request.Context())
		if traceID != "" {
			logFields[LogTraceIDField] = traceID
		}

		requestLogger := logger.WithFields(logFields)

		requestLogger.Debug("request started")

		// Add logger to request
		SetLoggerToGin(c, requestLogger)
		c.Request = c.Request.WithContext(SetLoggerToContext(c.Request.Context(), requestLogger))

		// Next
		c.Next()

		// Get status
		status := c.Writer.Status()
		bytes := c.Writer.Size()

		// Create new fields
		endFields := map[string]any{
			"resp_status":       status,
			"resp_bytes_length": bytes,
			"resp_elapsed_ms":   float64(time.Since(t1).Nanoseconds()) / nsToMs,
		}

		endRequestLogger := requestLogger.WithFields(endFields)

		logFunc := endRequestLogger.Info

		if status >= http.StatusMultipleChoices && status < http.StatusInternalServerError {
			logFunc = endRequestLogger.Warn
		}

		if status >= http.StatusInternalServerError {
			logFunc = endRequestLogger.Error
		}

		logFunc("request complete")
	}
}

func GraphqlAroundRootFieldsMiddleware() func(ctx context.Context, next gqlgraphql.RootResolver) gqlgraphql.Marshaler {
	return func(ctx context.Context, next gqlgraphql.RootResolver) gqlgraphql.Marshaler {
		rc := gqlgraphql.GetRootFieldContext(ctx)
		oc := gqlgraphql.GetOperationContext(ctx)

		opName := "no-operation-context"
		opType := "no-operation-context"

		if oc != nil {
			opName = oc.OperationName
			if opName == "" {
				opName = "Anonymous"
			}

			opType = string(oc.Operation.Operation)
		}

		lInitFields := map[string]any{
			"graphql_operation_name": opName,
			"graphql_operation_type": opType,
		}

		if rc != nil {
			lInitFields["graphql_root_object_name"] = rc.Object
			lInitFields["graphql_root_object_field_name"] = rc.Field.Name
		}

		logger := GetLoggerFromContext(ctx).WithFields(lInitFields)
		ctx = SetLoggerToContext(ctx, logger)

		start := time.Now()

		logger.Debug("Starting Root GraphQL operation")

		res := next(ctx)

		if rc != nil && rc.Object != "__Query" && rc.Object != "__Type" && rc.Object != "__Schema" {
			duration := time.Since(start)
			logger = logger.WithField("graphql_operation_elapsed_ms", float64(duration.Nanoseconds())/nsToMs)

			errs := gqlgraphql.GetErrors(ctx)
			if len(errs) > 0 {
				logger.Errorf("Failed Root GraphQL operation")

				for _, err := range errs {
					logger.Error(err)
				}
			} else {
				logger.Info("Success Root GraphQL operation")
			}
		}

		return res
	}
}

func GraphqlAroundFieldsMiddleware() func(ctx context.Context, next gqlgraphql.Resolver) (any, error) {
	return func(ctx context.Context, next gqlgraphql.Resolver) (any, error) {
		fc := gqlgraphql.GetFieldContext(ctx)

		lInitFields := map[string]any{}
		if fc != nil {
			lInitFields["graphql_object_name"] = fc.Object
			lInitFields["graphql_field_name"] = fc.Field.Name
			lInitFields["graphql_field_path"] = fc.Path().String()
		}

		logger := GetLoggerFromContext(ctx).WithFields(lInitFields)
		ctx = SetLoggerToContext(ctx, logger)

		start := time.Now()

		logger.Debug("Starting Field GraphQL operation")

		res, err := next(ctx)

		if fc != nil && fc.Object != "__Query" && fc.Object != "__Type" && fc.Object != "__Schema" {
			duration := time.Since(start)
			logger = logger.WithField("graphql_field_elapsed_ms", float64(duration.Nanoseconds())/nsToMs)

			if err != nil {
				logger.Error("Failed Field GraphQL operation")
				logger.Error(err)
			} else {
				logger.Debug("Success Field GraphQL operation")
			}
		}

		return res, err
	}
}
