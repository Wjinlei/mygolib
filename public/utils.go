package public

import (
	"os"
	"strings"
)

// 判断文件是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 判断文件是否不存在
func NotExists(path string) bool {
	_, err := os.Stat(path)
	return err != nil || os.IsNotExist(err)
}

// 判断元素是否存在于Slice中, String类型
func ContainsString(s []string, e string) bool {
	for _, a := range s {
		if strings.TrimSpace(a) == e {
			return true
		}
	}
	return false
}

// 判断元素是否存在于Slice中, Int类型
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
