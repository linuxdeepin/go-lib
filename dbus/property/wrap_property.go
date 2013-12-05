package property

import "reflect"
import "dlib/dbus"

type WrapProperty struct {
	obj  dbus.DBusObject
	name string
	core dbus.Property
}

func NewWrapProperty(obj dbus.DBusObject, propName string, core dbus.Property) WrapProperty {
	return WrapProperty{obj, propName, core}
}

func (p WrapProperty) Set(v interface{}) {
	p.core.Set(v)
	dbus.NotifyChange(p.obj, p.name)
}

func (p WrapProperty) Get() interface{} {
	return p.core.Get()
}

func (p WrapProperty) GetType() reflect.Type {
	return p.core.GetType()
}
