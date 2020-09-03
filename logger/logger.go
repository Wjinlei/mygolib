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
	LogLevel     LogLevel      // 日志级别
	LogType      string        // 日志类型: json, text
	MaxAge       time.Duration // 日志文件清理前的最长保存时间
	RotationTime time.Duration // 日志文件多长时间清理(切割)一次
	PrettyPrint  bool          // 美化输出
}

// Logger 日志结构体
type Logger struct {
	logger *logrus.Logger
}

// Fields 自定义字段
type Fields map[string]interface{}

// LogLevel 日志级别类型
type LogLevel uint32

// 日志级别枚举
const (
	TraceLevel LogLevel = iota
	PanicLevel
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

// 默认选项
const (
	DefaultDateFormat = "%Y%m%d%H%M"          // 默认日期格式
	DefaultTimeFormat = "2006-01-02 15:04:05" // 默认时间格式
	DefaultDataKey    = "data"                // 附加字段都会作为该字段的嵌入字段
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
			PrettyPrint:     option.PrettyPrint,
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
func (l *Logger) Debug(message string, data ...interface{}) {
	l.logger.WithField("Fields", data).Debug(message)
}

// Info 打印Info日志
func (l *Logger) Info(message string, data ...interface{}) {
	l.logger.WithField("Fields", data).Info(message)
}

// Warn 打印Warn日志
func (l *Logger) Warn(message string, data ...interface{}) {
	l.logger.WithField("Fields", data).Warn(message)
}

// Error 打印Error日志
func (l *Logger) Error(message string, data ...interface{}) {
	l.logger.WithField("Fields", data).Error(message)
}

// Fatal 打印Fatal日志
func (l *Logger) Fatal(message string, data ...interface{}) {
	l.logger.WithField("Fields", data).Fatal(message)
}

// Panic 打印Panic日志
func (l *Logger) Panic(message string, data ...interface{}) {
	l.logger.WithField("Fields", data).Panic(message)
}

// Trace 打印Trace日志
func (l *Logger) Trace(err error, data ...interface{}) {
	l.logger.WithField("Fields", data).Trace(err)
}

// Debug 全局方式调用Debug函数
func Debug(message string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Debug(message)
		return
	}
	logger.logger.WithField("Fields", data).Debug(message)
}

// Info 全局方式调用Info函数
func Info(message string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Info(message)
		return
	}
	logger.logger.WithField("Fields", data).Info(message)
}

// Warn 全局方式调用Warn函数
func Warn(message string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Warn(message)
		return
	}
	logger.logger.WithField("Fields", data).Warn(message)
}

// Error 全局方式调用Error函数
func Error(message string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Error(message)
		return
	}
	logger.logger.WithField("Fields", data).Error(message)
}

// Fatal 全局方式调用Fatal函数
func Fatal(message string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Fatal(message)
		return
	}
	logger.logger.WithField("Fields", data).Fatal(message)
}

// Panic 全局方式调用Panic函数
func Panic(message string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Panic(message)
		return
	}
	logger.logger.WithField("Fields", data).Panic(message)
}

// Trace 全局方式调用Trace函数
func Trace(err error, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Trace(err)
		return
	}
	logger.logger.WithField("Fields", data).Trace(err)
}
