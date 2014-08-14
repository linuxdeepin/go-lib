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
)

var logapi *Logapi

type backendDeepinlog struct {
	name string
	id   string
}

func newBackendDeepinlog(name string) (b *backendDeepinlog) {
	b = &backendDeepinlog{}
	b.name = name
	err := initLogapi()
	if err != nil {
		gologPrintln("initialize deepin log dbus interface failed:", err)
		return nil
	}
	b.id, err = logapi.NewLogger(name)
	if err != nil {
		gologPrintln("create deepin log object failed:", err)
		return nil
	}
	return
}
func initLogapi() (err error) {
	if logapi == nil {
		logapi, err = newLogapi("/com/deepin/api/Logger")
	}
	return
}

func (b *backendDeepinlog) log(level Priority, msg string) (err error) {
	if logapi == nil {
		err = fmt.Errorf("deepin log dbus interface could not access")
		return
	}
	switch level {
	case LevelDebug:
		err = logapi.Debug(b.id, b.name, msg)
	case LevelInfo:
		err = logapi.Info(b.id, b.name, msg)
	case LevelWarning:
		err = logapi.Warning(b.id, b.name, msg)
	case LevelError:
		err = logapi.Error(b.id, b.name, msg)
	case LevelPanic:
		err = logapi.Error(b.id, b.name, msg)
	case LevelFatal:
		err = logapi.Fatal(b.id, b.name, msg)
	default:
		err = errUnknownLogLevel
		return
	}
	return
}

func (b *backendDeepinlog) close() (err error) {
	return
}
