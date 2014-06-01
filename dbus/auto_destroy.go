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
