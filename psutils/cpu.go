package psutils

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"time"

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

// 获取CPU信息
func GetCPUInfo() ([]CpuStat, error) {
	v, err := cpu.Info()
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New(fmt.Sprintf("could not get CPU Info: %v", v))
	}
	var cpuStats []CpuStat
	for _, vv := range v {
		var cpuStat CpuStat
		if vv.ModelName == "" {
			return nil, errors.New(fmt.Sprintf("could not get CPU Info: %v", vv))
		}
		cpuStat.Cores = vv.Cores
		cpuStat.CPU = vv.CPU
		cpuStat.ModelName = vv.ModelName
		cpuStats = append(cpuStats, cpuStat)
	}
	return cpuStats, nil
}

// 获取CPU使用率
func GetCPUPercent() (float64, error) {
	numcpu := runtime.NumCPU()
	result := 0.0
	v, err := cpu.Percent(time.Second, false)
	if err != nil {
		return 0, err
	}
	for _, percent := range v {
		// Check for slightly greater then 100% to account for any rounding issues.
		if percent < 0.0 || percent > 100.0001*float64(numcpu) {
			return 0, errors.New(fmt.Sprintf("CPUPercent value is invalid: %f", percent))
		}
		result = percent
	}
	return result, nil
}
