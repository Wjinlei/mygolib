package check

import "testing"

func TestChkDN(t *testing.T) {
	dn := ChkDN("www.test.com")
	if !dn {
		t.Error("check failed")
	}
}
