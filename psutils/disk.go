package psutils

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/disk"
)

// 获取指定路径的磁盘使用率
func GetDiskUsage(path string) (*disk.UsageStat, error) {
	v, err := disk.Usage(path)
	if err != nil {
		return nil, err
	}
	if v.Path != path {
		return nil, errors.New(fmt.Sprintf("path does not match, yourpath: %s, getpath: %s", path, v.Path))
	}
	return v, nil
}

// 获取磁盘分区
func GetDiskPart() ([]disk.PartitionStat, error) {
	ret, err := disk.Partitions(false)
	if err != nil {
		return nil, err
	}
	if len(ret) == 0 {
		return nil, errors.New("Not found partitions")
	}
	empty := disk.PartitionStat{}
	for _, disk := range ret {
		if disk == empty {
			return nil, errors.New(fmt.Sprintf("Could not get device info %v", disk))
		}
	}
	return ret, nil
}
