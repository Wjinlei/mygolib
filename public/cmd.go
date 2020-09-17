package public

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// ExecShell 执行一条命令
func ExecShell(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = getSysctrlEnv(os.Environ())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	defer stdout.Reset()
	defer stderr.Reset()
	err := cmd.Run()
	outStr := strings.TrimSpace(stdout.String())
	errStr := strings.TrimSpace(stderr.String())
	if err != nil {
		return fmt.Sprintf("[ERROR]: %s\n", errStr), err
	}
	if stderr.String() != "" {
		return fmt.Sprintf("[ERROR]: %s\n", errStr), fmt.Errorf(errStr)
	}
	return outStr, nil
}

// ExecScript 执行一个脚本
func ExecScript(params ...string) (string, error) {
	cmd := exec.Command("bash", params...)
	cmd.Env = getSysctrlEnv(os.Environ())
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	defer stdout.Reset()
	defer stderr.Reset()
	err := cmd.Run()
	outStr := strings.TrimSpace(stdout.String())
	errStr := strings.TrimSpace(stderr.String())
	if err != nil {
		return fmt.Sprintf("[ERROR]: %s\n", errStr), err
	}
	if stderr.String() != "" {
		return fmt.Sprintf("[ERROR]: %s\n", errStr), fmt.Errorf(errStr)
	}
	return outStr, nil
}

// DoSysctrl 执行sysctl
func DoSysctrl(mib string) ([]string, error) {
	sysctl, err := exec.LookPath("sysctl")
	if err != nil {
		return []string{}, err
	}
	cmd := exec.Command(sysctl, "-n", mib)
	cmd.Env = getSysctrlEnv(os.Environ())
	out, err := cmd.Output()
	if err != nil {
		return []string{}, err
	}
	v := strings.Replace(string(out), "{ ", "", 1)
	v = strings.Replace(string(v), " }", "", 1)
	values := strings.Fields(string(v))
	return values, nil
}
