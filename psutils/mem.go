package psutils

import (
	"fmt"

	"github.com/Wjinlei/mygolib/public"
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
		Info:        public.ByteToString(v.Used) + "/" + public.ByteToString(v.Total),
		Data:        int(v.UsedPercent),
	}
	return res, nil
}
