package mystring

import "testing"

func TestFilterSql(t *testing.T) {
	sql := FilterSql("()][].")
	if sql != "_" {
		t.Error(sql)
	}
}
