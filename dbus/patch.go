/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
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
type DMessage struct {
	m *Message
	c *Conn
}

type Property interface {
	GetValue() interface{}
	SetValue(interface{})
	ConnectChanged(func())
	GetType() reflect.Type
}

func (msg DMessage) GetSender() string {
	sender := msg.m.Headers[FieldSender].value.(string)
	return sender
}

func (msg DMessage) GetSenderPID() (r uint32) {
	msg.c.BusObject().Call("org.freedesktop.DBus.GetConnectionUnixProcessID", 0, msg.GetSender()).Store(&r)
	return
}
func (msg DMessage) GetSenderUID() (r uint32) {
	msg.c.BusObject().Call("org.freedesktop.DBus.GetConnectionUnixUser", 0, msg.GetSender()).Store(&r)
	return
}

var (
	dbusObject          DBusObject
	dbusObjectInterface = reflect.TypeOf((*DBusObject)(nil)).Elem()
	introspectProxyType = reflect.TypeOf((*IntrospectProxy)(nil))
	dbusStructType      = reflect.TypeOf((*[]interface{})(nil)).Elem()
	dbusMessageType     = reflect.TypeOf((*DMessage)(nil)).Elem()
)

const (
	UserMethodFlagNone = 1 << iota
	UserMethodFlagWillThrowError
	UserMethodFlagNeedDMessage
)

func detectExportMethodFlags(t reflect.Type) int {
	flags := UserMethodFlagNone
	if n := t.NumOut(); n > 0 {
		if t.Out(n - 1).Implements(goErrorType) {
			flags = flags | UserMethodFlagWillThrowError
		}
	}

	if n := t.NumIn(); n > 0 {
		if t.In(0).ConvertibleTo(dbusMessageType) {
			flags = flags | UserMethodFlagNeedDMessage
		}
	}
	return flags
}

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
			if p, ok := value.Interface().(Property); ok {
				v := p.GetValue()
				if v == nil {
					log.Printf("The value of %q is nil. This shouldn't happen.", propName)
					return
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

func ReleaseName(obj DBusObject) {
	if obj == nil {
		return
	}
	con := detectConnByDBusObject(obj)
	if con != nil {
		con.busObj.Call("org.freedesktop.DBus.ReleaseName", 0, obj.GetDBusInfo().Dest)
	}
}

func Emit(obj DBusObject, name string, ins ...interface{}) error {
	c := detectConnByDBusObject(obj)
	if c == nil {
		return fmt.Errorf("%s is not installed on any bus", obj.GetDBusInfo().Dest)
	}
	v := getValueOf(obj)
	fn := v.FieldByName(name)
	if !fn.IsValid() {
		panic(fmt.Errorf("Can't find method of %s in %s", name, obj.GetDBusInfo().Interface))
	}

	fnType := getTypeOf(fn.Interface())
	if fnType.NumOut() != 0 || fnType.NumIn() != len(ins) {
		panic(fmt.Errorf("Invalid signal type (%d != %d)", fnType.NumIn(), len(ins)))
	}
	for i, in := range ins {
		if fnType.In(i) != reflect.TypeOf(in) {
			panic(fmt.Errorf("Invalid signal type %d (%v != %v)", i, fnType.In(i), reflect.TypeOf(in)))
		}
	}

	info := obj.GetDBusInfo()
	return c.Emit(ObjectPath(info.ObjectPath), info.Interface+"."+name, ins...)
}
