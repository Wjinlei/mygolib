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
