package public

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

// 执行一条命令
func ExecShell(command string) (string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
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
