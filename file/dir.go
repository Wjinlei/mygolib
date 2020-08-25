package file

import (
	"os"
	"path"
)

// 创建目录
func MakeDir(filepath string) error {
	dirpath := path.Dir(filepath)
	if NotExists(dirpath) {
		if err := os.Mkdir(dirpath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// 创建目录(所有层级)
func MakeDirAll(filepath string) error {
	dirpath := path.Dir(filepath)
	if NotExists(dirpath) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
