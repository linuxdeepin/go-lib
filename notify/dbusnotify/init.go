// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusnotify

import (
	"fmt"
	"sync"

	"github.com/godbus/dbus/v5"
)

var __conn *dbus.Conn = nil
var __connLock sync.Mutex

var __ruleCounter map[string]int = nil
var __ruleCounterLock sync.Mutex

func getBus() *dbus.Conn {
	__connLock.Lock()
	defer __connLock.Unlock()
	if __conn == nil {
		var err error
		__conn, err = dbus.SessionBus()
		if err != nil {
			panic(err)
		}
	}
	return __conn
}

func getRuleCounter() map[string]int {
	__ruleCounterLock.Lock()
	defer __ruleCounterLock.Unlock()
	if __ruleCounter == nil {
		__ruleCounter = make(map[string]int)
	}
	return __ruleCounter
}

func dbusCall(method string, flags dbus.Flags, args ...interface{}) (err error) {
	err = getBus().BusObject().Call(method, flags, args...).Err
	if err != nil {
		fmt.Println(err)
	}
	return
}

func dbusAddMatch(rule string) (err error) {
	ruleCounter := getRuleCounter()

	__ruleCounterLock.Lock()
	defer __ruleCounterLock.Unlock()
	if _, ok := ruleCounter[rule]; !ok {
		err = dbusCall("org.freedesktop.DBus.AddMatch", 0, rule)
	}
	ruleCounter[rule]++
	return
}

func dbusRemoveMatch(rule string) (err error) {
	ruleCounter := getRuleCounter()

	__ruleCounterLock.Lock()
	defer __ruleCounterLock.Unlock()
	if _, ok := ruleCounter[rule]; !ok {
		return
	}
	ruleCounter[rule]--
	if ruleCounter[rule] == 0 {
		delete(ruleCounter, rule)
		err = dbusCall("org.freedesktop.DBus.RemoveMatch", 0, rule)
	}
	return
}
