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
