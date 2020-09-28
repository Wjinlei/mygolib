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
func newLogger(option *Option) (logger.Interface, error) {
	// 设置日志文件路径
	if option.LogPath == "" {
		option.LogPath = "log/sql/sql.log"
	}
	filePath, err := filepath.Abs(option.LogPath)
	if err != nil {
		return nil, err
	}
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
	// 产生新的日志器对象
	newLogger := logger.New(
		log.New(rotator, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second,
			LogLevel:      option.LogLevel,
			Colorful:      false,
		},
	)
	// 返回新日志器
	return newLogger, nil
}
