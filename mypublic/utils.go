package mypublic

import (
	"os"
	"strings"

	"github.com/mozillazg/go-pinyin"
)

// 中文转拼音
func Pinyin(s string) string {
	pinyinArgs := pinyin.NewArgs()
	pinyinArgs.Style = pinyin.FirstLetter
	pys := []string{}
	for _, r := range s {
		py := pinyin.SinglePinyin(r, pinyinArgs)
		if len(py) > 0 {
			pys = append(pys, py[0])
		} else {
			pys = append(pys, string(r))
		}
	}
	ns := strings.Join(pys, "")
	return ns
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	if _, err := os.Lstat(path); err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
