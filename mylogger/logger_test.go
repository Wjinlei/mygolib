package mylogger

import (
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	logger, err := NewLogger(&Option{
		LogPath:        "log/test.log",
		LogLevel:       DebugLevel,
		LogType:        "text",
		RotationMaxAge: time.Duration(180) * time.Second, // 保留最近3分钟的日志
		RotationTime:   time.Duration(60) * time.Second,  //每1分钟切割一次
		PrettyPrint:    false,
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
