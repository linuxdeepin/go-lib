/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package lib

import "pkg.deepin.io/lib/dbus"

const (
	SystemBus  = 1
	SessionBus = 2
)

func UniqueOnSession(name string) bool {
	con, err := dbus.SessionBus()
	if err != nil {
		return false
	}
	return uniqueOnAny(con, name)
}
func UniqueOnSystem(name string) bool {
	con, err := dbus.SystemBus()
	if err != nil {
		return false
	}
	return uniqueOnAny(con, name)
}

func uniqueOnAny(bus *dbus.Conn, name string) bool {
	reply, err := bus.RequestName(name, dbus.NameFlagDoNotQueue)
	if err != nil || reply != dbus.RequestNameReplyPrimaryOwner {
		return false
	}
	return true
}
