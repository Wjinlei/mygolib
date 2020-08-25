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
