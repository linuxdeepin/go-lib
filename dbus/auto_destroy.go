/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
