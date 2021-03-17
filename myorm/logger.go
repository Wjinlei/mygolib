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
func NewLogger(option *Option) logger.Interface {
	return logger.New(log.New(newRotator(option.LogPath), "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      option.LogLevel,
		Colorful:      false,
	})
}

// newRotator 产生新的日志切割器
func newRotator(logpath string) *rotatelogs.RotateLogs {
	absPath := abs(logpath)
	// 日志切割器,日志文件最长保留3小时,每小时滚动一次
	rotator, _ := rotatelogs.New(
		fmt.Sprintf("%s-%s", absPath, "%Y%m%d%H%M"),
		rotatelogs.WithLinkName(absPath),
		rotatelogs.WithMaxAge(time.Duration(3*3600)*time.Second),     // 日志文件清理前的最长保存时间,保存24小时的
		rotatelogs.WithRotationTime(time.Duration(3600)*time.Second), // 按时间滚动,1小时滚动一次
	)
	return rotator
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
