package mylogger

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	logger, err := GetLogger(&Option{
		LogPath:      "./log/test.log",
		LogLevel:     DebugLevel,
		LogType:      "json",
		MaxAge:       time.Duration(180) * time.Second,
		RotationSize: 1024,
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
