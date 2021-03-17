package myorm

import (
	"testing"

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
			LogPath:  "log/sqlite.log",
			LogLevel: logger.Info,
		}))
	if err != nil {
		panic(err)
	}
	return db
}
