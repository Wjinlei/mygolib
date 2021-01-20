package myorm

import (
	"testing"

	"gorm.io/gorm/logger"
)

func TestNewMySQL(t *testing.T) {
	_, err := NewMySQL(Option{
		DataSource: "root:123@tcp(127.0.0.1:3306)/",
		LogMode:    true,
		LogLevel:   logger.Info,
	})
	if err != nil {
		t.Error(err)
	}
}
