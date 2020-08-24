package file

import "testing"

func TestWriteFile(t *testing.T) {
	if err := WriteFile("./testfile1.txt", "TestWriteFile"); err != nil {
		t.Error(err)
	}
}

func TestReadFile(t *testing.T) {
	_, err := ReadFile("./testfile1.txt")
	if err != nil {
		t.Error(err)
	}
}

func TestFileExists(t *testing.T) {
	if ok := FileExists("./testfile1.txt"); !ok {
		t.Error("TestFileExists failed")
	}
}

func TestFileNotExists(t *testing.T) {
	if ok := FileNotExists("./testfile1.txt"); ok {
		t.Error("TestFileNotExists failed")
	}
}

func TestMoveFile(t *testing.T) {
	if err := MoveFile("./testfile1.txt", "./testfile1.txt.bak"); err != nil {
		t.Error(err)
	}
}
