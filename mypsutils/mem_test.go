package mypsutils

import (
	"fmt"
	"testing"
)

func TestGetMemory(t *testing.T) {
	v, err := GetMemory()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
