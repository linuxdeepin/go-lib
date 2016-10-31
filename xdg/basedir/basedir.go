package basedir

import (
	"os"
	"path/filepath"
	"strings"
)

const (
	ApplicationsDir = "applications"
)

func GetUserDataDir() string {
	userDir := os.Getenv("XDG_DATA_HOME")
	if userDir == "" {
		// $HOME/.local/share
		userDir = filepath.Join(os.Getenv("HOME"), ".local/share")
	}
	return userDir
}

func GetSystemDataDirs() []string {
	dirs := os.Getenv("XDG_DATA_DIRS")
	if dirs == "" {
		dirs = "/usr/local/share:/usr/share"
	}
	systemDirs := strings.Split(dirs, ":")
	return systemDirs
}

func GetUserConfigDir() string {
	dir := os.Getenv("XDG_CONFIG_HOME")
	if dir == "" {
		dir = filepath.Join(os.Getenv("HOME"), ".config")
	}
	return dir
}

func GetSystemConfigDir() []string {
	dirs := os.Getenv("XDG_CONFIG_DIRS")
	if dirs == "" {
		dirs = "/etc/xdg"
	}
	return strings.Split(dirs, ":")
}
