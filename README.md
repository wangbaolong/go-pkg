```shell
go get github.com/wangbaolong/go-pkg
```
#### Example 
##### Log
```go
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
```
###### 日志输出
```text
{"level":"debug","ts":"2022-01-19T11:46:59.151+0800","caller":"log/log_test.go:21","msg":"example debug","int":12}
{"level":"info","ts":"2022-01-19T11:46:59.151+0800","caller":"log/log_test.go:22","msg":"example info","string":"example"}
{"level":"warn","ts":"2022-01-19T11:46:59.152+0800","caller":"log/log_test.go:23","msg":"example warn","cfg":{"Level":"debug","File":"logs/example.log","Format":"json","Caller":true,"MaxSize":1024,"MaxDays":1,"Rotate":true}}
```

##### Log for context
```go
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
```

###### 日志输出
```text
2022-01-19T11:49:11.144+0800	debug	log/log_test.go:41	example debug	{"string": "example", "x-trace-id": "123abc"}
2022-01-19T11:49:11.145+0800	info	log/log_test.go:42	example info	{"string": "example", "x-trace-id": "123abc"}
2022-01-19T11:49:11.145+0800	warn	log/log_test.go:43	example warn	{"cfg": {"Level":"debug","File":"logs/example.log","Format":"console","Caller":true,"MaxSize":1024,"MaxDays":1,"Rotate":true}, "x-trace-id": "123abc"}
```

