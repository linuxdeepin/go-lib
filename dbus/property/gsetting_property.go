package property

import "pkg.deepin.io/lib/gio-2.0"
import "pkg.deepin.io/lib/dbus"
import "reflect"
import "fmt"
import "unicode"

type _GSettingsProperty struct {
	*BaseObserver
	valueType reflect.Type
	getFn     func() interface{}
	setFn     func(interface{})
	core      dbus.DBusObject
	propName  string
}

var (
	boolType    = reflect.TypeOf(false)
	int32Type   = reflect.TypeOf(int32(0))
	uint32Type  = reflect.TypeOf(uint32(0))
	float64Type = reflect.TypeOf(float64(0))
	stringType  = reflect.TypeOf("")
	strvType    = reflect.TypeOf([]string{})
)

type GSettingsBoolProperty struct{ *_GSettingsProperty }
type GSettingsIntProperty struct{ *_GSettingsProperty }
type GSettingsEnumProperty struct{ *_GSettingsProperty }
type GSettingsUintProperty struct{ *_GSettingsProperty }
type GSettingsFloatProperty struct{ *_GSettingsProperty }
type GSettingsStringProperty struct{ *_GSettingsProperty }
type GSettingsStrvProperty struct{ *_GSettingsProperty }

func NewGSettingsBoolProperty(obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *GSettingsBoolProperty {
	return &GSettingsBoolProperty{newGSettingsProperty("b", obj, propName, s, keyName)}
}
func (prop *GSettingsBoolProperty) Get() bool {
	return prop.GetValue().(bool)
}
func (prop *GSettingsBoolProperty) Set(v bool) {
	prop.SetValue(v)
}

func NewGSettingsIntProperty(obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *GSettingsIntProperty {
	return &GSettingsIntProperty{newGSettingsProperty("i", obj, propName, s, keyName)}
}
func (prop *GSettingsIntProperty) Get() int32 {
	return prop.GetValue().(int32)
}
func (prop *GSettingsIntProperty) Set(v int32) {
	prop.SetValue(v)
}

func NewGSettingsEnumProperty(obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *GSettingsEnumProperty {
	return &GSettingsEnumProperty{newGSettingsProperty("e", obj, propName, s, keyName)}
}
func (prop *GSettingsEnumProperty) Get() int32 {
	return prop.GetValue().(int32)
}
func (prop *GSettingsEnumProperty) Set(v int32) {
	prop.SetValue(v)
}

func NewGSettingsUintProperty(obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *GSettingsUintProperty {
	return &GSettingsUintProperty{newGSettingsProperty("u", obj, propName, s, keyName)}
}
func (prop *GSettingsUintProperty) Get() uint32 {
	return prop.GetValue().(uint32)
}
func (prop *GSettingsUintProperty) Set(v uint32) {
	prop.SetValue(v)
}

func NewGSettingsFloatProperty(obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *GSettingsFloatProperty {
	return &GSettingsFloatProperty{newGSettingsProperty("d", obj, propName, s, keyName)}
}
func (prop *GSettingsFloatProperty) Get() float64 {
	return prop.GetValue().(float64)
}
func (prop *GSettingsFloatProperty) Set(v float64) {
	prop.SetValue(v)
}

func NewGSettingsStringProperty(obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *GSettingsStringProperty {
	return &GSettingsStringProperty{newGSettingsProperty("s", obj, propName, s, keyName)}
}
func (prop *GSettingsStringProperty) Get() string {
	return prop.GetValue().(string)
}
func (prop *GSettingsStringProperty) Set(v string) {
	prop.SetValue(v)
}

func NewGSettingsStrvProperty(obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *GSettingsStrvProperty {
	return &GSettingsStrvProperty{newGSettingsProperty("as", obj, propName, s, keyName)}
}
func (prop *GSettingsStrvProperty) Get() []string {
	return prop.GetValue().([]string)
}
func (prop *GSettingsStrvProperty) Set(v []string) {
	prop.SetValue(v)
}

func newGSettingsProperty(sig string, obj dbus.DBusObject, propName string, s *gio.Settings, keyName string) *_GSettingsProperty {
	if !unicode.IsUpper(rune(propName[0])) {
		panic(fmt.Sprintf("Can't bind an unexport field(%s) to %s", propName, keyName))
	}
	real_type := s.GetValue(keyName).GetTypeString()
	if real_type == "s" && sig == "e" {
		// note: "e" is not matched with glib type system
		real_type = "e"
	}
	if real_type != sig {
		var correct_method string
		switch real_type {
		case "b":
			correct_method = "NewGSettingsBoolProperty"
		case "i":
			correct_method = "NewGSettingsIntProperty"
		case "u":
			correct_method = "NewGSettingsUintProperty"
		case "d":
			correct_method = "NewGSettingsFloatProperty"
		case "s":
			correct_method = "NewGSettingsStringProperty"
		case "as":
			correct_method = "NewGSettingsStrvProperty"
		default:
			panic("GSettingsProperty didn't type " + sig)
		}
		panic("Type signal " + sig + " didn't match " + keyName + "'s type(" + real_type + ")" + ", please use the method " + correct_method)
	}

	prop := &_GSettingsProperty{}
	prop.core = obj
	prop.BaseObserver = &BaseObserver{}
	prop.propName = propName
	switch real_type {
	case "b":
		prop.valueType = boolType
		prop.getFn = func() interface{} {
			return s.GetBoolean(keyName)
		}
		prop.setFn = func(v interface{}) {
			s.SetBoolean(keyName, v.(bool))
		}
	case "i":
		prop.valueType = int32Type
		prop.getFn = func() interface{} {
			return int32(s.GetInt(keyName))
		}
		prop.setFn = func(v interface{}) {
			s.SetInt(keyName, int32(reflect.ValueOf(v).Int()))
		}
	case "e":
		prop.valueType = int32Type
		prop.getFn = func() interface{} {
			return int32(s.GetEnum(keyName))
		}
		prop.setFn = func(v interface{}) {
			s.SetEnum(keyName, int32(reflect.ValueOf(v).Int()))
		}
	case "u":
		prop.valueType = uint32Type
		prop.getFn = func() interface{} {
			return uint32(s.GetUint(keyName))
		}
		prop.setFn = func(v interface{}) {
			s.SetUint(keyName, uint32(reflect.ValueOf(v).Uint()))
		}
	case "d":
		prop.valueType = float64Type
		prop.getFn = func() interface{} {
			return s.GetDouble(keyName)
		}
		prop.setFn = func(v interface{}) {
			s.SetDouble(keyName, reflect.ValueOf(v).Float())
		}
	case "s":
		prop.valueType = stringType
		prop.getFn = func() interface{} {
			return s.GetString(keyName)
		}
		prop.setFn = func(v interface{}) {
			s.SetString(keyName, reflect.ValueOf(v).String())
		}
	case "as":
		prop.valueType = strvType
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
		prop.Notify()
	})
	prop.ConnectChanged(func() {
		dbus.NotifyChange(prop.core, prop.propName)
	})
	return prop
}

func strvEqual(a []string, b []string) bool {
	an := len(a)
	bn := len(b)
	if an != bn {
		return false
	}
	for i, v := range a {
		if b[i] != v {
			return false
		}
	}
	return true
}
func (p _GSettingsProperty) SetValue(v interface{}) {
	// []string is not comparable
	if strv, ok := v.([]string); ok {
		oldv, ok := p.getFn().([]string)
		if ok && !strvEqual(strv, oldv) {
			p.setFn(v)
			dbus.NotifyChange(p.core, p.propName)
		}
	} else {
		if v != p.getFn() {
			p.setFn(v)
			dbus.NotifyChange(p.core, p.propName)
		}
	}
}

func (p _GSettingsProperty) GetValue() interface{} {
	return p.getFn()
}

func (p _GSettingsProperty) GetType() reflect.Type {
	return p.valueType
}
