package myorm

import (
	"testing"
	"time"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestNewSqlite(t *testing.T) {
	defer handleRecover(t)
	db := newSqlite()
	defer close(db)
	db.AutoMigrate(&TestModel{})
	db.Create(&TestModel{Name: "zhangSan"})
	db.Create(&TestModel{Name: "LiSi"})
	db.Create(&TestModel{Name: "WangWu"})
}

func newSqlite() *gorm.DB {
	db, err := NewSqlite("sqlite.db",
		NewLogger(&Option{
			LogPath:        "log/sqlite.log",
			LogLevel:       logger.Info,
			RotationMaxAge: time.Duration(180) * time.Second, // 保留最近3分钟的日志
			RotationTime:   time.Duration(60) * time.Second,  // 每1分钟切割一次
		}))
	if err != nil {
		panic(err)
	}
	return db
}
