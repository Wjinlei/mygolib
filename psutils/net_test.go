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

func TestGetIOCounters(t *testing.T) {
	v, err := GetIOCounters()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetNetInterfaces(t *testing.T) {
	v, err := GetNetInterfaces()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}

func TestGetConnections(t *testing.T) {
	v, err := GetConnections()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(v)
}
