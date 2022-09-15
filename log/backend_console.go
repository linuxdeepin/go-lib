// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package log

import (
	"fmt"
	"os"
	"github.com/linuxdeepin/go-lib/utils"
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
