/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package dbus

import "time"

var _HasNewMessage = false
var quitTimer *time.Timer

func SetAutoDestroyHandler(d time.Duration, cb func() bool) {
	if quitTimer != nil {
		quitTimer.Stop()
	}
	if d == 0 {
		return
	}
	quitTimer = time.AfterFunc(d, func() {
		if !_HasNewMessage {
			if cb == nil || cb() {
				quit(nil)
				return
			}
		}
		_HasNewMessage = false
		SetAutoDestroyHandler(d, cb)
	})
}

var _quit = make(chan error)

func Wait() error {
	return <-_quit
}

func quit(err error) {
	_quit <- err
}
