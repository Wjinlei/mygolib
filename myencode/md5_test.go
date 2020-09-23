package myencode

import (
	"testing"
)

func TestStringToMd5(t *testing.T) {
	strMd5 := StringToMd5("Hello World!")
	if strMd5 != "ed076287532e86365e841e92bfc50d8c" {
		t.Error("md5 does not match")
	}
}

func TestFileToMd5(t *testing.T) {
	fileMd5, err := FileToMd5("./md5.go")
	if err != nil {
		t.Error(err)
	}
	if fileMd5 != "1035281b2855e5d9cc428219890adf3d" {
		t.Error("md5 does not match")
	}
}
