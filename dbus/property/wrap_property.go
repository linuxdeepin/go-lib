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
