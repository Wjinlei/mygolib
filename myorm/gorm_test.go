package myorm

import (
	"testing"

	"gorm.io/gorm/logger"
)

func TestGetInstance(t *testing.T) {
	db, err := GetInstance(&Option{
		Driver:     "sqlite",
		DataSource: "test.db",
		LogMode:    true,
		LogPath:    "log/sql/sqlite.log",
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

/*
func TestGetInstance(t *testing.T) {
	_, err := GetInstance(&Option{
		Driver:     "mysql",
		DataSource: "root:123@tcp(192.168.2.126:3306)/",
		LogMode:    true,
		LogPath:    "log/sql/mysql.log",
		LogLevel:   logger.Info,
	})
	if err != nil {
		t.Error(err)
		return
	}
}
*/
