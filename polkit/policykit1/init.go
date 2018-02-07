/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
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

package policykit1

import "pkg.deepin.io/lib/dbus"
import "fmt"
import "sync"

var __conn *dbus.Conn = nil
var __connLock sync.Mutex

var __ruleCounter map[string]int = nil
var __ruleCounterLock sync.Mutex

func getBus() *dbus.Conn {
	__connLock.Lock()
	defer __connLock.Unlock()
	if __conn == nil {
		var err error
		__conn, err = dbus.SystemBus()
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
