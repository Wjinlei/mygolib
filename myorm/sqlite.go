package myorm

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewSqlite 产生sqlite实例
func NewSqlite(option *Option) (*DB, error) {
	if option == nil {
		return nil, fmt.Errorf("[ERROR]: Option is nil")
	}
	if option.LogMode == true {
		dbLogger, err := newLogger(option)
		if err != nil {
			return nil, err
		}
		db, err := gorm.Open(sqlite.Open(option.DataSource),
			&gorm.Config{Logger: dbLogger})
		if err != nil {
			return nil, err
		}
		if err := openForeginKey(db); err != nil {
			return nil, err
		}
		return &DB{Instance: db}, nil
	}
	db, err := gorm.Open(sqlite.Open(option.DataSource), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	if err := openForeginKey(db); err != nil {
		return nil, err
	}
	return &DB{Instance: db}, nil
}

// openForeginKey 打开Sqlite3外键支持
func openForeginKey(db *gorm.DB) error {
	if err := db.Exec("PRAGMA foreign_keys = ON").Error; err != nil {
		return err
	}
	return nil
}
