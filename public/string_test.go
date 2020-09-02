package public

import "testing"

func TestFilterSql(t *testing.T) {
	sql := FilterSQL("()][].")
	if sql != "_" {
		t.Error(sql)
	}
}
