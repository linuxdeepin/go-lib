/**
 * Copyright (c) 2013 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 Xu FaSheng
 *
 * Author:      Xu FaSheng <fasheng.xu@gmail.com>
 * Maintainer:  Xu FaSheng <fasheng.xu@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 **/

package log

import (
	"os"
	"path/filepath"
	"strings"
)

// TODO
var crashReporterArgs = []string{crashReporterExe, "--remove-config", "--config"}

// restartConfig stores data to be used by deepin-crash-reporter
type restartConfig struct {
	AppName          string
	RestartCommand   []string
	RestartEnv       map[string]string
	RestartDirectory string
	LogDetail        string
}

func newRestartConfig(logname string) *restartConfig {
	config := &restartConfig{}
	config.AppName = logname
	config.RestartCommand = os.Args
	config.RestartCommand[0], _ = filepath.Abs(os.Args[0])
	config.RestartDirectory, _ = os.Getwd()

	// setup envrionment variables
	config.RestartEnv = make(map[string]string)
	environs := os.Environ()
	for _, env := range environs {
		values := strings.SplitN(env, "=", 2)
		// values[0] is environment variable name, values[1] is the value
		if len(values) == 2 {
			config.RestartEnv[values[0]] = values[1]
		}
	}
	return config
}
