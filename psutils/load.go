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
		return nil, errors.New(fmt.Sprintf("Get loadavg failed: %v", v))
	}
	return v, nil
}

func GetMisc() (*load.MiscStat, error) {
	v, err := load.Misc()
	if err != nil {
		return nil, err
	}
	empty := &load.MiscStat{}
	if v == empty {
		return nil, errors.New(fmt.Sprintf("Get misc failed: %v", v))
	}
	return v, nil
}
