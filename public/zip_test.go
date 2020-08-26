package public

import "testing"

func TestZIP(t *testing.T) {
	if err := ZIP("./.sayHello.sh", "./test1.zip"); err != nil {
		t.Error(err)
	}
}

func TestZIPEncrypt(t *testing.T) {
	if err := ZIPEncrypt("./.sayHello.sh", "./test2.zip", "123"); err != nil {
		t.Error(err)
	}
}
