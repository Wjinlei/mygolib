package psutils

import (
	"fmt"
	"testing"
)

func TestGetDiskUsage(t *testing.T) {
	v, err := GetDiskUsage("/")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetDiskPart(t *testing.T) {
	v, err := GetDiskPart()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
