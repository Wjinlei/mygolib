package myorm

import (
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gorm.io/gorm/logger"
)

// newLogger 产生新的日志器对象
func newLogger(logpath string, loglevel logger.LogLevel) (logger.Interface, error) {
	rotator, err := newRotator(logpath)
	if err != nil {
		return nil, err
	}
	mylogger := logger.New(log.New(rotator, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      loglevel,
			Colorful:      false,
		},
	)
	return mylogger, nil
}

// newRotator 产生新的日志切割器
func newRotator(logpath string) (*rotatelogs.RotateLogs, error) {
	filePath := abs(logpath)
	// 日志切割器,日志文件最长保留3小时,每小时滚动一次
	rotator, err := rotatelogs.New(
		fmt.Sprintf("%s-%s", filePath, "%Y%m%d%H%M"),
		rotatelogs.WithLinkName(filePath),
		rotatelogs.WithMaxAge(time.Duration(3*3600)*time.Second),     // 日志文件清理前的最长保存时间
		rotatelogs.WithRotationTime(time.Duration(3600)*time.Second), // 多久滚动一次
	)
	if err != nil {
		return nil, err
	}
	return rotator, nil
}

func abs(logpath string) string {
	newpath := strings.TrimSpace(logpath)
	if newpath == "" {
		return "log/sql/sql.log"
	}
	abspath, err := filepath.Abs(newpath)
	if err != nil {
		return "log/sql/sql.log"
	}
	return abspath
}
