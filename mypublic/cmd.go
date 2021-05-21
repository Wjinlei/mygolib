package mypublic

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

// ExecShell 执行一条命令
func ExecShell(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	//cmd.Env = GetSysctrlEnv(os.Environ())
	out, err := cmd.CombinedOutput()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] ExecShell: %s\n%s", command, out))
	}
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

// ExecScript 执行一个脚本
func ExecScript(params ...string) (string, error) {
	cmd := exec.Command("bash", params...)
	//cmd.Env = GetSysctrlEnv(os.Environ())
	out, err := cmd.CombinedOutput()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] ExecScript: %s\n%s", strings.Join(params, " "), out))
	}
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func RunCmd(name string, params ...string) (string, error) {
	cmd := exec.Command(name, params...)
	//cmd.Env = GetSysctrlEnv(os.Environ())
	out, err := cmd.CombinedOutput()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] Run: %s %s\n%s", name, strings.Join(params, " "), out))
	}
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func StartCmd(name string, params ...string) error {
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] Start: %s %s", name, strings.Join(params, " ")))
	}
	cmd := exec.Command(name, params...)
	if err := cmd.Start(); err != nil {
		return err
	}
	return nil
}
