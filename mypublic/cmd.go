package mypublic

import (
	"os"
	"os/exec"
)

// ExecShell 执行一条命令
func ExecShell(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = GetSysctrlEnv(os.Environ())
	out, err := cmd.Output()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

// ExecScript 执行一个脚本
func ExecScript(params ...string) (string, error) {
	cmd := exec.Command("bash", params...)
	cmd.Env = GetSysctrlEnv(os.Environ())
	out, err := cmd.Output()
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}
