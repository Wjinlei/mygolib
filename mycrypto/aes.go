package mycrypto

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
func AESEncrypt(decrypt string, key string, encoding string) (encrypt string) {
	blockSize := 16
	enCoder := mahonia.NewEncoder(encoding)
	defer func() {
		if err := recover(); err != nil {
			encrypt = decrypt
		}
	}()
	keyByte := make([]byte, 16)
	copy(keyByte, enCoder.ConvertString(key))
	newCipher, err := aes.NewCipher(keyByte)
	if err != nil {
		return decrypt
	}
	decryptByte := []byte(enCoder.ConvertString(decrypt))
	decryptLength := len(decryptByte)
	plain := make([]byte, (len(decryptByte)+blockSize)/blockSize*blockSize)
	plainLength := len(plain)
	copy(plain, decryptByte)
	padding := byte(plainLength - decryptLength)
	for i := decryptLength; i < plainLength; i++ {
		plain[i] = padding
	}
	encryptByte := make([]byte, plainLength)
	for bs, be := 0, blockSize; bs <= decryptLength; bs, be = bs+blockSize, be+blockSize {
		newCipher.Encrypt(encryptByte[bs:be], plain[bs:be])
	}
	encrypt = strings.TrimSpace(hex.EncodeToString(encryptByte))
	return encrypt // 如果加密失败,则返回原字符串
}

// AESDecrypt AES ECB 模式解密
func AESDecrypt(encrypt string, key string, encoding string) (decrypt string) {
	blockSize := 16
	deCoder := mahonia.NewDecoder(encoding)
	defer func() {
		if err := recover(); err != nil {
			decrypt = "解密错误"
		}
	}()
	keyByte := make([]byte, 16)
	copy(keyByte, deCoder.ConvertString(key))
	newCipher, err := aes.NewCipher(keyByte)
	if err != nil {
		return "解密错误"
	}
	encryptByte, _ := hex.DecodeString(encrypt)
	encryptByteLength := len(encryptByte)
	decryptByte := make([]byte, encryptByteLength)
	for bs, be := 0, blockSize; bs < encryptByteLength; bs, be = bs+blockSize, be+blockSize {
		newCipher.Decrypt(decryptByte[bs:be], encryptByte[bs:be])
	}
	decrypt = strings.TrimSpace(string(bytes.Trim([]byte(deCoder.ConvertString(string(decryptByte))), "\x00")))
	return decrypt // 返回的结果,不能用 == 判断是否相等,因为[]byte可能并不相同
}
