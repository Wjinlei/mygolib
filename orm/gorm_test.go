package orm

import "testing"

func TestNewDB(t *testing.T) {
	db := GetInstance()
	if db == nil {
		var err error
		db, err = NewInstance(&OptionStat{
			DBDriver:   Sqlite,
			DataSource: "test.db",
			LogMode:    true,
		})
		if err != nil {
			t.Error(err)
			return
		}
	}
	if err := db.Instance.DB().Ping(); err != nil {
		t.Error(err)
	}
	db.Instance.AutoMigrate(&User{})
	db.Instance.Create(&User{Name: "user1"})
	db.Instance.Create(&User{Name: "user2"})
}
