package mypublic

import "testing"

func TestFilterSql(t *testing.T) {
	sql := FilterSQL("()][].")
	if sql != "_" {
		t.Error(sql)
	}
}

func TestCreateRandomString(t *testing.T) {
	strLen := 8
	str := CreateRandomString(strLen)
	if len(str) != strLen {
		t.Errorf("string len != %d", strLen)
	}
}
