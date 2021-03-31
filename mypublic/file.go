package mypublic

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// DownloadFile 下载文件
func DownloadFile(url string, path string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	// 目标文件
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	// 保存响应数据到文件
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// WriteFile 创建文件
func WriteFile(filepath string, content string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(content)
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

// Delete 删除文件或目录,如果是目录,只能删除空目录
func Delete(path string) error {
	return os.Remove(path)
}

// DeleteAll 删除文件或目录
func DeleteAll(path string) error {
	return os.RemoveAll(path)
}

// Move 移动文件或目录,如果目标已存在,并且不是目录,则会覆盖
func Move(oldpath string, newpath string) error {
	return os.Rename(oldpath, newpath)
}

// 复制文件或目录
func Copy(oldpath string, newpath string) error {
	oldfileInfo, err := os.Lstat(oldpath)
	if err != nil {
		return err
	}
	if oldfileInfo.IsDir() {
		return copyD(oldpath, newpath)
	} else {
		return copyF(oldpath, newpath)
	}
}

// ReadFile 读取文件内容
func ReadFile(filename string) (string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// ReadLines 读所有行, ReadLinesOffsetN简单封装
func ReadLines(filename string) ([]string, error) {
	return ReadLinesOffsetN(filename, 0, -1)
}

// ReadLinesOffsetN 读几行, offset表示从第几行开始读, n表示读几行, 返回读取到的行的Slice
func ReadLinesOffsetN(filename string, offset uint, n int) ([]string, error) {
	f, err := os.Open(filename)
	if err != nil {
		return []string{""}, err
	}
	defer f.Close()
	var ret []string
	r := bufio.NewReader(f)
	for i := 0; i < n+int(offset) || n < 0; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			break
		}
		if i < int(offset) {
			continue
		}
		ret = append(ret, strings.Trim(line, "\n"))
	}
	return ret, nil
}

// DirSize 获取目录大小
func DirSize(path string) (int64, error) {
	var size int64
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

func IsDir(path string) bool {
	s, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// copyF 复制文件
func copyF(oldpath string, newpath string) error {
	symlink, err := os.Readlink(oldpath)
	if err != nil {
		// 源文件
		oldfile, err := os.Open(oldpath)
		if err != nil {
			return err
		}
		defer oldfile.Close()

		// 目标文件
		newfile, err := os.OpenFile(newpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
		if err != nil {
			return err
		}
		defer newfile.Close()

		// 保存响应数据到文件
		_, err = io.Copy(newfile, oldfile)
		if err != nil {
			return err
		}
	} else {
		err := os.Symlink(symlink, newpath)
		if err != nil {
			return err
		}
	}
	return nil
}

// copyD 复制目录
func copyD(oldpath string, newpath string) error {
	// 创建目标目录
	if err := MakeDir(newpath); err != nil {
		return err
	}

	// 打开源目录
	oldDir, err := os.Open(oldpath)
	if err != nil {
		return err
	}
	defer oldDir.Close()

	// 读取目录中的文件信息
	fileStats, err := oldDir.Readdir(-1)
	if err != nil {
		return err
	}

	// 处理目录下的内容
	for _, fileStat := range fileStats {
		fsrc := fmt.Sprintf("%s/%s", oldpath, fileStat.Name())
		fdst := fmt.Sprintf("%s/%s", newpath, fileStat.Name())
		if fileStat.IsDir() {
			// 递归创建子目录
			if err := copyD(fsrc, fdst); err != nil {
				return err
			}
		} else {
			// 复制文件
			if err := copyF(fsrc, fdst); err != nil {
				return err
			}
		}
	}
	return nil
}
