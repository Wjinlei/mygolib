package psutils

import (
	"fmt"
	"testing"
)

func TestGetOSRelease(t *testing.T) {
	platform, version, err := GetOSRelease()
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("platform: %s, version: %s\n", platform, version)
}
