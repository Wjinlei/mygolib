package psutils

import (
	"fmt"
	"runtime"

	"github.com/shirou/gopsutil/load"
)

// GetLoadAvg 获取平均负载
func GetLoadAvg() (*ResStat, error) {
	// 获取负载情况
	v, err := load.Avg()
	if err != nil {
		return nil, err
	}
	empty := &load.AvgStat{}
	if v == empty {
		return nil, fmt.Errorf("Get loadavg failed: %v", v)
	}
	// 计算平均负载 https://www.tecmint.com/understand-linux-load-averages-and-monitor-performance/
	cpuCount := runtime.NumCPU()
	loadAvg := (cpuCount*100 + (int(v.Load1)-cpuCount)*100) / cpuCount / cpuCount
	loadAvgInfo := toStringLoadAvg(loadAvg)
	res := &ResStat{
		Load1:  v.Load1,
		Load5:  v.Load5,
		Load15: v.Load15,
		Title:  "平均负载",
		Info:   loadAvgInfo,
		Data1:  loadAvg,
	}
	return res, nil
}

// GetMisc 获取扩展信息
func GetMisc() (*load.MiscStat, error) {
	v, err := load.Misc()
	if err != nil {
		return nil, err
	}
	empty := &load.MiscStat{}
	if v == empty {
		return nil, fmt.Errorf("Get misc failed: %v", v)
	}
	return v, nil
}

// toStringLoadAvg 转换loadAvg为字符串
func toStringLoadAvg(loadAvg int) string {
	var ret string
	if loadAvg > 90 {
		ret = "运行堵塞"
	} else if loadAvg > 60 {
		ret = "运行缓慢"
	} else if loadAvg > 30 {
		ret = "运行正常"
	} else {
		ret = "运行流畅"
	}
	return ret
}
