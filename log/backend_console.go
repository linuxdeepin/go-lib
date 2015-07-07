/**
 * Copyright (b) 2013 ~ 2014 Deepin, Inc.
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
	"os"
	"pkg.deepin.io/lib/utils"
	"time"
)

const defaultDebugConsoleEnv = "DDE_DEBUG_CONSOLE"

var (
	// DebugConsoleEnv is the name of environment variable that used to control
	// the console backend print log in syslog format.
	DebugConsoleEnv = defaultDebugConsoleEnv
)

type backendConsole struct {
	name       string
	syslogMode bool
}

func newBackendConsole(name string) (b *backendConsole) {
	b = &backendConsole{}
	b.name = name
	if utils.IsEnvExists(DebugConsoleEnv) {
		b.syslogMode = true
	}
	return
}

func (b *backendConsole) log(level Priority, msg string) (err error) {
	formatMsg, err := b.formatMsg(level, msg)
	if err != nil {
		return
	}
	if b.syslogMode {
		fmt.Println(getSyslogPrefix(b.name), formatMsg)
	} else {
		fmt.Println(formatMsg)
	}
	return
}
func getSyslogPrefix(name string) (prefix string) {
	hostname, _ := os.Hostname()
	prefix = fmt.Sprintf("%s %s %s[%d]:", time.Now().Format("Jan 2 15:04:05"), hostname, name, os.Getpid())
	return
}

func (b *backendConsole) formatMsg(level Priority, msg string) (fmtMsg string, err error) {
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
	default:
		err = errUnknownLogLevel
		return
	}
	fmtMsg = levelStr + " " + msg
	return
}

func (b *backendConsole) close() (err error) {
	return
}
