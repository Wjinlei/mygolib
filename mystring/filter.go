package mystring

import "strings"

// 过滤可能造成sql注入的字符
func FilterSql(sql string) string {
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
	return strings.TrimSpace(sql)
}
