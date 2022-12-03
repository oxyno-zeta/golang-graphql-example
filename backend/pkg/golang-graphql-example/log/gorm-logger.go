package log

import (
	"context"
	"time"

	"emperror.dev/errors"
	gormlogger "gorm.io/gorm/logger"
)

type gormLogger struct {
	logger Logger
}

func (gl *gormLogger) LogMode(gormlogger.LogLevel) gormlogger.Interface {
	return gl
}

func (gl *gormLogger) getCtxLoggerOrDefault(ctx context.Context) Logger {
	// Get logger from context
	ctxL := GetLoggerFromContext(ctx)
	// Check if it exists
	if ctxL != nil {
		return ctxL
	}

	// Default
	return gl.logger
}

func (gl *gormLogger) Info(ctx context.Context, v string, rest ...interface{}) {
	val := []interface{}{v}
	val = append(val, rest...)
	gl.getCtxLoggerOrDefault(ctx).Info(val...)
}

func (gl *gormLogger) Warn(ctx context.Context, v string, rest ...interface{}) {
	val := []interface{}{v}
	val = append(val, rest...)
	gl.getCtxLoggerOrDefault(ctx).Warn(val...)
}

func (gl *gormLogger) Error(ctx context.Context, v string, rest ...interface{}) {
	val := []interface{}{errors.New(v)}
	val = append(val, rest...)
	gl.getCtxLoggerOrDefault(ctx).Error(val...)
}

func (gl *gormLogger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	gl.getCtxLoggerOrDefault(ctx).WithField("rows", rows).WithError(err).Debug(sql)
}
