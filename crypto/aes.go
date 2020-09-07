package crypto

// +---------------------------------------------------------------------
// | Description: AES加解密
// +---------------------------------------------------------------------
// | Copyright (c) 2004-2020 护卫神(http://hws.com) All rights reserved.
// +---------------------------------------------------------------------
// | Author: Wjinlei <1976883731@qq.com>
// +---------------------------------------------------------------------
//
//                  ___====-_  _-====___
//             _--^^^#####/      \#####^^^--_
//          _-^##########/ (    ) \##########^-_
//         -############/  |\^^/|  \############-
//       _/############/   (@::@)   \############\_
//     /#############((     \  /     ))#############\
//     -###############\    (oo)    /###############-
//    -#################\  / VV \  /#################-
//   -###################\/      \/###################-
// _#/|##########/\######(   /\   )######/\##########|\#_
// |/ |#/\#/\#/\/  \#/\##\  |  |  /##/\#/  \/\#/\#/\#| \|
// '  |/  V  V      V  \#\| |  | |/#/  V      V  V  \|  '
//    '   '  '      '   / | |  | | \   '      '  '   '
//                     (  | |  | |  )
//                    __\ | |  | | /__
//                   (vvv(VVV)(VVV)vvv)
//
//                  神龙护体
//                代码无bug!

import (
	"bytes"
	"crypto/aes"
	"encoding/hex"
	"strings"

	"github.com/axgle/mahonia"
)

// AESEncrypt AES ECB 模式加密
func AESEncrypt(str string, key string, encoding string) (encrypt string) {
	defer func() {
		if err := recover(); err != nil {
			// 如果加密错误就返回原字符串
			encrypt = str
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

// AESDecrypt AES ECB 模式解密
func AESDecrypt(str string, key string, encoding string) (decrypt string) {
	defer func() {
		if err := recover(); err != nil {
			decrypt = "解密错误"
		}
	}()
	decoder := mahonia.NewDecoder(encoding)
	byteKey := make([]byte, 16)
	// key 只取16位
	copy(byteKey, decoder.ConvertString(key))
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
		string(bytes.Trim([]byte(decoder.ConvertString(string(decrypted[:trim]))),
			"\x00")))
	if decrypt == "" {
		decrypt = "解密错误"
	}
	return decrypt
}
