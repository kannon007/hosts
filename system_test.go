package hosts

import (
	"os"
	"runtime"
)

func GetUserHomeDir() string {
	if runtime.GOOS == "windows" {
		return os.Getenv("USERPROFILE")
	}
	return os.Getenv("HOME")
}
