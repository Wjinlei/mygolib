package psutils

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/host"
)

func GetHostInfo() (*host.InfoStat, error) {
	v, err := host.Info()
	if err != nil {
		return nil, err
	}
	empty := &host.InfoStat{}
	if v == empty {
		return nil, errors.New(fmt.Sprintf("Could not get hostinfo: %v", v))
	}
	if v.Procs == 0 {
		return nil, errors.New("Could not determine the number of host processes")
	}
	return v, nil
}
