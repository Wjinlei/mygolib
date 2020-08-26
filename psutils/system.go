package psutils

import (
	"mylib/public"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

//GetEnv retrieves the environment variable key. If it does not exist it returns the default.
func GetEnv(key string, dfault string, combineWith ...string) string {
	value := os.Getenv(key)
	if value == "" {
		value = dfault
	}
	switch len(combineWith) {
	case 0:
		return value
	case 1:
		return filepath.Join(value, combineWith[0])
	default:
		all := make([]string, len(combineWith)+1)
		all[0] = value
		copy(all[1:], combineWith)
		return filepath.Join(all...)
	}
}

func HostProc(combineWith ...string) string {
	return GetEnv("HOST_PROC", "/proc", combineWith...)
}

func HostSys(combineWith ...string) string {
	return GetEnv("HOST_SYS", "/sys", combineWith...)
}

func HostEtc(combineWith ...string) string {
	return GetEnv("HOST_ETC", "/etc", combineWith...)
}

func HostVar(combineWith ...string) string {
	return GetEnv("HOST_VAR", "/var", combineWith...)
}

func HostRun(combineWith ...string) string {
	return GetEnv("HOST_RUN", "/run", combineWith...)
}

func HostDev(combineWith ...string) string {
	return GetEnv("HOST_DEV", "/dev", combineWith...)
}

// getSysctrlEnv sets LC_ALL=C in a list of env vars for use when running
// sysctl commands (see DoSysctrl).
func getSysctrlEnv(env []string) []string {
	foundLC := false
	for i, line := range env {
		if strings.HasPrefix(line, "LC_ALL") {
			env[i] = "LC_ALL=C"
			foundLC = true
		}
	}
	if !foundLC {
		env = append(env, "LC_ALL=C")
	}
	return env
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

// 获取系统发行版本
func GetOSRelease() (platform string, version string, err error) {
	contents, err := public.ReadLines(HostEtc("os-release"))
	if err != nil {
		return "", "", nil // return empty
	}
	for _, line := range contents {
		field := strings.Split(line, "=")
		if len(field) < 2 {
			continue
		}
		switch field[0] {
		case "ID": // use ID for lowercase
			platform = trimQuotes(field[1])
		case "VERSION":
			version = trimQuotes(field[1])
		}
	}
	return platform, version, nil
}

// Remove quotes of the source string(删除源字符的引号)
func trimQuotes(s string) string {
	if len(s) >= 2 {
		if s[0] == '"' && s[len(s)-1] == '"' {
			return s[1 : len(s)-1]
		}
	}
	return s
}
