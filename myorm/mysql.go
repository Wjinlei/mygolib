package myorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQL 产生MySQL实例
func NewMySQL(option Option) (*gorm.DB, error) {
	if option.WriteLog {
		db, err := gorm.Open(mysql.Open(option.DataSource), &gorm.Config{Logger: newLogger(option.FilePath, option.Level)})
		if err != nil {
			return nil, err
		}
		return db, nil
	} else {
		db, err := gorm.Open(mysql.Open(option.DataSource), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		return db, nil
	}
}
