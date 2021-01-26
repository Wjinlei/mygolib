package myorm

import (
	"testing"

	"gorm.io/gorm/logger"
)

func TestNewMySQL(t *testing.T) {
	mysql, err := NewMySQL(Option{
		DataSource: "root:123@tcp(127.0.0.1:3306)/",
		WriteLog:   true,
		Level:      logger.Info,
	})
	if err != nil {
		t.Fatal(err)
	}
	conn, err := mysql.DB()
	if err != nil {
		t.Error(err)
	}
	conn.Close()
}
