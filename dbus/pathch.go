package dbus

import "reflect"

type DBusInfo struct {
	Dest, ObjectPath, Interface string
}
type DBusObject interface {
	GetDBusInfo() DBusInfo
}

var dbusObject DBusObject
var dbusObjectInterface = reflect.TypeOf((*DBusObject)(nil)).Elem()
var introspectProxyType = reflect.TypeOf((*IntrospectProxy)(nil)).Elem()
var propertyType = reflect.TypeOf((*Property)(nil)).Elem()
