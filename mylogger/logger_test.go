package mylogger

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	logger, err := NewLogger(&Option{
		LogPath:      "./log/test.log",
		LogLevel:     DebugLevel,
		LogType:      "json",
		MaxAge:       time.Duration(30) * time.Second,
		RotationTime: time.Duration(10) * time.Second,
		PrettyPrint:  false,
	})
	if err != nil {
		t.Error(err)
		return
	}
	for {
		logger.Info("测试消息", Fields{"name": "wjl"})
		time.Sleep(time.Duration(2) * time.Second)
	}
}
