package property

import "dlib/gio-2.0"
import "reflect"
import "dlib/dbus"

type GSettingsProperty struct {
	*BaseObserver
	valueType reflect.Type
	getFn     func() interface{}
	setFn     func(interface{})
	core      dbus.DBusObject
	propName  string
}

func NewGSettingsProperty(obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *GSettingsProperty {
	prop := &GSettingsProperty{}
	prop.core = obj
	prop.BaseObserver = &BaseObserver{}
	prop.propName = propName
	switch s.GetValue(keyName).GetTypeString() {
	case "b":
		prop.valueType = reflect.TypeOf(false)
		prop.getFn = func() interface{} {
			return s.GetBoolean(keyName)
		}
		prop.setFn = func(v interface{}) {
			s.SetBoolean(keyName, v.(bool))
		}
	case "i":
		prop.valueType = reflect.TypeOf(int32(0))
		prop.getFn = func() interface{} {
			return int32(s.GetInt(keyName))
		}
		prop.setFn = func(v interface{}) {
			s.SetInt(keyName, int(reflect.ValueOf(v).Int()))
		}
	case "u":
		prop.valueType = reflect.TypeOf(uint32(0))
		prop.getFn = func() interface{} {
			return uint32(s.GetUint(keyName))
		}
		prop.setFn = func(v interface{}) {
			s.SetUint(keyName, int(reflect.ValueOf(v).Uint()))
		}
	case "d":
		prop.valueType = reflect.TypeOf(float64(0))
		prop.getFn = func() interface{} {
			return s.GetDouble(keyName)
		}
		prop.setFn = func(v interface{}) {
			s.SetDouble(keyName, reflect.ValueOf(v).Float())
		}
	case "s":
		prop.valueType = reflect.TypeOf("")
		prop.getFn = func() interface{} {
			return s.GetString(keyName)
		}
		prop.setFn = func(v interface{}) {
			s.SetString(keyName, reflect.ValueOf(v).String())
		}
	case "as":
		prop.valueType = reflect.TypeOf([]string{})
		prop.getFn = func() interface{} {
			return s.GetStrv(keyName)
		}
		prop.setFn = func(v interface{}) {
			s.SetStrv(keyName, v.([]string))
		}
	default:
		panic("GSettingsProperty didn't support gsettings key " + keyName)
	}
	s.Connect("changed::"+keyName, func(s *gio.Settings, name string) {
		dbus.NotifyChange(prop.core, prop.propName)
	})
	return prop
}

func (p GSettingsProperty) Set(v interface{}) {
	if v != p.getFn() {
		p.setFn(v)
		dbus.NotifyChange(p.core, p.propName)
	}
}

func (p GSettingsProperty) Get() interface{} {
	return p.getFn()
}

func (p GSettingsProperty) GetType() reflect.Type {
	return p.valueType
}
