package myorm

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQL 产生MySQL实例
func NewMySQL(option *Option) (*DB, error) {
	if option == nil {
		return nil, fmt.Errorf("[ERROR]: Option is nil")
	}
	if option.LogMode == true {
		dbLogger, err := newLogger(option)
		if err != nil {
			return nil, err
		}
		db, err := gorm.Open(mysql.Open(option.DataSource),
			&gorm.Config{Logger: dbLogger})
		if err != nil {
			return nil, err
		}
		return &DB{Instance: db}, nil
	}
	db, err := gorm.Open(mysql.Open(option.DataSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return &DB{Instance: db}, nil
}
