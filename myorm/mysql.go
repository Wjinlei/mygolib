package myorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// NewMySQL 产生MySQL实例
func NewMySQL(datasource string, logger logger.Interface) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(datasource), &gorm.Config{Logger: logger})
	if err != nil {
		return nil, err
	}
	return db, nil
}
