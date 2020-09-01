package psutils

import (
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

// GetMemory 获取内存资源
func GetMemory() (*ResStat, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	empty := &mem.VirtualMemoryStat{}
	if v == empty {
		return nil, fmt.Errorf("error %v", v)
	}
	res := &ResStat{
		Total:       v.Total,
		Available:   v.Free,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
		Title:       "内存使用率",
		Info:        toString(v.Used, v.Total),
		Data1:       int(v.UsedPercent),
	}
	return res, nil
}
