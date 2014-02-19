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
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package logger

import (
	"dlib/dbus"
	"fmt"
	"log"
	"runtime"
)

const (
	_DEST = "com.deepin.api.Logger"
	_PATH = "/com/deepin/api/Logger"
)

var __conn *dbus.Conn = nil

func getBus() *dbus.Conn {
	if __conn == nil {
		var err error
		__conn, err = dbus.SystemBus()
		if err != nil {
			panic(err)
		}
	}
	return __conn
}

// TODO
func getBusObject() {
}

// TODO remove
func NewImage(path dbus.ObjectPath) (*Image, error) {
	if !path.IsValid() {
		return nil, errors.New("The path of '" + string(path) + "' is invalid.")
	}

	core := getBus().Object("com.deepin.api.Image", path)
	var v string
	core.Call("org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&v)
	if strings.Index(v, "com.deepin.api.Image") == -1 {
		return nil, errors.New("'" + string(path) + "' hasn't interface 'com.deepin.api.Image'.")
	}

	obj := &Image{Path: path, core: core}

	return obj, nil
}

func Println(v ...interface{}) {
	r := fmt.Sprintln(v...)
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d :%s", file, line, r)
}

func Printf(format string, v ...interface{}) {
	r := fmt.Sprintf(format, v...)
	_, file, line, _ := runtime.Caller(1)
	log.Printf("%s:%d :%s", file, line, r)
}

func Assert(exp bool, v ...interface{}) {
	if exp == false {
		panic(fmt.Sprintln(v...))
	}
}
func AssertNotReached() {
	panic("Shouldn't reached")
}
