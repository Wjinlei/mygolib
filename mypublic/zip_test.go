package mypublic

import "testing"

func TestTarGZ(t *testing.T) {
	if err := TarGZ("file.go", "file.tar.gz"); err != nil {
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
