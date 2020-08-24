package file

import (
	"io/ioutil"
	"os"
)

// 写入字符串到文件
func WriteFile(filepath string, content string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(content)
	return nil
}

// 读取文件内容
func ReadFile(filePath string) (string, error) {
	file, err := os.OpenFile(filePath, os.O_RDONLY, 0644)
	if err != nil {
		return "", err
	}
	defer file.Close()
	content, err := ioutil.ReadAll(file)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// 判断文件是否存在
func FileExists(filePath string) bool {
	_, err := os.Stat(filePath)
	ok := os.IsNotExist(err)
	if ok {
		return false
	}
	return true
}

// 判断文件是否不存在
func FileNotExists(filePath string) bool {
	_, err := os.Stat(filePath)
	ok := os.IsNotExist(err)
	if ok {
		return true
	}
	return false
}

// 移动或重命名文件
func MoveFile(oldpath string, newpath string) error {
	if err := os.Rename(oldpath, newpath); err != nil {
		return err
	}
	return nil
}
