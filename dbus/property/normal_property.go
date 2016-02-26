/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package property

import "reflect"
import "pkg.deepin.io/lib/dbus"
import "log"

type NormalProperty struct {
	value     interface{}
	valueType reflect.Type
	notify    func()
}

func NewNormalProperty(con *dbus.Conn, path, ifc string, propName string, v interface{}) *NormalProperty {
	prop := NormalProperty{v, reflect.TypeOf(v), nil}
	prop.notify = func() {
		inputs := make(map[string]dbus.Variant)
		inputs[propName] = dbus.MakeVariant(prop.value)
		e := con.Emit(dbus.ObjectPath(path), "org.freedesktop.DBus.Properties.PropertiesChanged", ifc, inputs, make([]string, 0))
		if e != nil {
			log.Print(e)
		}
	}
	return &prop
}

func (p NormalProperty) SetValue(v interface{}) {
	if reflect.TypeOf(v) != p.valueType {
		panic("This property need type of " + p.valueType.String())
	}
	p.value = v
	p.notify()
}

func (p NormalProperty) GetValue() interface{} {
	return p.value
}

func (p NormalProperty) GetType() reflect.Type {
	return p.valueType
}
