package log_test

import (
	"context"
	"testing"

	"github.com/wangbaolong/go-pkg/pkg/log"
)

func TestLog(t *testing.T) {
	cfg := &log.Config{
		Level:   "debug",
		File:    "logs/example.log",
		Format:  "json",
		Caller:  true,
		MaxSize: 1024,
		MaxDays: 1,
		Rotate:  true,
	}
	log.Init(cfg)
	log.Debug("example debug", log.Int("int", 12))
	log.Info("example info", log.String("string", "example"))
	log.Warn("example warn", log.Any("cfg", cfg))
	//log.Error("example error", log.Err(fmt.Errorf("example error : test")))
	// log.Panic("example panic", log.Err(fmt.Errorf("example panic : test")))
	//log.Fatal("example fatal", log.Err(fmt.Errorf("example fatal : test")))
}

func TestLogForCtx(t *testing.T) {
	cfg := &log.Config{
		Level:   "debug",
		File:    "logs/example.log",
		Format:  "console",
		Caller:  true,
		MaxSize: 1024,
		MaxDays: 1,
		Rotate:  true,
	}
	log.Init(cfg)
	ctx := context.WithValue(context.Background(), log.TraceIdKey, "123abc")
	log.For(ctx).Debug("example debug", log.String("string", "example"))
	log.For(ctx).Info("example info", log.String("string", "example"))
	log.For(ctx).Warn("example warn", log.Any("cfg", cfg))
	//log.For(ctx).Error("example error", log.Err(fmt.Errorf("example error : test")))
	//log.For(ctx).Panic("example panic", log.Err(fmt.Errorf("example panic : test")))
	//log.For(ctx).Fatal("example fatal", log.Err(fmt.Errorf("example fatal : test")))
}
