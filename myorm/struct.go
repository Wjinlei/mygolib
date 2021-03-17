package myorm

import "gorm.io/gorm/logger"

// Option 选项定义
type Option struct {
	LogPath  string          // 日志路径
	LogLevel logger.LogLevel // 日志级别
}
