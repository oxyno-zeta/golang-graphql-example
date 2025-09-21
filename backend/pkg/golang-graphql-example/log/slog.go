package log

import (
	"context"
	"log/slog"

	"go.uber.org/zap/zapcore"

	slogzap "github.com/samber/slog-zap/v2"
)

type SlogHandler struct {
	*slogzap.ZapHandler
	*loggerIns
}

func (ll *loggerIns) GetSlogInstance() *slog.Logger {
	zapHandler := slogzap.Option{
		// Fake level
		Level:  slog.LevelDebug,
		Logger: ll.Desugar(),
	}.NewZapHandler().(*slogzap.ZapHandler)

	slHandler := &SlogHandler{
		ZapHandler: zapHandler,
		loggerIns:  ll,
	}

	return slog.New(slHandler)
}

func (h *SlogHandler) Enabled(_ context.Context, level slog.Level) bool {
	return level == h.getSLogLevel()
}

func (h *SlogHandler) getSLogLevel() slog.Level {
	switch h.Level() {
	case zapcore.PanicLevel:
		fallthrough
	case zapcore.FatalLevel:
		fallthrough
	case zapcore.ErrorLevel:
		return slog.LevelError
	case zapcore.WarnLevel:
		return slog.LevelWarn
	case zapcore.InfoLevel:
		return slog.LevelInfo
	case zapcore.DebugLevel:
		return slog.LevelDebug
	default:
		return slog.LevelInfo
	}
}
