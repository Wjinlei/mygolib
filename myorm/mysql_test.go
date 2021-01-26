package myorm

import (
	"testing"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestNewMySQL(t *testing.T) {
	defer handleRecover(t)
	db := newMysql()
	defer close(db)
}

func newMysql() *gorm.DB {
	mysql, err := NewMySQL(Option{
		DataSource: "root:123@tcp(127.0.0.1:3306)/",
		WriteLog:   true,
		Level:      logger.Info,
	})
	if err != nil {
		panic(err)
	}
	return mysql
}
