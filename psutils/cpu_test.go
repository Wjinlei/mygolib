package psutils

import (
	"fmt"
	"testing"
)

func TestGetCpu(t *testing.T) {
	v, err := GetCpu()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
