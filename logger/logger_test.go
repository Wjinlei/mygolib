package logger

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	logger, err := New(&Option{
		LogPath:      "./log/test.log",
		LogLevel:     DebugLevel,
		LogType:      "json",
		MaxAge:       time.Duration(180) * time.Second,
		RotationTime: time.Duration(60) * time.Second,
		ReportCaller: true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for {
		logger.Debug(Fields{"name": "wjl"}, "测试消息")
		logger.Info(Fields{"name": "wjl"}, "测试消息")
		logger.Warn(Fields{"name": "wjl"}, "测试消息")
		//logger.Error(Fields{"name": "wjl"}, "测试消息")
		//logger.Fatal(Fields{"name": "wjl"}, "测试消息")
		//logger.Panic(Fields{"name": "wjl"}, "测试消息")
		//logger.Trace(Fields{"name": "wjl"}, "测试消息")
		time.Sleep(time.Duration(2) * time.Second)
	}
}
