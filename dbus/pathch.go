package dbus

import "fmt"
import "log"

import "reflect"

type DBusInfo struct {
	Dest, ObjectPath, Interface string
}
type DBusObject interface {
	GetDBusInfo() DBusInfo
}

var (
	dbusObject          DBusObject
	dbusObjectInterface = reflect.TypeOf((*DBusObject)(nil)).Elem()
	introspectProxyType = reflect.TypeOf((*IntrospectProxy)(nil)).Elem()
	propertyType        = reflect.TypeOf((*Property)(nil)).Elem()
	dbusStructType      = reflect.TypeOf((*[]interface{})(nil)).Elem()
)

func isStructureMatched(structValue interface{}, dbusValue interface{}) bool {
	dValues, ok := dbusValue.([]interface{})
	if !ok {
		fmt.Println(dbusValue, "is Not an []interface{}")
		return false
	}

	t1 := reflect.TypeOf(structValue)
	if t1.Kind() == reflect.Ptr {
		t1 = t1.Elem()
	}

	if t1.Kind() != reflect.Struct {
		fmt.Println(t1.Kind(), "!=", reflect.Struct)
		return false
	}

	j := -1
	for i := 0; i < t1.NumField(); i++ {
		field := t1.Field(i)
		if isExportedStructField(field) {
			j++
			if j >= len(dValues) {
				fmt.Println("J:", j, "values:", len(dValues))
				return false
			}

			t := field.Type
			if t.Kind() == reflect.Ptr {
				t = t.Elem()
			}
			if reflect.TypeOf(dValues[j]).Kind() != t.Kind() {
				fmt.Println(reflect.TypeOf(dValues[j]).Kind(), "!=", t.Kind())
				return false
			}
		}
	}
	j++
	if j != len(dValues) {
		fmt.Println("exported Num:", j, "But dValues has NUm:", len(dValues))
		return false
	}
	return true
}

func isExportedStructField(field reflect.StructField) bool {
	if field.PkgPath == "" && field.Tag.Get("dbus") != "-" {
		return true
	} else {
		return false
	}
}

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
	if obj == nil {
		panic("detecConnByDBusObject must not feed nil")
	}
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
	if obj == nil {
		return
	}
	con := detectConnByDBusObject(obj)
	if con != nil {
		value := getValueOf(obj).FieldByName(propName)
		if value.IsValid() {
			if value.Type().Implements(propertyType) {
				v := value.MethodByName("GetValue").Interface().(func() interface{})()
				if v == nil {
					log.Println("dbus.NotifyChange", propName, "is an nil value! This shouldn't happen.")
				}
				value = reflect.ValueOf(v)
			}
			value = tryTranslateDBusObjectToObjectPath(con, value)
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
