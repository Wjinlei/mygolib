package file

import "os"

// 判断是否存在
func Exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

// 判断是否不存在
func NotExists(path string) bool {
	_, err := os.Stat(path)
	return err != nil || os.IsNotExist(err)
}
