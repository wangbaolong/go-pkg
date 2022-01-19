package log

import (
	"go.uber.org/zap"
)

type DBLogger struct {
	prefix string
	log    *zap.SugaredLogger
}

func NewGormLogger(prefix string) *DBLogger {
	return &DBLogger{
		prefix: prefix,
		log:    sugar.Named(prefix),
	}
}

func (t *DBLogger) Printf(format string, args ...interface{}) {
	format = "%s [%.3fms] [rows:%v] %s"
	t.log.Infof(format, args...)
}
