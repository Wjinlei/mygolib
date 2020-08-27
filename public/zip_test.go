package public

import "testing"

func TestZIP(t *testing.T) {
	if err := ZIP("./zip.go", "./test_nopassword.zip"); err != nil {
		t.Error(err)
	}
}

func TestZIPDecrypt(t *testing.T) {
	if err := ZIPDecrypt("./test_nopassword.zip", "/tmp", "", "utf-8"); err != nil {
		t.Error(err)
	}
}
