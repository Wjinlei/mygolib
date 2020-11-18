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

func RunCmd(name string, params ...string) error {
	cmd := exec.Command(name, params...)
	if err := cmd.Run(); err != nil {
		return err
	}
	return nil
}

func StartCmd(name string, params ...string) error {
	cmd := exec.Command(name, params...)
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
