package mypublic

import (
	"bytes"
	"crypto/rand"
	"fmt"
	"math/big"
	"regexp"
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

func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// Int64ToKBMBGB 将int64类型表示的字节转换为带上单位的String
func Int64ToKBMBGB(value int64) string {
	var retStr string
	if value > 1024*1024*1024*1024 {
		retStr = fmt.Sprintf("%.1f", float64(value)/float64(1024)/float64(1024)/float64(1024)/float64(1024)) + " TB"
	} else if value > 1024*1024*1024 {
		retStr = fmt.Sprintf("%.1f", float64(value)/float64(1024)/float64(1024)/float64(1024)) + " GB"
	} else if value > 1024*1024 {
		retStr = fmt.Sprintf("%.1f", float64(value)/float64(1024)/float64(1024)) + " MB"
	} else if value > 1024 {
		retStr = fmt.Sprintf("%.1f", float64(value)/float64(1024)) + " KB"
	} else {
		retStr = fmt.Sprintf("%.1f", float64(value)) + " B"
	}
	return retStr
}

// Uint64ToKBMBGB 将uint64类型表示的字节转换为带上单位的String
func Uint64ToKBMBGB(value uint64) string {
	var retStr string
	if value > 1024*1024*1024*1024 {
		retStr = fmt.Sprintf("%.1f", float64(value)/float64(1024)/float64(1024)/float64(1024)/float64(1024)) + " TB"
	} else if value > 1024*1024*1024 {
		retStr = fmt.Sprintf("%.1f", float64(value)/float64(1024)/float64(1024)/float64(1024)) + " GB"
	} else if value > 1024*1024 {
		retStr = fmt.Sprintf("%.1f", float64(value)/float64(1024)/float64(1024)) + " MB"
	} else if value > 1024 {
		retStr = fmt.Sprintf("%.1f", float64(value)/float64(1024)) + " KB"
	} else {
		retStr = fmt.Sprintf("%.1f", float64(value)) + " B"
	}
	return retStr
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

func GetPath(path string) string {
	// 把正斜杠替换为反斜杠
	path = strings.ReplaceAll(path, "\\", "/")
	// 正则匹配规则
	re, _ := regexp.Compile("/+")
	// 替换多个"///"为一个"/"
	path = re.ReplaceAllLiteralString(path, "/")
	// 取出最右边(末尾)的所有"/"
	path = strings.TrimRight(path, "/")
	// 如果最终结果为"",则返回"/"
	if path == "" {
		path = "/"
	}
	return path
}
