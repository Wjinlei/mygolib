package public

import "testing"

func TestExists(t *testing.T) {
	if ok := Exists("./utils.go"); !ok {
		t.Error("TestExists failed")
	}
}

func TestNotExists(t *testing.T) {
	if ok := NotExists("./utils.go"); ok {
		t.Error("TestNotExists failed")
	}
}
