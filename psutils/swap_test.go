package psutils

import (
	"fmt"
	"testing"
)

func TestGetSwap(t *testing.T) {
	v, err := GetSwap()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
