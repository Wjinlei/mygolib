package log

import "testing"

func TestNew(t *testing.T) {
	logger, err := New(&Option{
		LogLevel:     DebugLevel,
		LogType:      "json",
		ReportCaller: true,
	})
	if err != nil {
		t.Error(err)
		return
	}
	logger.Debug(Fields{"name": "wjl"}, "测试消息")
	logger.Info(Fields{"name": "wjl"}, "测试消息")
	logger.Warn(Fields{"name": "wjl"}, "测试消息")
	logger.Error(Fields{"name": "wjl"}, "测试消息")
	//logger.Fatal(Fields{"name": "wjl"}, "测试消息")
	//logger.Panic(Fields{"name": "wjl"}, "测试消息")
	//logger.Trace(Fields{"name": "wjl"}, "测试消息")
}
