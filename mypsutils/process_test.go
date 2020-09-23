package mypsutils

import (
	"fmt"
	"os"
	"testing"
)

func TestNewProcess(t *testing.T) {
	ret, err := NewProcess(int32(os.Getpid()))
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("Pid: %d\n", ret.Pid)
}
