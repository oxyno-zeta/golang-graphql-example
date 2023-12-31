package log

import (
	"context"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

type Logger interface {
	Configure(level string, format string, filePath string) error
	// This must be called in the main function to force flush and write.
	Sync() error

	WithField(key string, value interface{}) Logger
	WithFields(fields map[string]interface{}) Logger
	WithError(err error) Logger

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
	Fatalf(format string, args ...interface{})
	Panicf(format string, args ...interface{})

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})
	Fatal(args ...interface{})
	Panic(args ...interface{})

	GetLockDistributorLogger() LockDistributorLogger
	GetGormLogger() gormlogger.Interface
}

type LockDistributorLogger interface {
	Println(args ...interface{})
}

type GormLogger interface {
	LogMode(lvl gormlogger.LogLevel) GormLogger
	Info(ctx context.Context, msg string, args ...interface{})
	Warn(ctx context.Context, msg string, args ...interface{})
	Error(ctx context.Context, msg string, args ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

func NewLogger() Logger {
	logger := initLogger()

	return &loggerIns{
		SugaredLogger: logger,
	}
}
