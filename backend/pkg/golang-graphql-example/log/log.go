package log

import (
	"context"
	"time"

	gormlogger "gorm.io/gorm/logger"
)

type Logger interface {
	Configure(level, format, filePath string) error
	// This must be called in the main function to force flush and write.
	Sync() error

	WithField(key string, value any) Logger
	WithFields(fields map[string]any) Logger
	WithError(err error) Logger

	Debugf(format string, args ...any)
	Infof(format string, args ...any)
	Warnf(format string, args ...any)
	Errorf(format string, args ...any)
	Fatalf(format string, args ...any)
	Panicf(format string, args ...any)

	Debug(args ...any)
	Info(args ...any)
	Warn(args ...any)
	Error(args ...any)
	Fatal(args ...any)
	Panic(args ...any)

	GetLockDistributorLogger() LockDistributorLogger
	GetGormLogger() gormlogger.Interface
}

type LockDistributorLogger interface {
	Println(args ...any)
}

type GormLogger interface {
	LogMode(lvl gormlogger.LogLevel) GormLogger
	Info(ctx context.Context, msg string, args ...any)
	Warn(ctx context.Context, msg string, args ...any)
	Error(ctx context.Context, msg string, args ...any)
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

func NewLogger() Logger {
	logger := initLogger()

	return &loggerIns{
		SugaredLogger: logger,
	}
}
