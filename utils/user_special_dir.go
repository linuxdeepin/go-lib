// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package utils

import (
	"os/user"
	"path"
)

func GetHomeDir() string {
	info, err := user.Current()
	if err != nil {
		return ""
	}

	return info.HomeDir
}

func GetConfigDir() string {
	if homeDir := GetHomeDir(); len(homeDir) > 0 {
		return path.Join(homeDir, ".config")
	}
	return ""
}

func GetCacheDir() string {
	if homeDir := GetHomeDir(); len(homeDir) > 0 {
		return path.Join(homeDir, ".cache")
	}
	return ""
}
