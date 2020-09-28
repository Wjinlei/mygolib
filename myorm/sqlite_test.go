package myorm

import (
	"testing"

	"gorm.io/gorm/logger"
)

func TestNewSqlite(t *testing.T) {
	db, err := NewSqlite(&Option{
		DataSource: "hws.db",
		LogMode:    true,
		LogLevel:   logger.Info,
	})
	if err != nil {
		t.Error(err)
	}
	db.Instance.AutoMigrate(&TestModel{})
	db.Instance.Create(&TestModel{Name: "zhangSan"})
	db.Instance.Create(&TestModel{Name: "LiSi"})
	db.Instance.Create(&TestModel{Name: "WangWu"})
}
