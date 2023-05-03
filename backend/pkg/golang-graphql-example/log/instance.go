package log

import (
	"os"
	"path/filepath"

	"emperror.dev/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	gormlogger "gorm.io/gorm/logger"
)

const LogFileMode = 0666
const LogTraceIDField = "trace_id"

func getInitConfig() *zap.Config {
	zconfig := zap.NewProductionConfig()
	zconfig.EncoderConfig.TimeKey = "time"
	zconfig.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	zconfig.DisableCaller = true

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

func (ll *loggerIns) GetGormLogger() gormlogger.Interface {
	return &gormLogger{
		logger: ll,
	}
}

func (ll *loggerIns) GetTracingLogger() TracingLogger {
	return &tracingLogger{
		logger: ll,
	}
}

func (ll *loggerIns) GetLockDistributorLogger() LockDistributorLogger {
	return &lockDistributorLogger{
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

	// Return new logger
	return &loggerIns{
		SugaredLogger: ll.SugaredLogger.With("error", err),
	}
}
