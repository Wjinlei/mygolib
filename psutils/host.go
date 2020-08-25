package psutils

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/host"
)

// 获取主机信息
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

// 获取系统启动时间
func GetUptime() (uint64, error) {
	v, err := host.Uptime()
	if err != nil {
		return 0, err
	}
	if v == 0 {
		return 0, errors.New(fmt.Sprintf("Could not get up time: %v", v))
	}
	return v, nil
}

// 获取系统用户列表
func GetUsers() ([]host.UserStat, error) {
	v, err := host.Users()
	if err != nil {
		return nil, err
	}
	empty := host.UserStat{}
	if len(v) == 0 {
		return nil, errors.New("Users is empty")
	}
	for _, u := range v {
		if u == empty {
			return nil, errors.New(fmt.Sprintf("Could not User: %v", v))
		}
	}
	return v, nil
}

// 获取系统内核版本
func GetKernelVersion() (string, error) {
	version, err := host.KernelVersion()
	if err != nil {
		return "", err
	}
	if version == "" {
		return "", errors.New(fmt.Sprintf("KernelVersion() returns empty: %s", version))
	}
	return version, nil
}

// 获取平台信息
func GetPlatformInfo() (platform string, family string, version string, err error) {
	platform, family, version, err = host.PlatformInformation()
	if err != nil {
		return "", "", "", err
	}
	if platform == "" {
		return "", "", "", errors.New(fmt.Sprintf("PlatformInformation() returns empty: %s", platform))
	}
	return platform, family, version, nil
}

// 获取主机ID
func GetHostId() (string, error) {
	v, err := GetHostInfo()
	if err != nil {
		return "", err
	}
	if v.HostID == "" {
		return "", errors.New("Host id is empty")
	}
	return v.HostID, nil
}
