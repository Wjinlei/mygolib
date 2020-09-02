package psutils

import (
	"fmt"

	"github.com/Wjinlei/mygolib/public"
	"github.com/shirou/gopsutil/mem"
)

// GetSwap 获取交换空间使用率
func GetSwap() (*ResStat, error) {
	v, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}
	empty := &mem.SwapMemoryStat{}
	if v == empty {
		return nil, fmt.Errorf("error %v", v)
	}
	res := &ResStat{
		Total:       v.Total,
		Available:   v.Free,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
		Title:       "swap(交换空间)",
		Info:        public.Uint64ToKBMBGB(v.Used) + "/" + public.Uint64ToKBMBGB(v.Total),
		Data:        int(v.UsedPercent),
	}
	return res, nil
}
