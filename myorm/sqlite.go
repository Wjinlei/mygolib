package myorm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewSqlite 产生sqlite实例
func NewSqlite(option Option) (*gorm.DB, error) {
	if option.WriteLog {
		db, err := gorm.Open(sqlite.Open(option.DataSource), &gorm.Config{Logger: newLogger(option.FilePath, option.Level)})
		if err != nil {
			return nil, err
		}
		db.Exec("PRAGMA foreign_keys = ON")
		return db, nil
	} else {
		db, err := gorm.Open(sqlite.Open(option.DataSource), &gorm.Config{})
		if err != nil {
			return nil, err
		}
		db.Exec("PRAGMA foreign_keys = ON")
		return db, nil
	}
}
