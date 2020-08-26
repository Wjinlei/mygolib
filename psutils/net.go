package psutils

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/net"
)

// 获取所有接口总IO
func GetIOCountersAll() ([]net.IOCountersStat, error) {
	v, err := net.IOCounters(false)
	if err != nil {
		return nil, err
	}
	if len(v) != 1 {
		return nil, errors.New(fmt.Sprintf("Could not get NetIOCounters: %v", v))
	}
	if v[0].Name != "all" {
		return nil, errors.New(fmt.Sprintf("Invalid NetIOCounters: %v", v))
	}
	per, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	var pr uint64
	for _, p := range per {
		pr += p.PacketsRecv
	}
	if v[0].PacketsRecv != pr {
		return nil, errors.New(fmt.Sprintf("invalid sum value: %v, %v", v[0].PacketsRecv, pr))
	}
	return v, nil
}

// 获取每个接口的IO
func GetIOCounters() ([]net.IOCountersStat, error) {
	v, err := net.IOCounters(true)
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New(fmt.Sprintf("Could not get NetIOCounters: %v", v))
	}
	for _, vv := range v {
		if vv.Name == "" {
			return nil, errors.New(fmt.Sprintf("Invalid NetIOCounters: %v", vv))
		}
	}
	return v, nil
}
