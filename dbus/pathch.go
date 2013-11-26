package dbus

import "log"

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

func isExitsInBus(con *Conn, obj DBusObject) bool {
	con.handlersLck.Lock()
	defer con.handlersLck.Unlock()
	info := obj.GetDBusInfo()
	if ifcs , ok := con.handlers[ObjectPath(info.ObjectPath)]; ok {
		return ifcs[info.Interface] == obj
	}
	return false
}

func NotifyChange(obj DBusObject, propName string) {
	value := getValueOf(obj).FieldByName(propName)
	if value.IsValid() {
		systemBusLck.Lock()
		defer systemBusLck.Unlock()
		sessionBusLck.Lock()
		defer sessionBusLck.Unlock()

		var err error
		if systemBus != nil && isExitsInBus(systemBus, obj) {
			inputs := make(map[string]Variant)
			inputs[propName] = MakeVariant(value.Interface())
			info := obj.GetDBusInfo()
			err = systemBus.Emit(ObjectPath(info.ObjectPath), "org.freedesktop.DBus.Properties.PropertiesChanged", info.Interface, inputs, make([]string, 0))
		} else if sessionBus != nil && isExitsInBus(sessionBus, obj) {
			inputs := make(map[string]Variant)
			inputs[propName] = MakeVariant(value.Interface())
			info := obj.GetDBusInfo()
			err = sessionBus.Emit(ObjectPath(info.ObjectPath), "org.freedesktop.DBus.Properties.PropertiesChanged", info.Interface, inputs, make([]string, 0))
		}
		if err != nil {
			log.Print(err)
		}
	} else {
		log.Printf(reflect.TypeOf(obj).String(), "hasn't the ", propName, "property")
	}
}
