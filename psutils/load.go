package psutils

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/load"
)

func GetLoadAvg() (*load.AvgStat, error) {
	v, err := load.Avg()
	if err != nil {
		return nil, err
	}
	empty := &load.AvgStat{}
	if v == empty {
		return nil, errors.New(fmt.Sprintf("get LoadAvg failed: %v", v))
	}
	return v, nil
}
