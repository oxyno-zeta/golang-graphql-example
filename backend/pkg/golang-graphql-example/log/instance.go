package log

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"emperror.dev/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const LogFileMode = 0666
const LogTraceIDField = "trace_id"

// This is dirty pkg/errors.
type stackTracer interface {
	StackTrace() errors.StackTrace
}

func getInitConfig() *zap.Config {
	zconfig := zap.NewProductionConfig()
	zconfig.EncoderConfig.TimeKey = "time"
	zconfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zconfig.DisableCaller = true
	zconfig.DisableStacktrace = true

	return &zconfig
}

func initLogger() *zap.SugaredLogger {
	zconfig := getInitConfig()

	logger, _ := zconfig.Build()

	// Create sugar logger
	suggarLogger := logger.Sugar()

	return suggarLogger
}

type loggerIns struct {
	*zap.SugaredLogger
}

func (ll *loggerIns) GetTracingLogger() TracingLogger {
	return &tracingLogger{
		logger: ll,
	}
}

func (ll *loggerIns) Configure(level string, format string, filePath string) error {
	// Get config
	zconfig := getInitConfig()

	// Parse log level
	lvl, err := zap.ParseAtomicLevel(level)
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}
	// Set log level
	zconfig.Level.SetLevel(lvl.Level())

	// Check format
	if format != "json" {
		zconfig.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		zconfig.Encoding = "console"
	}

	// Check if file path exists
	if filePath != "" {
		// Create directory if necessary
		err2 := os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
		// Check error
		if err2 != nil {
			return errors.WithStack(err2)
		}

		zconfig.OutputPaths = []string{filePath}
	}

	logger, err := zconfig.Build()
	// Check error
	if err != nil {
		return errors.WithStack(err)
	}

	ll.SugaredLogger = logger.Sugar()

	return nil
}

func (ll *loggerIns) WithField(key string, value interface{}) Logger {
	// Create new field logger
	fieldL := ll.SugaredLogger.With(key, value)

	return &loggerIns{
		SugaredLogger: fieldL,
	}
}

func (ll *loggerIns) WithFields(fields map[string]interface{}) Logger {
	// Create new field logger
	fieldL := ll.SugaredLogger
	// Loop over data
	for key, val := range fields {
		fieldL = fieldL.With(key, val)
	}

	return &loggerIns{
		SugaredLogger: fieldL,
	}
}

func (ll *loggerIns) WithError(err error) Logger {
	// Check if error doesn't exist
	if err == nil {
		return ll
	}

	// Create new field logger with error
	fieldL := ll.addPotentialWithError(err)
	// Return new logger
	return &loggerIns{
		SugaredLogger: fieldL.With("error", err),
	}
}

func (ll *loggerIns) addPotentialWithError(elem interface{}) *zap.SugaredLogger {
	// Try to cast element to error
	err, ok := elem.(error)
	// Check if can be casted to error
	if ok {
		// Check if error as an hidden stackTracer
		var st stackTracer
		if errors.As(err, &st) {
			// Get stack trace from error
			st := st.StackTrace()
			// Stringify stack trace
			valued := fmt.Sprintf("%+v", st)
			// Remove all tabs
			valued = strings.ReplaceAll(valued, "\t", "")
			// Split on new line
			stack := strings.Split(valued, "\n")
			// Remove first empty string
			stack = stack[1:]
			// Add stack trace to field logger
			return ll.SugaredLogger.With("stacktrace", strings.Join(stack, ","))
		}
	}

	// Default
	return ll.SugaredLogger
}

func (ll *loggerIns) Error(args ...interface{}) {
	// Add potential "WithError"
	l := ll.addPotentialWithError(args[0])

	// Call logger error method
	l.Error(args...)
}

func (ll *loggerIns) Fatal(args ...interface{}) {
	// Add potential "WithError"
	l := ll.addPotentialWithError(args[0])

	// Call logger fatal method
	l.Fatal(args...)
}

func (ll *loggerIns) Errorf(format string, args ...interface{}) {
	// Create error
	err := errors.WithStack(errors.Errorf(format, args...))

	// Log error
	ll.Error(err)
}

func (ll *loggerIns) Fatalf(format string, args ...interface{}) {
	// Create error
	err := errors.WithStack(errors.Errorf(format, args...))

	// Log fatal
	ll.Fatal(err)
}
