package mypublic

import (
	"errors"
	"fmt"
	"os"
	"path"
	"path/filepath"
)

// CopyDir 复制目录
func CopyDir(oldpath string, newpath string) error {
	// 获取源目录信息
	oldfileStat, err := os.Stat(oldpath)
	if err != nil {
		return err
	}

	// 创建目标目录
	if err := os.MkdirAll(newpath, oldfileStat.Mode()); err != nil {
		return err
	}

	// 打开源目录
	olddir, err := os.Open(oldpath)
	if err != nil {
		return err
	}
	defer olddir.Close()

	// 读取目录中的文件信息
	fileStats, err := olddir.Readdir(-1)
	if err != nil {
		return err
	}

	var errs []error // 用来保存错误

	for _, fileStat := range fileStats {
		fsrc := fmt.Sprintf("%s/%s", oldpath, fileStat.Name())
		fdst := fmt.Sprintf("%s/%s", newpath, fileStat.Name())
		if fileStat.IsDir() {
			// 递归创建子目录
			err = CopyDir(fsrc, fdst)
			if err != nil {
				errs = append(errs, err)
			}
		} else {
			// 复制文件
			err = CopyFile(fsrc, fdst)
			if err != nil {
				errs = append(errs, err)
			}
		}
	}

	// 处理错误信息
	var errString string
	for _, err := range errs {
		errString += err.Error() + "\n"
	}
	if errString != "" {
		return errors.New(errString)
	}

	return nil
}

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

// 判断目录大小
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

// 判断所给路径是否为文件夹
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}
