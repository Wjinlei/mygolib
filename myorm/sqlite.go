package myorm

import (
	"fmt"

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
				Sqlite, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s?_journal_mode=wal", option.DataSource)), &gorm.Config{Logger: logger})
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
				Sqlite.Exec("PRAGMA foreign_keys = ON")
				conn, err := Sqlite.DB()
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
				conn.SetMaxOpenConns(1)
			} else {
				Sqlite, err = gorm.Open(sqlite.Open(fmt.Sprintf("%s?_journal_mode=wal", option.DataSource)), &gorm.Config{})
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
				Sqlite.Exec("PRAGMA foreign_keys = ON")
				conn, err := Sqlite.DB()
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
				conn.SetMaxOpenConns(1)
			}
		}
		mutex.Unlock()
	}
	return Sqlite, nil
}
