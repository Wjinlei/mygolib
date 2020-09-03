package logger

import (
	"context"
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
	PrettyPrint  bool          // 是否美化输出
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
	TraceLevel
)

// Interface 和Gorm 2.0集成
type Interface interface {
	LogMode(LogLevel) Interface
	Debug(context.Context, string, ...interface{})
	Info(context.Context, string, ...interface{})
	Warn(context.Context, string, ...interface{})
	Error(context.Context, string, ...interface{})
	Fatal(context.Context, string, ...interface{})
	Panic(context.Context, string, ...interface{})
	Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error)
}

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
func New(option *Option) (Interface, error) {
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

// LogMode gorm 2.0 会用到这个函数来设置日志器的日志级别
func (l *Logger) LogMode(level LogLevel) Interface {
	newlogger := *l
	newlogger.logger.SetLevel(logrus.Level(level))
	return &newlogger
}

// Debug 打印Debug日志
func (l *Logger) Debug(ctx context.Context, msg string, data ...interface{}) {
	if l.logger.Level >= logrus.Level(DebugLevel) {
		l.logger.WithField("Fields", data).Debug(msg)
	}
}

// Info 打印Info日志
func (l *Logger) Info(ctx context.Context, msg string, data ...interface{}) {
	if l.logger.Level >= logrus.Level(InfoLevel) {
		l.logger.WithField("Fields", data).Info(msg)
	}
}

// Warn 打印Warn日志
func (l *Logger) Warn(ctx context.Context, msg string, data ...interface{}) {
	if l.logger.Level >= logrus.Level(WarnLevel) {
		l.logger.WithField("Fields", data).Warn(msg)
	}
}

// Error 打印Error日志
func (l *Logger) Error(ctx context.Context, msg string, data ...interface{}) {
	if l.logger.Level >= logrus.Level(ErrorLevel) {
		l.logger.WithField("Fields", data).Error(msg)
	}
}

// Fatal 打印Fatal日志
func (l *Logger) Fatal(ctx context.Context, msg string, data ...interface{}) {
	if l.logger.Level >= logrus.Level(FatalLevel) {
		l.logger.WithField("Fields", data).Fatal(msg)
	}
}

// Panic 打印Panic日志
func (l *Logger) Panic(ctx context.Context, msg string, data ...interface{}) {
	if l.logger.Level >= logrus.Level(PanicLevel) {
		l.logger.WithField("Fields", data).Panic(msg)
	}
}

// Trace 打印Trace日志
func (l *Logger) Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	l.logger.WithFields(logrus.Fields{"sql": sql, "rows": rows}).Trace(err)
}

// Debug 全局方式调用Debug函数
func Debug(ctx context.Context, msg string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Debug(msg)
		return
	}
	logger.logger.WithField("Fields", data).Debug(msg)
}

// Info 全局方式调用Info函数
func Info(ctx context.Context, msg string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Info(msg)
		return
	}
	logger.logger.WithField("Fields", data).Info(msg)
}

// Warn 全局方式调用Warn函数
func Warn(ctx context.Context, msg string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Warn(msg)
		return
	}
	logger.logger.WithField("Fields", data).Warn(msg)
}

// Error 全局方式调用Error函数
func Error(ctx context.Context, msg string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Error(msg)
		return
	}
	logger.logger.WithField("Fields", data).Error(msg)
}

// Fatal 全局方式调用Fatal函数
func Fatal(ctx context.Context, msg string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Fatal(msg)
		return
	}
	logger.logger.WithField("Fields", data).Fatal(msg)
}

// Panic 全局方式调用Panic函数
func Panic(ctx context.Context, msg string, data ...interface{}) {
	if logger.logger == nil {
		logrus.WithField("Fields", data).Panic(msg)
		return
	}
	logger.logger.WithField("Fields", data).Panic(msg)
}

// Trace 全局方式调用Trace函数
func Trace(ctx context.Context, begin time.Time, fc func() (string, int64), err error) {
	sql, rows := fc()
	if logger.logger == nil {
		logrus.WithFields(logrus.Fields{"sql": sql, "rows": rows}).Trace(err)
	}
	logger.logger.WithFields(logrus.Fields{"sql": sql, "rows": rows}).Trace(err)
}
