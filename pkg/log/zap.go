package log

import (
	"log"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func NewLoggerWithConfig(cfg *Config) *zap.Logger {
	var err error

	ws := make([]zapcore.WriteSyncer, 0, 2)
	ws = append(ws, zapcore.AddSync(os.Stdout))
	if cfg.File != "" {
		rotateLogger := &lumberjack.Logger{
			Filename:  cfg.File,
			MaxSize:   cfg.MaxSize,
			MaxAge:    cfg.MaxDays,
			LocalTime: true,
			Compress:  true,
		}
		ws = append(ws, zapcore.AddSync(rotateLogger))

		if cfg.Rotate {
			go scheduleRotate(rotateLogger)
		}
	}

	var level zapcore.Level
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		log.Panic(err)
	}

	atomicLevel = zap.NewAtomicLevelAt(level)

	writeSynced := zapcore.NewMultiWriteSyncer(ws...)
	writer = writeSynced

	encodingCfg := zap.NewProductionEncoderConfig()
	encodingCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if strings.ToLower(cfg.Format) == "json" {
		encoder = zapcore.NewJSONEncoder(encodingCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encodingCfg)
	}
	core := zapcore.NewCore(
		encoder,
		writeSynced,
		atomicLevel,
	)

	options := make([]zap.Option, 0, 3)
	options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	if cfg.Caller && level.Enabled(zapcore.DebugLevel) {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	lg := zap.New(core, options...)

	return lg
}

func scheduleRotate(log *lumberjack.Logger) {
	for {
		n := time.Now().Add(time.Hour * 24)
		next := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
		d := time.Until(next)
		time.Sleep(d)
		_ = log.Rotate()
	}
}
