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
	PanicLevel LogLevel = iota
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
func (l *Logger) Debug(msg string, data Fields) {
	if l.logger.Level >= logrus.Level(DebugLevel) {
		if data == nil {
			data = Fields{}
		}
		l.logger.WithFields(logrus.Fields(data)).Debug(msg)
	}
}

// Info 打印Info日志
func (l *Logger) Info(msg string, data Fields) {
	if l.logger.Level >= logrus.Level(InfoLevel) {
		if data == nil {
			data = Fields{}
		}
		l.logger.WithFields(logrus.Fields(data)).Info(msg)
	}
}

// Warn 打印Warn日志
func (l *Logger) Warn(msg string, data Fields) {
	if l.logger.Level >= logrus.Level(WarnLevel) {
		if data == nil {
			data = Fields{}
		}
		l.logger.WithFields(logrus.Fields(data)).Warn(msg)
	}
}

// Error 打印Error日志
func (l *Logger) Error(msg string, data Fields) {
	if l.logger.Level >= logrus.Level(ErrorLevel) {
		if data == nil {
			data = Fields{}
		}
		l.logger.WithFields(logrus.Fields(data)).Error(msg)
	}
}

// Fatal 打印Fatal日志
func (l *Logger) Fatal(msg string, data Fields) {
	if l.logger.Level >= logrus.Level(FatalLevel) {
		if data == nil {
			data = Fields{}
		}
		l.logger.WithFields(logrus.Fields(data)).Fatal(msg)
	}
}

// Panic 打印Panic日志
func (l *Logger) Panic(msg string, data Fields) {
	if l.logger.Level >= logrus.Level(PanicLevel) {
		if data == nil {
			data = Fields{}
		}
		l.logger.WithFields(logrus.Fields(data)).Panic(msg)
	}
}
