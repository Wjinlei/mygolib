package orm

import "testing"

func TestNewDB(t *testing.T) {
	db, err := NewDB(&OptionStat{
		Driver:     "sqlite3",
		DataSource: "test.db",
	})
	if err != nil {
		t.Error(err)
	}
	if err := db.OBJ.DB().Ping(); err != nil {
		t.Error(err)
	}
}
