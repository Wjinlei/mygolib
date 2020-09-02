package public

import (
	"os"
	"path"
)

// MakeDir 创建目录
func MakeDir(filepath string) error {
	dirpath := path.Dir(filepath)
	if !PathExists(dirpath) {
		if err := os.Mkdir(dirpath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// MakeDirAll 创建目录(所有层级)
func MakeDirAll(filepath string) error {
	dirpath := path.Dir(filepath)
	if !PathExists(dirpath) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}
