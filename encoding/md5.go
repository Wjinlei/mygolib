package encoding

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
)

// 生成给定字符串的md5值
func StringToMd5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	return fmt.Sprintf("%x", w.Sum(nil))
}

// 生成给定文件的md5值
func FileToMd5(path string) (string, error) {
	body, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", md5.Sum(body)), nil
}
