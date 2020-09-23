package mycrypto

import "testing"

func TestAESEncrypt(t *testing.T) {
	encrypt := AESEncrypt("Hello World!", "123", "utf-8")
	if encrypt != "69dd1d8bda910a5d547de96054a70542" {
		t.Error(encrypt)
	}
}

func TestAESDecrypt(t *testing.T) {
	decrypt := AESDecrypt("69dd1d8bda910a5d547de96054a70542", "123", "utf-8")
	if decrypt != "Hello World!" {
		t.Error(decrypt)
	}
}
