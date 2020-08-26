package public

import "testing"

func TestZIP(t *testing.T) {
	if err := ZIP("./zip.go", "./test_nopassword.zip"); err != nil {
		t.Error(err)
	}
}

func TestZIPEncrypt(t *testing.T) {
	if err := ZIPEncrypt("./zip.go", "./test_password123.zip", "123"); err != nil {
		t.Error(err)
	}
}

func TestZIPDecrypt(t *testing.T) {
	if err := ZIPDecrypt("./test_password123.zip", "tmp", "123", "utf-8"); err != nil {
		t.Error(err)
	}
}
