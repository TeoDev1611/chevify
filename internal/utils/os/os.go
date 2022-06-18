package os

import (
	"errors"
	"os"
	"runtime"
)

func System() (string, error) {
	switch runtime.GOOS {
	case "windows":
		return "windows", nil
	case "linux":
		return "linux", nil
	case "darwin":
		return "macos", nil
	default:
		return "", errors.New("Other Defaults Platforms Like: Linux, Windows, MacOs")
	}
}

func Exists(path string) (bool, error) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return false, nil
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		return true, nil
	}
	return false, errors.New("WTF Not exists and exists .-. report on github the error!")
}
