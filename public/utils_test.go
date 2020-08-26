package public

import "testing"

func TestPathExists(t *testing.T) {
	if ok := PathExists("./utils.go"); !ok {
		t.Error("TestExists failed")
	}
}
