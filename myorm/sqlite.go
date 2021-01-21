package myorm

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// NewSqlite 产生sqlite实例
func NewSqlite(option Option) (*gorm.DB, error) {
	if Sqlite == nil {
		mutex.Lock()
		if Sqlite == nil {
			if option.WriteLog {
				logger, err := newLogger(option.FilePath, option.Level)
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
				Sqlite, err = gorm.Open(sqlite.Open(option.DataSource), &gorm.Config{Logger: logger})
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
				Sqlite.Exec("PRAGMA foreign_keys = ON")
			} else {
				Sqlite, err = gorm.Open(sqlite.Open(option.DataSource), &gorm.Config{})
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
				Sqlite.Exec("PRAGMA foreign_keys = ON")
			}
		}
		mutex.Unlock()
	}
	return Sqlite, nil
}
