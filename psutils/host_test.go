package psutils

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
