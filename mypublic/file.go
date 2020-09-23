package mypublic

import (
	"bufio"
	"io"
	"io/ioutil"
	"net/http"
	"os"
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

// WriteFile 写入字符串到文件
func WriteFile(filepath string, content string) error {
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	file.WriteString(content)
	return nil
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

// MoveFile 移动或重命名文件
func MoveFile(oldpath string, newpath string) error {
	if err := os.Rename(oldpath, newpath); err != nil {
		return err
	}
	return nil
}
