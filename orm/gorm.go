package orm

// +---------------------------------------------------------------------
// | Description: gorm封装,包含日志记录功能
// +---------------------------------------------------------------------
// | Copyright (c) 2004-2020 护卫神(http://hws.com) All rights reserved.
// +---------------------------------------------------------------------
// | Author: Wjinlei <1976883731@qq.com>
// +---------------------------------------------------------------------
//
//                  ___====-_  _-====___
//             _--^^^#####/      \#####^^^--_
//          _-^##########/ (    ) \##########^-_
//         -############/  |\^^/|  \############-
//       _/############/   (@::@)   \############\_
//     /#############((     \  /     ))#############\
//     -###############\    (oo)    /###############-
//    -#################\  / VV \  /#################-
//   -###################\/      \/###################-
// _#/|##########/\######(   /\   )######/\##########|\#_
// |/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
// '  |/  V  V      V  \#\| |  | |/#/  V      V  V  \|  '
//    '   '  '      '   / | |  | | \   '      '  '   '
//                     (  | |  | |  )
//                    __\ | |  | | /__
//                   (vvv(VVV)(VVV)vvv)
//
//                  神龙护体
//                代码无bug!

import (
	"errors"
	"fmt"
	"log"
	"path/filepath"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// Option 选项
type Option struct {
	Driver     string // sqlite
	DataSource string
	LogMode    bool
	LogLevel   logger.LogLevel // 如果不传,默认为Silent(不记录日志)
}

// DB gorm.DB实例
type DB struct {
	Instance *gorm.DB
}

var (
	dbInstance *DB
	dbLogger   logger.Interface
)

// GetInstance 获取实例
func GetInstance(dbOption *Option) (*DB, error) {
	if dbInstance == nil {
		if dbOption == nil {
			return nil, errors.New("Option is nil")
		}
		if dbOption.LogMode == true {
			filePath, err := filepath.Abs("log/sql/sql.log")
			if err != nil {
				return nil, err
			}
			rotator, err := rotatelogs.New(
				fmt.Sprintf("%s-%s", filePath, "%Y%m%d%H%M"),
				rotatelogs.WithLinkName(filePath),
				rotatelogs.WithMaxAge(time.Duration(3*3600)*time.Second),     // 日志文件清理前的最长保存时间
				rotatelogs.WithRotationTime(time.Duration(3600)*time.Second), // 多久滚动一次
			)
			if err != nil {
				return nil, err
			}
			newLogger := logger.New(
				log.New(rotator, "\r\n", log.LstdFlags), // io writer
				logger.Config{
					SlowThreshold: time.Second,       // 慢 SQL 阈值
					LogLevel:      dbOption.LogLevel, // Log level
					Colorful:      false,             // 禁用彩色打印
				},
			)
			dbLogger = newLogger
		}
		if dbOption.Driver == "sqlite" {
			db, err := newSqlite(dbOption)
			if err != nil {
				return nil, err
			}
			return db, nil
		}
	}
	return dbInstance, nil
}

// newSqlite 产生sqlite实例
func newSqlite(dbOption *Option) (*DB, error) {
	db, err := gorm.Open(sqlite.Open(dbOption.DataSource), &gorm.Config{
		Logger: dbLogger,
	})
	if err != nil {
		return nil, err
	}
	dbInstance = &DB{Instance: db}
	return dbInstance, nil
}
