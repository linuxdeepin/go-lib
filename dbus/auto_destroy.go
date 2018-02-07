/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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
import "sync"

var quitTimer *time.Timer

var hasNewMessage, markHasNewMessage = func() (func() bool, func(bool)) {
	__v := false
	var mux sync.RWMutex
	return func() bool {
			mux.RLock()
			defer mux.RUnlock()
			return __v
		}, func(v bool) {
			mux.Lock()
			__v = v
			mux.Unlock()
		}
}()

func SetAutoDestroyHandler(d time.Duration, cb func() bool) {
	if quitTimer != nil {
		quitTimer.Stop()
	}
	if d == 0 {
		return
	}
	quitTimer = time.AfterFunc(d, func() {
		if !hasNewMessage() {
			if cb == nil || cb() {
				quit(nil)
				return
			}
		}
		markHasNewMessage(false)
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
