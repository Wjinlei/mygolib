package cmd

import "testing"

func TestExecShell(t *testing.T) {
	_, err := ExecShell("ls -ll")
	if err != nil {
		t.Error(err)
	}
}

func TestExecScript(t *testing.T) {
	_, err := ExecScript("echoHello.sh")
	if err != nil {
		t.Error(err)
	}
}
