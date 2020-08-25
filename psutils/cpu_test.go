package psutils

import (
	"fmt"
	"testing"
)

func TestGetCPUInfo(t *testing.T) {
	v, err := GetCPUInfo()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetCPUPercent(t *testing.T) {
	v, err := GetCPUPercent()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetCPUCount(t *testing.T) {
	v, err := GetCPUCount(false)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
