package file

import "testing"

func TestExists(t *testing.T) {
	if ok := Exists("./dir.go"); !ok {
		t.Error("TestExists failed")
	}
}

func TestNotExists(t *testing.T) {
	if ok := NotExists("./dir.go"); ok {
		t.Error("TestNotExists failed")
	}
}
