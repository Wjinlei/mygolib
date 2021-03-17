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
	mysql, err := NewMySQL("root:123@tcp(127.0.0.1:3306)/", NewLogger("mysql.log", logger.Info))
	if err != nil {
		panic(err)
	}
	return mysql
}
