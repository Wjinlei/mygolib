package mypublic

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

// ExecShell 执行一条命令
func ExecShell(command string) (string, error) {
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = GetSysctrlEnv(os.Environ()) // 去除所有本地化的设置,让命令可以正确执行
	out, err := cmd.CombinedOutput()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] ExecShell: %s\n%s", command, out))
	}
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

// ExecShell 执行一条命令,不去除本地化内容
func ExecShell2(command string) (string, error) {
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
	cmd.Env = GetSysctrlEnv(os.Environ()) // 去除所有本地化的设置,让命令可以正确执行
	out, err := cmd.CombinedOutput()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] ExecScript: %s\n%s", strings.Join(params, " "), out))
	}
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

// ExecScript2 执行一个脚本,不去除本地化内容
func ExecScript2(params ...string) (string, error) {
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
	cmd.Env = GetSysctrlEnv(os.Environ()) // 去除所有本地化的设置,让命令可以正确执行
	out, err := cmd.CombinedOutput()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] Run: %s %s\n%s", name, strings.Join(params, " "), out))
	}
	if err != nil {
		return string(out), err
	}
	return string(out), nil
}

func RunCmd2(name string, params ...string) (string, error) {
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
	cmd := exec.Command(name, params...)
	cmd.Env = GetSysctrlEnv(os.Environ()) // 去除所有本地化的设置,让命令可以正确执行
	err := cmd.Start()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] Start: %s %s\n", name, strings.Join(params, " ")))
	}
	if err != nil {
		return err
	}
	return nil
}

func StartCmd2(name string, params ...string) error {
	cmd := exec.Command(name, params...)
	//cmd.Env = GetSysctrlEnv(os.Environ())
	err := cmd.Start()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] Start: %s %s\n", name, strings.Join(params, " ")))
	}
	if err != nil {
		return err
	}
	return nil
}

// StartShell 执行一条命令,不等待结果
func StartShell(command string) error {
	cmd := exec.Command("bash", "-c", command)
	cmd.Env = GetSysctrlEnv(os.Environ()) // 去除所有本地化的设置,让命令可以正确执行
	err := cmd.Start()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] StartShell: %s\n", command))
	}
	if err != nil {
		return err
	}
	return nil
}

// StartShell 执行一条命令,不等待结果
func StartShell2(command string) error {
	cmd := exec.Command("bash", "-c", command)
	//cmd.Env = GetSysctrlEnv(os.Environ()) // 去除所有本地化的设置,让命令可以正确执行
	err := cmd.Start()
	if Exists("debug") {
		log.Println(fmt.Sprintf("[DEBUG] StartShell: %s\n", command))
	}
	if err != nil {
		return err
	}
	return nil
}
