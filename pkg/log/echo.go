package log

import (
	"io"

	"github.com/labstack/gommon/log"
	"go.uber.org/zap"
)

type EchoLogger struct {
	prefix string
	log    *zap.SugaredLogger
}

func NewEchoLogger(prefix string) *EchoLogger {
	return &EchoLogger{
		prefix: prefix,
		log:    sugar.Named(prefix),
	}
}

func (t *EchoLogger) Output() io.Writer {
	return Writer()
}

func (t *EchoLogger) SetOutput(w io.Writer) {
	//
}

func (t *EchoLogger) Prefix() string {
	return t.prefix
}

func (t *EchoLogger) SetPrefix(p string) {
	t.prefix = p
	t.log = sugar.Named(p)
}

func (t *EchoLogger) Level() log.Lvl {
	return log.INFO
}

func (t *EchoLogger) SetLevel(v log.Lvl) {
	// not work
}

func (t *EchoLogger) SetHeader(h string) {
	// not work
}

func (t *EchoLogger) Print(i ...interface{}) {
	t.log.Info(i...)
}

func (t *EchoLogger) Printf(format string, args ...interface{}) {
	t.log.Infof(format, args...)
}

func (t *EchoLogger) Printj(j log.JSON) {
	t.log.Info(j)
}

func (t *EchoLogger) Debug(i ...interface{}) {
	t.log.Debug(i...)
}

func (t *EchoLogger) Debugf(format string, args ...interface{}) {
	t.log.Debugf(format, args...)
}

func (t *EchoLogger) Debugj(j log.JSON) {
	t.log.Debug(j)
}

func (t *EchoLogger) Info(i ...interface{}) {
	t.log.Info(i...)
}

func (t *EchoLogger) Infof(format string, args ...interface{}) {
	t.log.Infof(format, args...)
}

func (t *EchoLogger) Infoj(j log.JSON) {
	t.log.Info(j)
}

func (t *EchoLogger) Warn(i ...interface{}) {
	t.log.Warn(i...)
}

func (t *EchoLogger) Warnf(format string, args ...interface{}) {
	t.log.Warnf(format, args...)
}

func (t *EchoLogger) Warnj(j log.JSON) {
	t.log.Warn(j)
}

func (t *EchoLogger) Error(i ...interface{}) {
	t.log.Error(i...)
}

func (t *EchoLogger) Errorf(format string, args ...interface{}) {
	t.log.Errorf(format, args...)
}

func (t *EchoLogger) Errorj(j log.JSON) {
	t.log.Error(j)
}

func (t *EchoLogger) Fatal(i ...interface{}) {
	t.log.Fatal(i...)
}

func (t *EchoLogger) Fatalj(j log.JSON) {
	t.log.Fatal(j)
}

func (t *EchoLogger) Fatalf(format string, args ...interface{}) {
	t.log.Fatalf(format, args...)
}

func (t *EchoLogger) Panic(i ...interface{}) {
	t.log.Panic(i...)
}

func (t *EchoLogger) Panicj(j log.JSON) {
	t.log.Panic(j)
}

func (t *EchoLogger) Panicf(format string, args ...interface{}) {
	t.log.Panicf(format, args...)
}
