package orm

import (
	"errors"
	"time"

	"github.com/Wjinlei/mygolib/logger"
	"github.com/jinzhu/gorm"
)

// OptionStat 选项
type OptionStat struct {
	DBDriver   Driver
	DataSource string
	LogMode    bool
	Logger     *logger.Logger
}

// Driver 驱动类型
type Driver string

const (
	// Sqlite sqlite 驱动名定义
	Sqlite Driver = "sqlite3"
)

// DBStat gorm.DB实例
type DBStat struct {
	Instance *gorm.DB
}

var (
	globalDB *DBStat
	dbOption *OptionStat
	mylogger *logger.Logger
)

// CustomLogger 自定义日志器
type CustomLogger struct{}

// Print 打印日志
func (*CustomLogger) Print(v ...interface{}) {
	switch v[0] {
	case "sql":
		mylogger.Debug(logger.Fields{
			"module":   "gorm",
			"type":     "sql",
			"file":     v[1], // sql代码行号
			"duration": v[2], // sql执行时间戳
			"sql":      v[3], // sql语句
			"values":   v[4], // sql数据
			"rows":     v[5], // sql影响的行数
		}, "sql日志")
	case "log":
		mylogger.Debug(logger.Fields{
			"module":   "gorm",
			"type":     "log",
			"duration": v[2],
		}, "普通日志")
	}
}

// GetInstance 获取实例
func GetInstance() *DBStat {
	return globalDB
}

// NewInstance 产生新的实例
func NewInstance(dbOption *OptionStat) (*DBStat, error) {
	if globalDB == nil {
		if dbOption == nil {
			return nil, errors.New("Option is nil")
		}
		if dbOption.LogMode == true {
			if dbOption.Logger == nil {
				newLogger, err := logger.New(&logger.Option{
					LogPath:      "./log/db/db.log",
					LogLevel:     logger.DebugLevel,
					LogType:      "json",
					MaxAge:       time.Duration(3*3600) * time.Second,
					RotationTime: time.Duration(3600) * time.Second,
				})
				if err != nil {
					return nil, err
				}
				mylogger = newLogger
			} else {
				mylogger = dbOption.Logger
			}
		}
		if dbOption.DBDriver == Sqlite {
			sqliteDB, err := newSqlite(dbOption)
			if err != nil {
				return nil, err
			}
			return sqliteDB, nil
		}
	}
	return globalDB, nil
}

// 产生sqlite实例
func newSqlite(dbOption *OptionStat) (*DBStat, error) {
	sqliteDB, err := gorm.Open("sqlite3", dbOption.DataSource)
	if err != nil {
		return nil, err
	}
	sqliteDB.SetLogger(&CustomLogger{})
	sqliteDB.LogMode(dbOption.LogMode)
	globalDB = &DBStat{Instance: sqliteDB}
	return globalDB, nil
}
