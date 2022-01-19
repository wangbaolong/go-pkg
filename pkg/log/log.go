package log

import (
	"context"
	"io"

	"go.uber.org/zap"
)

type Field = zap.Field
type ctxKey string

const TraceIdKey ctxKey = "traceId"

var (
	writer      io.Writer
	zapLog      *zap.Logger
	sugar       *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

type Config struct {
	Level   string `yaml:"level"`
	File    string `yaml:"file"`
	Format  string `yaml:"format"`
	Caller  bool   `yaml:"caller"`
	MaxSize int    `yaml:"maxSize"`
	MaxDays int    `yaml:"maxDays"`
	Rotate  bool   `yaml:"rotate"`
}

func Init(cfg *Config) {
	zapLog = NewLoggerWithConfig(cfg)
	sugar = zapLog.Sugar()
}

type trace interface {
	Debug(msg string, fields ...Field)
	Info(msg string, fields ...Field)
	Warn(msg string, fields ...Field)
	Error(msg string, fields ...Field)
	Panic(msg string, fields ...Field)
	Fatal(msg string, fields ...Field)
}

type traceImpl struct {
	traceId string
}

func (t *traceImpl) Debug(msg string, fields ...Field) {
	if len(t.traceId) > 0 {
		fields = append(fields, zap.String("x-trace-id", t.traceId))
	}
	zapLog.Debug(msg, fields...)
}

func (t *traceImpl) Info(msg string, fields ...Field) {
	if len(t.traceId) > 0 {
		fields = append(fields, zap.String("x-trace-id", t.traceId))
	}
	zapLog.Info(msg, fields...)
}

func (t *traceImpl) Warn(msg string, fields ...Field) {
	if len(t.traceId) > 0 {
		fields = append(fields, zap.String("x-trace-id", t.traceId))
	}
	zapLog.Warn(msg, fields...)
}

func (t *traceImpl) Error(msg string, fields ...Field) {
	if len(t.traceId) > 0 {
		fields = append(fields, zap.String("x-trace-id", t.traceId))
	}
	zapLog.Error(msg, fields...)
}

func (t *traceImpl) Panic(msg string, fields ...Field) {
	if len(t.traceId) > 0 {
		fields = append(fields, zap.String("x-trace-id", t.traceId))
	}
	zapLog.Panic(msg, fields...)
}

func (t *traceImpl) Fatal(msg string, fields ...Field) {
	if len(t.traceId) > 0 {
		fields = append(fields, zap.String("x-trace-id", t.traceId))
	}
	zapLog.Fatal(msg, fields...)
}

func newTrace(traceId string) trace {
	return &traceImpl{traceId: traceId}
}

func parseTraceId(ctx context.Context) string {
	if ctx != nil {
		value := ctx.Value(TraceIdKey)
		if traceId, ok := value.(string); ok {
			return traceId
		}
	}
	return ""
}

func For(ctx context.Context) trace {
	return newTrace(parseTraceId(ctx))
}

func Debug(msg string, fields ...zap.Field) {
	zapLog.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	zapLog.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	zapLog.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	zapLog.Error(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	zapLog.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	zapLog.Fatal(msg, fields...)
}

func Writer() io.Writer {
	return writer
}

func Sync() {
	_ = zapLog.Sync()
}
