package public

import (
	"fmt"
	"strings"
)

// FilterSQL 过滤可能造成sql注入的字符
func FilterSQL(sql string) string {
	sql = strings.ReplaceAll(sql, "'", "")
	sql = strings.ReplaceAll(sql, "\"", "")
	sql = strings.ReplaceAll(sql, "(", "")
	sql = strings.ReplaceAll(sql, ")", "")
	//sql = strings.ReplaceAll(sql, "*", "") // 由于域名可能会用到*,所以不过滤*
	sql = strings.ReplaceAll(sql, ".", "_")
	sql = strings.ReplaceAll(sql, "[", "")
	sql = strings.ReplaceAll(sql, "]", "")
	sql = strings.ReplaceAll(sql, " ", "")
	sql = strings.ReplaceAll(sql, "=", "")
	sql = strings.ReplaceAll(sql, ">", "")
	sql = strings.ReplaceAll(sql, "<", "")
	sql = strings.ReplaceAll(sql, "%", "")
	sql = strings.ReplaceAll(sql, "\t", "")
	return sql
}

// TrimQuotes 删除源字符的引号
func TrimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}

// ByteToString 将uint64类型表示的字节转换为带上单位的String
func ByteToString(value uint64) string {
	var retStr string
	if value > 1073741824 {
		retStr = fmt.Sprintf("%.2f", float64(value)/float64(1024)/float64(1024)/float64(1024)) + "(GB)"
	} else if value > 1048576 {
		retStr = fmt.Sprintf("%.2f", float64(value)/float64(1024)/float64(1024)) + "(MB)"
	} else {
		retStr = fmt.Sprintf("%.2f", float64(value)) + "(KB)"
	}
	return retStr
}
