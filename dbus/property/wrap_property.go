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

type BaseObserver struct {
	cbs []func()
}

func (o *BaseObserver) ConnectChanged(cb func()) {
	o.cbs = append(o.cbs, cb)
}

func (o *BaseObserver) Reset() {
	o.cbs = nil
}

func (o *BaseObserver) Notify() {
	for _, c := range o.cbs {
		c()
	}
}

type WrapProperty struct {
	BaseObserver
	obj  dbus.DBusObject
	name string
	core dbus.Property
}

func NewWrapProperty(obj dbus.DBusObject, propName string, core dbus.Property) *WrapProperty {
	p := &WrapProperty{BaseObserver{}, obj, propName, core}
	p.core = core
	core.ConnectChanged(func() {
		dbus.NotifyChange(p.obj, p.name)
	})
	return p
}

func (p WrapProperty) SetValue(v interface{}) {
	p.core.SetValue(v)
	dbus.NotifyChange(p.obj, p.name)
	p.Notify()
}

func (p WrapProperty) GetValue() interface{} {
	return p.core.GetValue()
}

func (p WrapProperty) GetType() reflect.Type {
	return p.core.GetType()
}
