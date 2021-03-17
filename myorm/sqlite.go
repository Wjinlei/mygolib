package myorm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewSqlite 产生sqlite实例
func NewSqlite(datasource string, logger logger.Interface) (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open(datasource), &gorm.Config{Logger: logger})
	if err != nil {
		return nil, err
	}
	db.Exec("PRAGMA foreign_keys = ON")
	return db, nil
}
