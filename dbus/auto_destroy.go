package dbus

import "fmt"
import "time"

var _HasNewMessage = false
var quitTimer *time.Timer

func SetAutoDestroyHandler(d time.Duration, cb func() bool) {
	if quitTimer != nil {
		quitTimer.Stop()
	}
	if cb == nil || d == 0 {
		fmt.Println("Clean AutoDestroyHandle")
		return
	}
	quitTimer = time.AfterFunc(d, func() {
		if !_HasNewMessage {
			if cb() {
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
