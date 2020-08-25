package psutils

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shirou/gopsutil/mem"
)

type SwapStat struct {
	Total       uint64  `json:"total"`
	Available   uint64  `json:"available"`
	Used        uint64  `json:"used"`
	UsedPercent float64 `json:"usedPercent"`
}

func (m SwapStat) String() string {
	s, _ := json.Marshal(m)
	return string(s)
}

func GetSwap() (*SwapStat, error) {
	v, err := mem.SwapMemory()
	if err != nil {
		return nil, err
	}
	empty := &mem.SwapMemoryStat{}
	if v == empty {
		return nil, errors.New(fmt.Sprintf("error %v", v))
	}
	return &SwapStat{
		Total:       v.Total,
		Available:   v.Free,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
	}, nil
}
