package property

import "reflect"
import "dlib/dbus"

type BaseObserver struct {
	cbs []func()
}

func (o *BaseObserver) ConnectChanged(cb func()) {
	o.cbs = append(o.cbs, cb)
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
	core.ConnectChanged(func() {
		dbus.NotifyChange(p.obj, p.name)
	})
	return p
}

func (p WrapProperty) Set(v interface{}) {
	p.core.Set(v)
	dbus.NotifyChange(p.obj, p.name)
	p.Notify()
}

func (p WrapProperty) Get() interface{} {
	return p.core.Get()
}

func (p WrapProperty) GetType() reflect.Type {
	return p.core.GetType()
}
