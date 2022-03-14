package log

import (
	"context"
	"io"

	"go.uber.org/zap"
)

type Field = zap.Field
type ctxKey string

const TraceIdKey ctxKey = "trace-id"

var (
	writer      io.Writer
	zapLog      *zap.Logger
	sugar       *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

type Config struct {
	Level   string `yaml:"level"`   // debug/info/warn/error
	File    string `yaml:"file"`    // log file example: logs/example.log
	Format  string `yaml:"format"`  // console/json
	Caller  bool   `yaml:"caller"`  // whether print caller name
	MaxSize int    `yaml:"maxSize"` // log file max size
	MaxDays int    `yaml:"maxDays"` // max log file usage time
	Rotate  bool   `yaml:"rotate"`  // whether to compress to generate a new file
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

func Debug(msg string, fields ...Field) {
	zapLog.Debug(msg, fields...)
}

func Info(msg string, fields ...Field) {
	zapLog.Info(msg, fields...)
}

func Warn(msg string, fields ...Field) {
	zapLog.Warn(msg, fields...)
}

func Error(msg string, fields ...Field) {
	zapLog.Error(msg, fields...)
}

func Panic(msg string, fields ...Field) {
	zapLog.Panic(msg, fields...)
}

func Fatal(msg string, fields ...Field) {
	zapLog.Fatal(msg, fields...)
}

func Writer() io.Writer {
	return writer
}

func Sync() {
	_ = zapLog.Sync()
}
