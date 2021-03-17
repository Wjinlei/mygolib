package mylogger

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	logger, err := NewLogger(&Option{
		LogPath:       "log/test.log",
		LogLevel:      DebugLevel,
		LogType:       "text",
		MaxAge:        time.Duration(30) * time.Second,
		RotationCount: 5,
		RotationSize:  100,
		PrettyPrint:   false,
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
