package myencode

import "testing"

func TestGetPunyCode(t *testing.T) {
	punycode := GetPunyCode("中文.中国")
	if punycode != "xn--fiq228c.xn--fiqs8s" {
		t.Error("Get PunyCode failed")
	}
}
