package myorm

import (
	"testing"

	"gorm.io/gorm/logger"
)

func TestNewSqlite(t *testing.T) {
	db, err := NewSqlite(Option{
		DataSource: "test.db",
		LogMode:    true,
		LogLevel:   logger.Info,
	})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&TestModel{})
	db.Create(&TestModel{Name: "zhangSan"})
	db.Create(&TestModel{Name: "LiSi"})
	db.Create(&TestModel{Name: "WangWu"})
}
