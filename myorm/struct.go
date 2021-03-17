package myorm

import "gorm.io/gorm/logger"

// Option 选项定义
type Option struct {
	LogPath       string          // 日志路径
	LogLevel      logger.LogLevel // 日志级别
	RotationCount uint            // 日志文件保留个数
	RotationSize  int64           // 日志文件滚动大小,单位是KB
}
