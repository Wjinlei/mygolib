package mypublic

import "testing"

func TestExists(t *testing.T) {
	if ok := Exists("./utils.go"); !ok {
		t.Error("TestExists failed")
	}
}
