package myencode

import "testing"

func TestGetEncoder(t *testing.T) {
	encoder := GetEncoder("utf-8")
	if encoder == nil {
		t.Error("GetEncoder failed...")
	}
}

func TestGetDecoder(t *testing.T) {
	decoder := GetDecoder("gbk")
	if decoder == nil {
		t.Error("GetDecoder failed...")
	}
}
