package psutils

import (
	"encoding/json"
	"github.com/shirou/gopsutil/mem"
)

type MemoryStat struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func (m MemoryStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}

func GetMemory() (*MemoryStat, error) {
	v, err := mem.VirtualMemory()
	if err != nil {
		return nil, err
	}
	return &MemoryStat{
		Total:       v.Total,
		Available:   v.Available,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
	}, nil
}
