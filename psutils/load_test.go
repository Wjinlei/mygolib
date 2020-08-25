package psutils

import (
	"fmt"
	"testing"
)

func TestGetLoadAvg(t *testing.T) {
	v, err := GetLoadAvg()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
