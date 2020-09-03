package logger

import (
	"fmt"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	logrus "github.com/sirupsen/logrus"
)

// Option 选项定义
type Option struct {
	LogPath      string        // 日志路径
	LogLevel     Level         // 日志级别
	LogType      string        // 日志类型: json, text
	MaxAge       time.Duration // 日志文件清理前的最长保存时间
	RotationTime time.Duration // 日志文件多长时间清理(切割)一次
}

// Logger 日志结构体
type Logger struct {
	logger *logrus.Logger
}

// Fields 自定义字段
type Fields map[string]interface{}

// Level 日志级别类型
type Level uint32

// 日志级别枚举
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
	DefaultDateFormat  = "%Y%m%d%H%M"          // 默认日期格式
	DefaultTimeFormat  = "2006-01-02 15:04:05" // 默认时间格式
	DefaultDataKey     = "data"                // 附加字段都会作为该字段的嵌入字段
	DefaultPrettyPrint = false                 // 默认美化输出
)

var (
	logger *Logger
)

// GetLogger 返回logger
func GetLogger() *Logger {
	return logger
}

// New 产生新的logger
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
	// 将路径转换为绝对路径
	absLogPath, err := filepath.Abs(option.LogPath)
	if err != nil {
		return nil, err
	}
	rotator, err := rotatelogs.New(
		fmt.Sprintf("%s-%s", absLogPath, DefaultDateFormat),
		rotatelogs.WithLinkName(absLogPath),
		rotatelogs.WithMaxAge(option.MaxAge),             // 日志文件清理前的最长保存时间
		rotatelogs.WithRotationTime(option.RotationTime), // 多久滚动一次
	)
	if err != nil {
		return nil, err
	}
	logrusLogger.SetOutput(rotator)
	logger = &Logger{
		logger: logrusLogger,
	}
	return logger, nil
}

// Debug 打印Debug日志
func (l *Logger) Debug(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Debug(message)
}

// Info 打印Info日志
func (l *Logger) Info(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Info(message)
}

// Warn 打印Warn日志
func (l *Logger) Warn(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Warn(message)
}

// Error 打印Error日志
func (l *Logger) Error(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Error(message)
}

// Fatal 打印Fatal日志
func (l *Logger) Fatal(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Fatal(message)
}

// Panic 打印Panic日志
func (l *Logger) Panic(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Panic(message)
}

// Trace 打印Trace日志
func (l *Logger) Trace(fields Fields, message string) {
	if fields == nil {
		fields = Fields{}
	}
	l.logger.WithFields(logrus.Fields(fields)).Trace(message)
}

// Debug 全局方式调用Debug函数
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

// Info 全局方式调用Info函数
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

// Warn 全局方式调用Warn函数
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

// Error 全局方式调用Error函数
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

// Fatal 全局方式调用Fatal函数
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

// Panic 全局方式调用Panic函数
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

// Trace 全局方式调用Trace函数
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
