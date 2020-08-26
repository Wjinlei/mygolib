package public

import "testing"

func TestMakeDir(t *testing.T) {
	if err := MakeDir("test1/test1.txt"); err != nil {
		t.Error(err)
	}
}

func TestMakeDirAll(t *testing.T) {
	if err := MakeDirAll("test2/test2.1/test2.txt"); err != nil {
		t.Error(err)
	}
}
