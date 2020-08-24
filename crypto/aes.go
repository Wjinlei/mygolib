package crypto

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"github.com/axgle/mahonia"
	"strings"
)

// AES ECB 模式加密
func AESEncrypt(str string, key string, encoding string) (encrypt string) {
	defer func() {
		if err := recover(); err != nil {
			encrypt = ""
		}
	}()
	encoder := mahonia.NewEncoder(encoding)
	byteStr := []byte(encoder.ConvertString(str))
	byteKey := make([]byte, 16)
	// key 只取16位
	copy(byteKey, encoder.ConvertString(key))

	for i := 16; i < len(byteKey); {
		for j := 0; j < 16 && i < len(byteKey); j, i = j+1, i+1 {
			byteKey[j] ^= byteKey[i]
		}
	}
	cipher, _ := aes.NewCipher(byteKey)
	length := (len(byteStr) + aes.BlockSize) / aes.BlockSize
	plain := make([]byte, length*aes.BlockSize)
	copy(plain, byteStr)
	pad := byte(len(plain) - len(byteStr))
	for i := len(byteStr); i < len(plain); i++ {
		plain[i] = pad
	}
	encrypted := make([]byte, len(plain))
	for bs, be := 0,
		cipher.BlockSize(); bs <= len(byteStr); bs,
		be = bs+cipher.BlockSize(), be+cipher.BlockSize() {
		cipher.Encrypt(encrypted[bs:be], plain[bs:be])
	}
	encrypt = strings.TrimSpace(hex.EncodeToString(encrypted))
	return encrypt
}

// AES ECB 模式解密
func AESDecrypt(str string, key string, encoding string) (decrypt string) {
	defer func() {
		if err := recover(); err != nil {
			decrypt = ""
		}
	}()
	encoder := mahonia.NewDecoder(encoding)
	byteKey := make([]byte, 16)
	// key 只取16位
	copy(byteKey, encoder.ConvertString(key))
	for i := 16; i < len(byteKey); {
		for j := 0; j < 16 && i < len(byteKey); j, i = j+1, i+1 {
			byteKey[j] ^= byteKey[i]
		}
	}
	decodeString, _ := hex.DecodeString(str)
	cipher, _ := aes.NewCipher(byteKey)
	decrypted := make([]byte, len(decodeString))
	for bs, be := 0,
		cipher.BlockSize(); bs < len(decodeString); bs, be = bs+cipher.BlockSize(),
		be+cipher.BlockSize() {
		cipher.Decrypt(decrypted[bs:be], decodeString[bs:be])
	}
	trim := 0
	if len(decrypted) > 0 {
		trim = len(decrypted) - int(decrypted[len(decrypted)-1])
	}
	// 去除两边多余的0,和多余的空格
	decrypt = strings.TrimSpace(
		string(bytes.Trim([]byte(encoder.ConvertString(string(decrypted[:trim]))),
			"\x00")))
	return decrypt
}