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
	"log/syslog"
)

const defaultSyslogTagPrefix = ""

var (
	// SyslogTagPrefix define the prefix of syslog tag, default is
	// empty.
	SyslogTagPrefix = defaultSyslogTagPrefix
)

type backendSyslog struct {
	name   string
	writer *syslog.Writer
}

func newBackendSyslog(name string) (b *backendSyslog) {
	b = &backendSyslog{}
	b.name = name
	var err error
	b.writer, err = newSyslogWriter(name)
	if err != nil {
		std.Println("<info> syslog is not available:", err)
		return nil
	}
	return
}
func newSyslogWriter(name string) (l *syslog.Writer, err error) {
	tag := SyslogTagPrefix + name
	l, err = syslog.New(syslog.LOG_DAEMON, tag)
	return
}

func (b *backendSyslog) log(level Priority, msg string) (err error) {
	switch level {
	case LevelDebug:
		err = b.writer.Debug(msg)
	case LevelInfo:
		err = b.writer.Info(msg)
	case LevelWarning:
		err = b.writer.Warning(msg)
	case LevelError:
		err = b.writer.Err(msg)
	case LevelPanic:
		err = b.writer.Emerg(msg)
	case LevelFatal:
		err = b.writer.Emerg(msg)
	default:
		err = errUnknownLogLevel
	}
	return
}

func (b *backendSyslog) close() (err error) {
	err = b.writer.Close()
	return
}
