package public

import (
	"os"
	"strings"
)

// 判断文件是否存在
func PathExists(path string) bool {
	if _, err := os.Stat(path); err == nil {
		return true
	}
	return false
}

// 判断路径是否是文件夹
func IsDir(path string) (bool, error) {
	s, err := os.Stat(path)
	if err != nil {
		return false, err
	}
	return s.IsDir(), nil
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
