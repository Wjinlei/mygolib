package mycrypto

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func TestAESEncrypt(t *testing.T) {
	err := SetAesKey("123", 16) // 可以指定密钥的长度, 16, 24, 32, 不足为会自动填充0
	if err != nil {
		t.Error(err)
	}
	//ciphertext, err := AesECBEncrypt([]byte("Hello World!"), "ZeroPadding") // 还可以指定零填充
	ciphertext, err := AesECBEncrypt([]byte("Hello World!")) // 默认采用PKCS5算法来Padding填充
	if err != nil {
		t.Error(err)
	}
	encrypt := hex.EncodeToString(ciphertext) // 如果要加密结果转为十六禁止,则需要编码
	fmt.Println(encrypt)
}

func TestAesECBDecrypt(t *testing.T) {
	err := SetAesKey("123", 16) // 可以指定密钥的长度, 16, 24, 32, 不足为会自动填充0
	if err != nil {
		t.Error(err)
	}
	decode, err := hex.DecodeString("69dd1d8bda910a5d547de96054a70542") // 如果密文是16进制编码,则需要解码
	if err != nil {
		t.Error(err)
	}
	//ciphertext, err := AesECBDecrypt(decode, "ZeroUnPadding") // 还可以指定零反填充
	ciphertext, err := AesECBDecrypt(decode) // 默认采用PKCS5算法来UnPadding填充
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(ciphertext))
}
