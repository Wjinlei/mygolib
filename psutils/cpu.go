package psutils

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/shirou/gopsutil/cpu"
)

// GetCPUInfo 获取CPU信息
func GetCPUInfo() ([]cpu.InfoStat, error) {
	v, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, fmt.Errorf("could not get CPU Info: %v", v)
	}
	return v, nil
}

// GetCPUPercent 获取CPU使用率
func GetCPUPercent() (*ResStat, error) {
	numcpu := runtime.NumCPU()
	result := 0.0
	v, err := cpu.Percent(time.Second, false)
	if err != nil {
		return nil, err
	}
	for _, percent := range v {
		// Check for slightly greater then 100% to account for any rounding issues.
		if percent < 0.0 || percent > 100.0001*float64(numcpu) {
			return nil, fmt.Errorf("CPUPercent value is invalid: %f", percent)
		}
		result = percent
	}
	// 获取CPU核心数
	cpuCount, err := GetCPUCount(true)
	if err != nil {
		return nil, err
	}
	res := &ResStat{
		UsedPercent: result,
		Title:       "CPU使用率",
		Info:        strconv.Itoa(cpuCount) + "核心",
		Data1:       int(result),
	}
	return res, nil
}

// GetCPUCount 获取CPU个数, logical = true 则统计逻辑核心(线程)
func GetCPUCount(logical bool) (int, error) {
	v, err := cpu.Counts(logical)
	if err != nil {
		return 0, err
	}
	if v == 0 {
		return 0, fmt.Errorf("could not get CPU counts: %v", v)
	}
	return v, nil
}
