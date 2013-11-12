package property

import "dlib"
import "reflect"
import "dlib/dbus"

type GSettingsProperty struct {
	core      *dlib.Settings
	key       string
	valueType reflect.Type
}

func NewGSettingsProperty(s *dlib.Settings, key string, t interface{}) *GSettingsProperty {
	switch reflect.TypeOf(t).Kind() {
	case reflect.Bool:
	case reflect.Float32, reflect.Float64:
	case reflect.String:
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint32, reflect.Uint64:
	case reflect.Array:
		panic("Don't support array of string use v[:] to pass an slice of string!")
	case reflect.Slice:
		if reflect.TypeOf(t).Elem().Kind() != reflect.String {
			panic("GSetting only support slice of string!")
		}
	default:
		panic("Don't support array of string use v[:] to pass an slice of string!")
	}
	return &GSettingsProperty{s, key, reflect.TypeOf(t)}
}

func NewGSettingsPropertyFull(s *dlib.Settings, key string, t interface{}, con *dbus.Conn, path, ifc, propName string) *GSettingsProperty {
	prop := NewGSettingsProperty(s, key, t)
	prop.core.Connect("changed::"+key, func(s *dlib.Settings, key string) {
		inputs := make(map[string]dbus.Variant)
		inputs[propName] = dbus.MakeVariant(prop.Get())
	})
	return prop
}

func (p GSettingsProperty) Set(v interface{}) {
	if reflect.TypeOf(v) != p.valueType {
		panic("This property need type of " + p.valueType.String())
	}
	switch p.valueType.Kind() {
	case reflect.Bool:
		p.core.SetBoolean(p.key, v.(bool))
		return
	case reflect.Float32, reflect.Float64:
		p.core.SetDouble(p.key, reflect.ValueOf(v).Float())
		return
	case reflect.String:
		p.core.SetString(p.key, v.(string))
		return
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint32, reflect.Uint64:
		p.core.SetInt(p.key, int(reflect.ValueOf(v).Int()))
		return
	case reflect.Slice:
		p.core.SetStrv(p.key, v.([]string))
		return

	}
	panic("Didn't support type " + reflect.TypeOf(v).String())
}

func (p GSettingsProperty) Get() interface{} {
	switch p.valueType.Kind() {
	case reflect.Bool:
		return p.core.GetBoolean(p.key)
	case reflect.Float32, reflect.Float64:
		return p.core.GetDouble(p.key)
	case reflect.String:
		return p.core.GetString(p.key)
	case reflect.Int, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint32, reflect.Uint64:
		return p.core.GetInt(p.key)
	case reflect.Slice:
		return p.core.GetStrv(p.key)
	}
	panic("Didn't support type!")
}

func (p GSettingsProperty) GetType() reflect.Type {
	return p.valueType
}
