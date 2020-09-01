package psutils

import "fmt"

// toString 将Used和Total转换为字符串
func toString(used uint64, total uint64) string {
	var totalStr, usedStr string
	if total > 1073741824 {
		totalStr = fmt.Sprintf("%.2f", float64(total)/float64(1024)/float64(1024)/float64(1024)) + "(GB)"
	} else if total > 1048576 {
		totalStr = fmt.Sprintf("%.2f", float64(total)/float64(1024)/float64(1024)) + "(MB)"
	} else {
		totalStr = fmt.Sprintf("%.2f", float64(total)) + "(KB)"
	}
	if used > 1073741824 {
		usedStr = fmt.Sprintf("%.2f", float64(used)/float64(1024)/float64(1024)/float64(1024)) + "(GB)"
	} else if used > 1048576 {
		usedStr = fmt.Sprintf("%.2f", float64(used)/float64(1024)/float64(1024)) + "(MB)"
	} else {
		usedStr = fmt.Sprintf("%.2f", float64(used)) + "(KB)"
	}
	return usedStr + "/" + totalStr
}
