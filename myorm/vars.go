package myorm

import (
	"time"

	"gorm.io/gorm/logger"
)

// Option 选项定义
type Option struct {
	LogPath        string          // 日志路径
	LogLevel       logger.LogLevel // 日志级别
	RotationMaxAge time.Duration   // 日志文件按时间切割的保留时间
	RotationTime   time.Duration   // 日志文件按时间切割
}
