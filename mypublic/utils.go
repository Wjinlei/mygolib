package mypublic

import (
	"os"
	"strings"
)

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// 判断所给路径是否为文件
func IsFile(path string) bool {
	return !IsDir(path)
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
