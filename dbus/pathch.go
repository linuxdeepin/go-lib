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
	if ifcs, ok := con.handlers[ObjectPath(info.ObjectPath)]; ok {
		return ifcs[info.Interface] == obj
	}
	return false
}

func detectConnByDBusObject(obj DBusObject) *Conn {
	systemBusLck.Lock()
	defer systemBusLck.Unlock()
	sessionBusLck.Lock()
	defer sessionBusLck.Unlock()

	if systemBus != nil && isExitsInBus(systemBus, obj) {
		return systemBus
	} else if sessionBus != nil && isExitsInBus(sessionBus, obj) {
		return sessionBus
	}
	return nil
}

func NotifyChange(obj DBusObject, propName string) {
	con := detectConnByDBusObject(obj)
	if con != nil {
		value := getValueOf(obj).FieldByName(propName)
		value = tryTranslateDBusObjectToObjectPath(con, value)
		if value.IsValid() {
			inputs := make(map[string]Variant)
			inputs[propName] = MakeVariant(value.Interface())
			info := obj.GetDBusInfo()
			err := con.Emit(ObjectPath(info.ObjectPath), "org.freedesktop.DBus.Properties.PropertiesChanged", info.Interface, inputs, make([]string, 0))
			if err != nil {
				log.Println("NotifyChange send message error:", err)
			}
		} else {
			log.Println(reflect.TypeOf(obj).String(), "hasn't the ", propName, "property")
		}
	}
}
