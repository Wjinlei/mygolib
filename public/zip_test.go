package public

import "testing"

func TestZIP(t *testing.T) {
	if err := ZIP("./dir.go", "./test_nopassword.zip", "GBK"); err != nil {
		t.Error(err)
	}
}

func TestZIPDecrypt(t *testing.T) {
	if err := ZIPDecrypt("./test_nopassword.zip", "/tmp", "", "GBK"); err != nil {
		t.Error(err)
	}
}
