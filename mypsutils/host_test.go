package mypsutils

import (
	"fmt"
	"testing"
)

func TestGetHostInfo(t *testing.T) {
	v, err := GetHostInfo()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetUptime(t *testing.T) {
	v, err := GetUptime()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetUsers(t *testing.T) {
	v, err := GetUsers()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestKernelVersion(t *testing.T) {
	v, err := GetKernelVersion()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetPlatformInfo(t *testing.T) {
	platform, family, version, err := GetPlatformInfo()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("platform: %s, family: %s, version: %s\n", platform, family, version)
}

func TestGetHostId(t *testing.T) {
	v, err := GetHostID()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
