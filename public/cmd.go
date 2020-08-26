package public

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// 执行一条命令
func ExecShell(command string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = getSysctrlEnv(os.Environ())
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	defer stdout.Reset()
	defer stderr.Reset()
	err := cmd.Run()
	if err != nil {
		errOut := strings.TrimSpace(stderr.String())
		return fmt.Sprintf("[DEBUG]: %s [ERROR]: %s [INFO]: %s\n", command, err, errOut), err
	}
	return strings.TrimSpace(stdout.String()), nil
}

// 执行一个脚本
func ExecScript(params ...string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", params...)
	cmd.Env = getSysctrlEnv(os.Environ())
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	defer stdout.Reset()
	defer stderr.Reset()
	err := cmd.Run()
	if err != nil {
		errOut := strings.TrimSpace(stderr.String())
		return fmt.Sprintf("[DEBUG]: %s [ERROR]: %s [INFO]: %s\n", params, err, errOut), err
	}
	return strings.TrimSpace(stdout.String()), nil
}

// 执行sysctl
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
