package mypsutils

import (
	"testing"
)

func TestCheckEmploy(t *testing.T) {
	Employ, info := CheckEmploy("tcp", "127.0.0.1", 8080)
	if Employ == false {
		t.Error("检测到端口被占用,占用进程信息", info)
	}
}
