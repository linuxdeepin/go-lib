/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
