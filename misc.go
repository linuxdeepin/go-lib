package dlib

import "dlib/dbus"

/*
#include <glib.h>
#include <locale.h>
#include <stdlib.h>
#include <libintl.h>
void _run() {
	g_main_loop_run(g_main_loop_new(0, FALSE));
}
void _init_i18n() { setlocale(LC_ALL, ""); }
#cgo pkg-config: glib-2.0
*/
import "C"
import "unsafe"

func StartLoop() {
	C._run()
}

const (
	SystemBus  = 1
	SessionBus = 2
)

func Tr(id string) string {
	_id := C.CString(id)
	defer C.free(unsafe.Pointer(_id))
	return C.GoString(C.gettext(_id))
}

func DGettext(domain, id string) string {
	_id := C.CString(id)
	_d := C.CString(domain)
	defer C.free(unsafe.Pointer(_id))
	defer C.free(unsafe.Pointer(_d))
	return C.GoString(C.dgettext(_d, _id))
}

func Textdomain(domain string) {
	_domain := C.CString(domain)
	defer C.free(unsafe.Pointer(_domain))
	C.textdomain(_domain)
}

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
