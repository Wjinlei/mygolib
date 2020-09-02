package psutils

import (
	"fmt"
	"testing"
)

func TestGetUsage(t *testing.T) {
	v, err := GetUsage("/")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetPartitions(t *testing.T) {
	v, err := GetPartitions()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
