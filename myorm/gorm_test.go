package myorm

import (
	"testing"

	"gorm.io/gorm/logger"
)

func TestNewDB(t *testing.T) {
	db, err := GetInstance(&Option{
		Driver:     "sqlite",
		DataSource: "test.db",
		LogMode:    true,
		LogLevel:   logger.Info,
	})
	if err != nil {
		t.Error(err)
		return
	}
	db.Instance.AutoMigrate(&User{})
	db.Instance.Create(&User{Name: "user1"})
	db.Instance.Create(&User{Name: "user2"})
}
