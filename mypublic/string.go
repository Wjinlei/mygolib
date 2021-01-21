package mypublic

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
	"strings"
)

// 字符串转整型,错误返回0
func Atoi(s string) int {
	sInt, err := strconv.Atoi(s)
	if err != nil {
		return 0
	}
	return sInt
}

func FilterStr(str string) string {
	str = strings.ReplaceAll(str, "'", "")
	str = strings.ReplaceAll(str, "\"", "")
	str = strings.ReplaceAll(str, "(", "")
	str = strings.ReplaceAll(str, ")", "")
	//str = strings.ReplaceAll(str, "*", "")
	str = strings.ReplaceAll(str, ".", "_")
	str = strings.ReplaceAll(str, "[", "")
	str = strings.ReplaceAll(str, "]", "")
	str = strings.ReplaceAll(str, " ", "")
	str = strings.ReplaceAll(str, "=", "")
	str = strings.ReplaceAll(str, ">", "")
	str = strings.ReplaceAll(str, "<", "")
	str = strings.ReplaceAll(str, "%", "")
	str = strings.ReplaceAll(str, "\t", "")
	return str
}

// Int64ToKBMBGB 将int64类型表示的字节转换为带上单位的String
func Int64ToKBMBGB(value int64) string {
	var retStr string
	if value > 1073741824 {
		retStr = fmt.Sprintf("%.2f", float64(value)/float64(1024)/float64(1024)/float64(1024)) + "(GB)"
	} else if value > 1048576 {
		retStr = fmt.Sprintf("%.2f", float64(value)/float64(1024)/float64(1024)) + "(MB)"
	} else if value > 1024 {
		retStr = fmt.Sprintf("%.2f", float64(value)/float64(1024)) + "(KB)"
	} else {
		retStr = fmt.Sprintf("%.2f", float64(value)) + "(B)"
	}
	return retStr
}

// Uint64ToKBMBGB 将uint64类型表示的字节转换为带上单位的String
func Uint64ToKBMBGB(value uint64) string {
	var retStr string
	if value > 1073741824 {
		retStr = fmt.Sprintf("%.2f", float64(value)/float64(1024)/float64(1024)/float64(1024)) + "(GB)"
	} else if value > 1048576 {
		retStr = fmt.Sprintf("%.2f", float64(value)/float64(1024)/float64(1024)) + "(MB)"
	} else if value > 1024 {
		retStr = fmt.Sprintf("%.2f", float64(value)/float64(1024)) + "(KB)"
	} else {
		retStr = fmt.Sprintf("%.2f", float64(value)) + "(B)"
	}
	return retStr
}

// Float64ToKBMBGB 将float64类型表示的字节转换为带上单位的String
func Float64ToKBMBGB(value float64) string {
	var retStr string
	if value > 1073741824 {
		retStr = fmt.Sprintf("%.2f", value) + "(GB)"
	} else if value > 1048576 {
		retStr = fmt.Sprintf("%.2f", value) + "(MB)"
	} else {
		retStr = fmt.Sprintf("%.2f", value) + "(KB)"
	}
	return retStr
}

// ContainsString 判断元素是否存在于Slice中, String类型
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if strings.TrimSpace(a) == e {
			return true
		}
	}
	return false
}

// ContainsInt 判断元素是否存在于Slice中, Int类型
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// ContainsInt 判断元素是否存在于Slice中, Int类型
func ContainsUInt(s []uint, e uint) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// CreateRandomString 生成随机字符串
func CreateRandomString(len int) string {
	var container string
	var str = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"
	b := bytes.NewBufferString(str)
	length := b.Len()
	bigInt := big.NewInt(int64(length))
	for i := 0; i < len; i++ {
		randomInt, _ := rand.Int(rand.Reader, bigInt)
		container += string(str[randomInt.Int64()])
	}
	return container
}
