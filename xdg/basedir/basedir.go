// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package basedir

// XDG Base Directory Specification

import (
	"errors"
	"os"
	"os/user"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func GetUserHomeDir() string {
	dir := os.Getenv("HOME")
	if dir != "" {
		return filepath.Clean(dir)
	}
	user, err := user.Current()
	if err != nil {
		return ""
	}
	return filepath.Clean(user.HomeDir)
}

func GetUserDataDir() string {
	// default $HOME/.local/share
	defaultDir := filepath.Join(GetUserHomeDir(), ".local/share")
	return getUserDir("XDG_DATA_HOME", defaultDir)
}

func GetSystemDataDirs() []string {
	defaultDirs := []string{"/usr/local/share", "/usr/share"}
	return getSystemDirs("XDG_DATA_DIRS", defaultDirs)
}

func GetUserConfigDir() string {
	// default $HOME/.config
	defaultDir := filepath.Join(GetUserHomeDir(), ".config")
	return getUserDir("XDG_CONFIG_HOME", defaultDir)
}

func GetSystemConfigDirs() []string {
	defaultDirs := []string{"/etc/xdg"}
	return getSystemDirs("XDG_CONFIG_DIRS", defaultDirs)
}

func GetUserCacheDir() string {
	defaultDir := filepath.Join(GetUserHomeDir(), ".cache")
	return getUserDir("XDG_CACHE_HOME", defaultDir)
}

func GetUserRuntimeDir(strict bool) (string, error) {
	dir := os.Getenv("XDG_RUNTIME_DIR")
	if dir == "" || !filepath.IsAbs(dir) {
		if strict {
			return "", errors.New("env $XDG_RUNTIME_DIR is invalid")
		}

		// fallback runtime user dir
		fallback := "/tmp/goxdg-runtime-dir-fallback-" + strconv.Itoa(os.Getuid())
		create := false

		fi, err := os.Lstat(fallback)
		if err != nil {
			if os.IsNotExist(err) {
				create = true
			} else {
				return "", err
			}
			// Lstat ok
		} else {
			// The fallback must be a directory
			if !fi.IsDir() {
				if err := os.Remove(fallback); err != nil {
					return "", err
				}
				create = true

				// Must be owned by the user and not accessible by anyone else
			} else if getFileInfoUid(fi) != os.Getuid() ||
				(fi.Mode()&os.ModePerm) != 0700 {
				if err := os.RemoveAll(fallback); err != nil {
					return "", err
				}
				create = true
			}
		}

		if create {
			_ = os.Mkdir(fallback, 0700)
		}
		return fallback, nil
	}

	return filepath.Clean(dir), nil
}

func getUserDir(envVarName string, defaultDir string) string {
	dir := os.Getenv(envVarName)
	if dir == "" {
		return defaultDir
	}
	if !filepath.IsAbs(dir) {
		return defaultDir
	}
	return filepath.Clean(dir)
}

func getSystemDirs(envVarName string, defaultDirs []string) []string {
	dirsEnv := os.Getenv(envVarName)
	if dirsEnv == "" {
		return defaultDirs
	}
	dirs := filterNotAbs(strings.Split(dirsEnv, ":"))
	if len(dirs) == 0 {
		return defaultDirs
	}
	return dirs
}

// filter not absolute path and clean path
func filterNotAbs(slice []string) []string {
	result := make([]string, 0, len(slice))
	for _, path := range slice {
		if !filepath.IsAbs(path) {
			continue
		}
		result = append(result, filepath.Clean(path))
	}
	return result
}

func getFileInfoUid(fi os.FileInfo) int {
	return int(fi.Sys().(*syscall.Stat_t).Uid)
}
