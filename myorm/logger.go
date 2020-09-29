package myorm

import (
	"fmt"
	"log"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gorm.io/gorm/logger"
)

// newLogger 产生新的日志器对象
func newLogger(loglevel logger.LogLevel, rotator *rotatelogs.RotateLogs) (logger.Interface, error) {
	// 产生新的日志器对象
	newLogger := logger.New(
		log.New(rotator, "\r\n", log.LstdFlags), logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      loglevel,
			Colorful:      false,
		},
	)
	// 返回新日志器
	return newLogger, nil
}

// newRotator 产生新的日志切割器
func newRotator(logpath string) (*rotatelogs.RotateLogs, error) {
	logpath = handleLogPath(logpath)
	// 日志切割器,日志文件最长保留3小时,每小时滚动一次
	rotator, err := rotatelogs.New(
		fmt.Sprintf("%s-%s", logpath, "%Y%m%d%H%M"),
		rotatelogs.WithLinkName(logpath),
		rotatelogs.WithMaxAge(time.Duration(3*3600)*time.Second),     // 日志文件清理前的最长保存时间
		rotatelogs.WithRotationTime(time.Duration(3600)*time.Second), // 多久滚动一次
	)
	if err != nil {
		return nil, err
	}
	return rotator, nil
}

// handleLogPath 处理日志路径
func handleLogPath(path string) string {
	if path == "" {
		path = "log/sql/sql.log"
	}
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "/tmp/hwsmaster/sql/sql.log"
	}
	return absPath
}
