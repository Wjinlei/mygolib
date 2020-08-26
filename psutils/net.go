package psutils

import (
	"errors"
	"fmt"
	"github.com/shirou/gopsutil/net"
	gonet "net"
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

// 获取网络接口
func GetNetInterfaces() ([]net.InterfaceStat, error) {
	v, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New(fmt.Sprintf("Could not get NetInterfaceStat: %v", v))
	}
	for _, vv := range v {
		if vv.Name == "" {
			return nil, errors.New(fmt.Sprintf("Invalid NetInterface: %v", vv))
		}
	}
	return v, nil
}

// 获取接口连接信息
func GetConnections() ([]net.ConnectionStat, error) {
	v, err := net.Connections("inet")
	if err != nil {
		return nil, err
	}
	if len(v) == 0 {
		return nil, errors.New(fmt.Sprintf("Could not get NetConnections: %v", v))
	}
	for _, vv := range v {
		if vv.Family == 0 {
			return nil, errors.New(fmt.Sprintf("invalid NetConnections: %v", vv))
		}
	}
	return v, nil
}

// 获取出口IP,注意不是外网IP,而是数据包流出的接口的IP地址
func GetOutboundIP() (string, error) {
	conn, err := gonet.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*gonet.UDPAddr)
	return localAddr.IP.String(), nil
}
