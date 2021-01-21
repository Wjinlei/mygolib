package myorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// NewMySQL 产生MySQL实例
func NewMySQL(option Option) (*gorm.DB, error) {
	if Mysql == nil {
		mutex.Lock()
		if Mysql == nil {
			if option.WriteLog {
				logger, err := newLogger(option.FilePath, option.Level)
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
				Mysql, err = gorm.Open(mysql.Open(option.DataSource), &gorm.Config{Logger: logger})
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
			} else {
				Mysql, err = gorm.Open(mysql.Open(option.DataSource), &gorm.Config{})
				if err != nil {
					mutex.Unlock()
					return nil, err
				}
			}
		}
		mutex.Unlock()
	}
	return Mysql, nil
}
