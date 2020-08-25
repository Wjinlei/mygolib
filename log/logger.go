package log

import (
	logrus "github.com/sirupsen/logrus"
)

// 选项定义
type Option struct {
	LogLevel     Level
	LogType      string // 日志类型: json, text
	ReportCaller bool   // 将调用方法添加为字段,这会带来开销
}

type Logger struct {
	logger *logrus.Logger
}

type Fields map[string]interface{}

type Level uint32

// 日志级别
const (
	PanicLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)

// 默认选项
const (
	DefaultDateFormat  = "%Y%m%d"              // 默认日期格式
	DefaultTimeFormat  = "2006-01-02 15:04:05" // 默认时间格式
	DefaultDataKey     = "data"                // 附加字段都会作为该字段的嵌入字段
	DefaultPrettyPrint = true                  // 默认美化输出
)

var (
	logger *Logger
)

// 返回logger
func GetLogger() *Logger {
	return logger
}

// 产生新的logger
func New(option *Option) (*Logger, error) {
	logrusLogger := logrus.New()
	// 设置日志格式
	switch option.LogType {
	case "json":
		logrusLogger.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: DefaultTimeFormat,
			DataKey:         DefaultDataKey,
			PrettyPrint:     DefaultPrettyPrint,
		})
	default:
		logrusLogger.SetFormatter(&logrus.TextFormatter{
			TimestampFormat: DefaultTimeFormat,
		})
	}
	logrusLogger.Level = logrus.Level(option.LogLevel)
	logrusLogger.SetReportCaller(option.ReportCaller)
	logger = &Logger{
		logger: logrusLogger,
	}
	return logger, nil
}

// 打印Debug日志
func (l *Logger) Debug(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Debug(message)
}

// 打印Info日志
func (l *Logger) Info(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Info(message)
}

// 打印Warn日志
func (l *Logger) Warn(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Warn(message)
}

// 打印Error日志
func (l *Logger) Error(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Error(message)
}

// 打印Fatal日志
func (l *Logger) Fatal(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Fatal(message)
}

// 打印Panic日志
func (l *Logger) Panic(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Panic(message)
}

// 打印Trace日志
func (l *Logger) Trace(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Trace(message)
}

// 全局方式调用Debug函数
func Debug(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	if logger.logger == nil {
		logrus.Debug(fields, message)
		return
	}
	logger.logger.WithFields(logrus.Fields(fields)).Debug(message)
}

// 全局方式调用Info函数
func Info(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	if logger.logger == nil {
		logrus.Info(fields, message)
		return
	}
	logger.logger.WithFields(logrus.Fields(fields)).Info(message)
}

// 全局方式调用Warn函数
func Warn(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	if logger.logger == nil {
		logrus.Warn(fields, message)
		return
	}
	logger.logger.WithFields(logrus.Fields(fields)).Warn(message)
}

// 全局方式调用Error函数
func Error(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	if logger.logger == nil {
		logrus.Error(fields, message)
		return
	}
	logger.logger.WithFields(logrus.Fields(fields)).Error(message)
}

// 全局方式调用Fatal函数
func Fatal(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	if logger.logger == nil {
		logrus.Fatal(fields, message)
		return
	}
	logger.logger.WithFields(logrus.Fields(fields)).Fatal(message)
}

// 全局方式调用Panic函数
func Panic(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	if logger.logger == nil {
		logrus.Panic(fields, message)
		return
	}
	logger.logger.WithFields(logrus.Fields(fields)).Panic(message)
}

// 全局方式调用Trace函数
func Trace(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	if logger.logger == nil {
		logrus.Trace(fields, message)
		return
	}
	logger.logger.WithFields(logrus.Fields(fields)).Trace(message)
}
