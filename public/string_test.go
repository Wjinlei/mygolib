package public

import "testing"

func TestChkDNFormat(t *testing.T) {
	dn := ChkDNFormat("www.test.com")
	if !dn {
		t.Error("check failed")
	}
}

func TestFilterSql(t *testing.T) {
	sql := FilterSql("()][].")
	if sql != "_" {
		t.Error(sql)
	}
}
