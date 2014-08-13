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
	golog "log"
)

var logapi *Logapi

type backendDeepinlog struct {
	id string
}

func newBackendDeepinlog(name string) (b *backendDeepinlog) {
	b = &backendDeepinlog{}
	err := initLogapi()
	if err != nil {
		golog.Println("initialize deepin log dbus interface failed:", err)
		return
	}
	b.id, err = logapi.NewLogger(name)
	if err != nil {
		golog.Println("create deepin log object failed:", err)
		return
	}
	return
}
func initLogapi() (err error) {
	if logapi == nil {
		logapi, err = newLogapi("/com/deepin/api/Logger")
	}
	return
}

// GetBackendConsole new and return a deepinlog back-end object.
func GetBackendDeepinlog(name string) Backend {
	return newBackendDeepinlog(name)
}

func (b *backendDeepinlog) log(name string, level Priority, msg string) (err error) {
	if logapi == nil {
		err = fmt.Errorf("deepin log dbus interface could not access")
		return
	}
	switch level {
	case LevelDebug:
		err = logapi.Debug(b.id, name, msg)
	case LevelInfo:
		err = logapi.Info(b.id, name, msg)
	case LevelWarning:
		err = logapi.Warning(b.id, name, msg)
	case LevelError:
		err = logapi.Error(b.id, name, msg)
	case LevelPanic:
		err = logapi.Error(b.id, name, msg)
	case LevelFatal:
		err = logapi.Fatal(b.id, name, msg)
	}
	return
}
