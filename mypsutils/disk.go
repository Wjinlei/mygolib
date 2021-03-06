package mypsutils

import (
	"errors"
	"fmt"

	"github.com/Wjinlei/mygolib/mypublic"
	"github.com/shirou/gopsutil/disk"
)

// GetUsage 获取指定路径的磁盘使用率
func GetUsage(path string) (*ResStat, error) {
	v, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	if v.Path != path {
		return nil, fmt.Errorf("path does not match, yourpath: %s, getpath: %s", path, v.Path)
	}
	res := &ResStat{
		Total:       v.Total,
		Available:   v.Free,
		Used:        v.Used,
		UsedPercent: v.UsedPercent,
		Title:       path,
		Info:        mypublic.Uint64ToKBMBGB(v.Used) + "/" + mypublic.Uint64ToKBMBGB(v.Total),
		Data:        int(v.UsedPercent),
	}
	return res, nil
}

// GetPartitions 获取磁盘分区
func GetPartitions(all bool) ([]disk.PartitionStat, error) {
	ret, err := disk.Partitions(all)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("Not found partitions")
	}
	empty := disk.PartitionStat{}
	for _, disk := range ret {
		if disk == empty {
			return nil, fmt.Errorf("Could not get device info %v", disk)
		}
	}
	return ret, nil
}
