package myorm

import (
	"testing"

	"gorm.io/gorm/logger"
)

func TestNewSqlite(t *testing.T) {
	db, err := NewSqlite(Option{
		DataSource: "test.db",
		WriteLog:   true,
		Level:      logger.Info,
	})
	if err != nil {
		t.Fatal(err)
	}
	conn, err := db.DB()
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()
	db.AutoMigrate(&TestModel{})
	db.Create(&TestModel{Name: "zhangSan"})
	db.Create(&TestModel{Name: "LiSi"})
	db.Create(&TestModel{Name: "WangWu"})
}
