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

func NotifyChange(con *Conn, obj DBusObject, propName string) {
	value := getValueOf(obj).FieldByName(propName)
	if value.IsValid() {
		inputs := make(map[string]Variant)
		inputs[propName] = MakeVariant(value.Interface())
		info := obj.GetDBusInfo()
		e := con.Emit(ObjectPath(info.ObjectPath), "org.freedesktop.DBus.Properties.PropertiesChanged", info.Interface, inputs, make([]string, 0))
		if e != nil {
			log.Print(e)
		}
	} else {
		log.Printf(reflect.TypeOf(obj).String(), "hasn't the ", propName, "property")
	}
}
