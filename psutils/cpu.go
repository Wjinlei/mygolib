package psutils

import (
	"encoding/json"
	"errors"

	"github.com/shirou/gopsutil/cpu"
)

type CpuStat struct {
	Cores     int32  `json:"cores"`
	CPU       int32  `json:"cpu"`
	ModelName string `json:"modelName"`
}

func (c CpuStat) String() string {
	s, _ := json.Marshal(c)
	return string(s)
}

func GetCpu() ([]CpuStat, error) {
	v, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New("could not get CPU Info")
	}
	var cpuStats []CpuStat
	for _, vv := range v {
		var cpuStat CpuStat
		if vv.ModelName == "" {
			return nil, errors.New("could not get CPU Info")
		}
		cpuStat.Cores = vv.Cores
		cpuStat.CPU = vv.CPU
		cpuStat.ModelName = vv.ModelName
		cpuStats = append(cpuStats, cpuStat)
	}
	return cpuStats, nil
}
