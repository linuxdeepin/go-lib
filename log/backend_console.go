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
	"fmt"
	"pkg.linuxdeepin.com/lib/utils"
	"strings"
)

const defaultDebugConsoleEnv = "DDE_DEBUG_CONSOLE"

var (
	// DebugConsoleEnv is the name of environment variable that used to control
	// the console backend print log in syslog format.
	DebugConsoleEnv = defaultDebugConsoleEnv

	backendConsole = newConsole()
)

type console struct {
	syslogMode bool
}

func newConsole() (c *console) {
	c = &console{}
	if utils.IsEnvExists(DebugConsoleEnv) {
		c.syslogMode = true
	}
	return
}

// GetBackendConsole return the only console back-end object.
func GetBackendConsole() Backend {
	return backendConsole
}

func (c *console) log(level Priority, msg string) (err error) {
	fmt.Println(c.formatMsg(level, msg))
	return
}

// TODO
func (c *console) formatMsg(level Priority, msg string) (fmtMsg string) {
	var levelStr string
	switch level {
	case LevelDebug:
		levelStr = "<debug>"
	case LevelInfo:
		levelStr = "<info>"
	case LevelWarning:
		levelStr = "<warning>"
	case LevelError:
		levelStr = "<error>"
	case LevelPanic:
		levelStr = "<error>"
	case LevelFatal:
		levelStr = "<error>"
	}
	fmtMsg = levelStr + " " + msg
	// fmtMsg = strings.Replace(fmtMsg, "\n", "\n"+levelStr+" ", -1) // format multi-lines message
	strings.Replace(fmtMsg, "\n", "\n"+levelStr+" ", -1) // format multi-lines message
	return
}
