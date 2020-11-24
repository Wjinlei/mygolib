package mypublic

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// MakeDir 创建目录
func MakeDir(dirpath string) error {
	if !Exists(dirpath) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

// MakeDirAll 创建目录(所有层级)
func MakeDirAll(filepath string) error {
	dirpath := path.Dir(filepath)
	if !Exists(dirpath) {
		if err := os.MkdirAll(dirpath, os.ModePerm); err != nil {
			return err
		}
	}
	return nil
}

func DirSize(path string) (int64, error) {
	var size int64
	if !Exists(path) {
		return 0, fmt.Errorf("目录不存在")
	}
	err := filepath.Walk(path, func(_ string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			size += info.Size()
		}
		return err
	})
	if err != nil {
		return 0, err
	}
	return size, nil
}
