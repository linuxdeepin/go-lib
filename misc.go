package dlib

import "dlib/dbus"

/*
#include <glib.h>
#include <locale.h>
void _run() {
	g_main_loop_run(g_main_loop_new(0, FALSE));
}
void _init_i18n() { setlocale(LC_ALL, ""); }
#cgo pkg-config: glib-2.0
*/
import "C"

func StartLoop() {
	C._run()
}

const (
	SystemBus  = 1
	SessionBus = 2
)

func InitI18n() {
	C._init_i18n()
}

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
