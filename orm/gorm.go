package orm

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"mylib/logger"
)

type OptionStat struct {
	DBDriver   Driver
	DataSource string
	LogMode    bool
}

type Driver string

const (
	Sqlite Driver = "sqlite3"
)

type DBStat struct {
	Instance *gorm.DB
}

var (
	globalDB *DBStat
	dbOption *OptionStat
	mylogger *logger.Logger
)

// 自定义日志器
type CustomLogger struct{}

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

func GetInstance() *DBStat {
	return globalDB
}

// 产生新的实例
func NewInstance(dbOption *OptionStat) (*DBStat, error) {
	if globalDB == nil {
		if dbOption == nil {
			return nil, errors.New("Option is nil")
		}
		if dbOption.LogMode == true {
			newLogger, err := logger.New(&logger.Option{
				LogPath:      "./log/db/db.log",
				LogLevel:     logger.DebugLevel,
				LogType:      "json",
				MaxAge:       time.Duration(3*24*3600) * time.Second,
				RotationTime: time.Duration(3600) * time.Second,
			})
			if err != nil {
				return nil, err
			}
			mylogger = newLogger
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
