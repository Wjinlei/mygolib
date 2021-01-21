package mypublic

import "math"

// Round 四舍五入,保留指定位数的小数
func Round(f float64, n int) float64 {
	n10 := math.Pow10(n)
	return math.Trunc((f+0.5/n10)*n10) / n10
}
