package random

import "testing"

func TestCreateRandomString(t *testing.T) {
	strLen := 8
	str := CreateRandomString(strLen)
	if len(str) != strLen {
		t.Errorf("string len != %d", strLen)
	}
}
