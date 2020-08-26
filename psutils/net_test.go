package psutils

import (
	"fmt"
	"testing"
)

func TestGetIOCountersAll(t *testing.T) {
	v, err := GetIOCountersAll()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
