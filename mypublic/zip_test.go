package mypublic

import "testing"

func TestTGZ(t *testing.T) {
	if err := TGZ("file.go", "file.tar.gz"); err != nil {
		t.Error(err)
	}
}

func TestZIP(t *testing.T) {
	if err := ZIP("file.go", "file.zip", "GBK"); err != nil {
		t.Error(err)
	}
}

func TestZIPDecrypt(t *testing.T) {
	if err := ZIPDecrypt("file.zip", "/tmp", "", "GBK"); err != nil {
		t.Error(err)
	}
}
