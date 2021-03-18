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
	rotator, _ := rotatelogs.New(
		fmt.Sprintf("%s-%s", abs(option.LogPath), "%Y%m%d%H%M"),
		rotatelogs.WithLinkName(abs(option.LogPath)),
		rotatelogs.WithMaxAge(option.RotationMaxAge),
		rotatelogs.WithRotationTime(option.RotationTime),
	)
	return logger.New(log.New(rotator, "\r\n", log.LstdFlags), logger.Config{
		SlowThreshold: time.Second,
		LogLevel:      option.LogLevel,
		Colorful:      false,
	})
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
