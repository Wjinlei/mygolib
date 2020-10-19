package mypublic

import (
	"fmt"
	"testing"
)

func TestExecShell(t *testing.T) {
	out, err := ExecShell("ls -ll")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
}

func TestExecScript(t *testing.T) {
	out, err := ExecScript("./.sayHello.sh")
	if err != nil {
		t.Error(err)
	}
	fmt.Println(out)
}
