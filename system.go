package hosts

import (
	"os"
	"path/filepath"
	"runtime"
)

func GetUserHomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}

func GetHostsPath() string {
	if isWindows() {
		//TODO 获取系统盘路径
		sysDir := getWinSystemDir()
		return filepath.Join(sysDir,"system32","drivers","etc","hosts")
	}else {
		return "/etc/hosts"
	}
}

func isWindows() bool {
	return runtime.GOOS == "windows";
}

func getWinSystemDir() string {
	dir := ""
	if isWindows() {
		dir = os.Getenv("windir")
	}

	return dir
}