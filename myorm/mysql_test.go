package myorm

import (
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestNewMySQL(t *testing.T) {
	defer handleRecover(t)
	db := newMysql()
	defer close(db)
}

func newMysql() *gorm.DB {
	mysql, err := NewMySQL("root:123@tcp(127.0.0.1:3306)/",
		NewLogger(&Option{
			LogPath:        "log/mysql.log",
			LogLevel:       logger.Info,
			RotationMaxAge: time.Duration(180) * time.Second, // 保留最近3分钟的日志
			RotationTime:   time.Duration(60) * time.Second,  // 每1分钟切割一次
		}))
	if err != nil {
		panic(err)
	}
	return mysql
}
