package orm

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type OptionStat struct {
	Driver     string
	DataSource string
}

type DB struct {
	OBJ *gorm.DB
}

var (
	db *DB
)

func GetDB() *DB {
	return db
}

func NewDB(option *OptionStat) (*DB, error) {
	newDB, err := gorm.Open(option.Driver, option.DataSource)
	if err != nil {
		return nil, err
	}
	db = &DB{
		OBJ: newDB,
	}
	return db, nil
}
