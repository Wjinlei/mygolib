package orm

import (
	"errors"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
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
	globalDB    *DBStat
	globalError error
	dbOption    *OptionStat
)

func GetInstance() *DBStat {
	return globalDB
}

func NewInstance(dbOption *OptionStat) (*DBStat, error) {
	if globalDB == nil {
		if dbOption == nil {
			return nil, errors.New("Option is nil")
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

func newSqlite(dbOption *OptionStat) (*DBStat, error) {
	sqliteDB, err := gorm.Open("sqlite3", dbOption.DataSource)
	if err != nil {
		return nil, err
	}
	sqliteDB.LogMode(dbOption.LogMode)
	globalDB = &DBStat{Instance: sqliteDB}
	return globalDB, nil
}
