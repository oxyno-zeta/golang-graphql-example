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

// func (ll *loggerIns) addPotentialWithError(elem interface{}) logrus.FieldLogger {
// 	// Try to cast element to error
// 	err, ok := elem.(error)
// 	// Check if can be casted to error
// 	if ok {
// 		// Create new field logger
// 		fieldL := ll.FieldLogger.WithError(err)

// 		addStackTrace := func(pError stackTracer) {
// 			// Get stack trace from error
// 			st := pError.StackTrace()
// 			// Stringify stack trace
// 			valued := fmt.Sprintf("%+v", st)
// 			// Remove all tabs
// 			valued = strings.ReplaceAll(valued, "\t", "")
// 			// Split on new line
// 			stack := strings.Split(valued, "\n")
// 			// Remove first empty string
// 			stack = stack[1:]
// 			// Add stack trace to field logger
// 			fieldL = fieldL.WithField("stack", strings.Join(stack, ","))
// 		}

// 		// Check if error as an hidden stackTracer
// 		var st stackTracer
// 		if errors.As(err, &st) {
// 			addStackTrace(st)
// 		}

// 		return fieldL
// 	}

// 	// Default
// 	return ll.FieldLogger
// }

// func (ll *loggerIns) Error(args ...interface{}) {
// 	// Add potential "WithError"
// 	l := ll.addPotentialWithError(args[0])

// 	// Call logger error method
// 	l.Error(args...)
// }

// func (ll *loggerIns) Fatal(args ...interface{}) {
// 	// Add potential "WithError"
// 	l := ll.addPotentialWithError(args[0])

// 	// Call logger fatal method
// 	l.Fatal(args...)
// }

// func (ll *loggerIns) Errorf(format string, args ...interface{}) {
// 	// Create error
// 	err := errors.Errorf(format, args...)

// 	// Log error
// 	ll.Error(err)
// }

// func (ll *loggerIns) Fatalf(format string, args ...interface{}) {
// 	// Create error
// 	err := errors.Errorf(format, args...)

// 	// Log fatal
// 	ll.Fatal(err)
// }

// func (ll *loggerIns) Errorln(args ...interface{}) {
// 	// Add potential "WithError"
// 	l := ll.addPotentialWithError(args[0])

// 	// Log error
// 	l.Errorln(args...)
// }

// func (ll *loggerIns) Fatalln(args ...interface{}) {
// 	// Add potential "WithError"
// 	l := ll.addPotentialWithError(args[0])

// 	// Log fatal
// 	l.Fatalln(args...)
// }
