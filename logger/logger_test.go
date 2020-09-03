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
	})
	if err != nil {
		t.Error(err)
		return
	}
	for {
		logger.Debug(nil, "测试消息", Fields{"name": "wjl", "age": 24})
		time.Sleep(time.Duration(2) * time.Second)
	}
}
