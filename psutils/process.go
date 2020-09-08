package psutils

import (
	"github.com/shirou/gopsutil/process"
)

// NewProcess 产生新进程
func NewProcess(pid int32) (*process.Process, error) {
	ret, err := process.NewProcess(pid)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
