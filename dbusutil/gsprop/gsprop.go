package gsprop

import (
	"reflect"

	"gir/gio-2.0"
	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/gsettings"
)

type baseProperty struct {
	valueType     reflect.Type
	getFn         func() interface{}
	setFn         func(interface{})
	notifyChanged func(val interface{})
}

var _ dbusutil.Property = &baseProperty{}

var (
	boolType    = reflect.TypeOf(false)
	int32Type   = reflect.TypeOf(int32(0))
	uint32Type  = reflect.TypeOf(uint32(0))
	float64Type = reflect.TypeOf(float64(0))
	stringType  = reflect.TypeOf("")
	strvType    = reflect.TypeOf([]string{})
)

func newBaseProperty(sig string, s *gio.Settings, keyName string) *baseProperty {
	realType := s.GetValue(keyName).GetTypeString()
	if realType == "s" && sig == "e" {
		// note: "e" is not matched with glib type system
		realType = "e"
	}
	if realType != sig {
		var correctMethod string
		switch realType {
		case "b":
			correctMethod = "NewBool"
		case "i":
			correctMethod = "NewInt"
		case "u":
			correctMethod = "NewUin"
		case "d":
			correctMethod = "NewFloat"
		case "s":
			correctMethod = "NewString"
		case "as":
			correctMethod = "NewStrv"
		default:
			panic("GSettingsProperty didn't type " + sig)
		}
		panic("Type signal " + sig + " didn't match " + keyName + "'s type(" + realType + ")" + ", please use the method " + correctMethod)
	}

	prop := &baseProperty{}
	switch realType {
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
	var schemaId string
	s.GetProperty("schema-id", &schemaId)

	gsettings.ConnectChanged(schemaId, keyName, func(key string) {
		prop.notifyChanged(prop.getFn())
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

func (p *baseProperty) GetValue() (interface{}, *dbus.Error) {
	return p.getFn(), nil
}

func (p *baseProperty) SetValue(val interface{}) (bool, *dbus.Error) {
	// []string is not comparable
	if strv, ok := val.([]string); ok {
		oldStrv, ok := p.getFn().([]string)
		if ok && !strvEqual(strv, oldStrv) {
			p.setFn(val)
		}
	} else {
		if val != p.getFn() {
			p.setFn(val)
		}
	}

	// NOTE: SetValue changed always is false
	return false, nil
}

func (p *baseProperty) SetNotifyChangedFunc(fn func(val interface{})) {
	p.notifyChanged = fn
}

func (p *baseProperty) GetType() reflect.Type {
	return p.valueType
}

type Bool struct{ *baseProperty }

func NewBool(s *gio.Settings, keyName string) *Bool {
	return &Bool{newBaseProperty("b", s, keyName)}
}

func (prop *Bool) Get() bool {
	v, _ := prop.GetValue()
	return v.(bool)
}

func (prop *Bool) Set(v bool) {
	prop.SetValue(v)
}

type Int struct{ *baseProperty }

func NewInt(s *gio.Settings, keyName string) *Int {
	return &Int{newBaseProperty("i", s, keyName)}
}

func (prop *Int) Get() int32 {
	v, _ := prop.GetValue()
	return v.(int32)
}

func (prop *Int) Set(v int32) {
	prop.SetValue(v)
}

type Enum struct{ *baseProperty }

func NewEnum(s *gio.Settings, keyName string) *Enum {
	return &Enum{newBaseProperty("e", s, keyName)}
}

func (prop *Enum) Get() int32 {
	v, _ := prop.GetValue()
	return v.(int32)
}

func (prop *Enum) Set(v int32) {
	prop.SetValue(v)
}

type Uint struct{ *baseProperty }

func NewUint(s *gio.Settings, keyName string) *Uint {
	return &Uint{newBaseProperty("u", s, keyName)}
}

func (prop *Uint) Get() uint32 {
	v, _ := prop.GetValue()
	return v.(uint32)
}

func (prop *Uint) Set(v uint32) {
	prop.SetValue(v)
}

type Float struct{ *baseProperty }

func NewFloat(s *gio.Settings, keyName string) *Float {
	return &Float{newBaseProperty("d", s, keyName)}
}

func (prop *Float) Get() float64 {
	v, _ := prop.GetValue()
	return v.(float64)
}

func (prop *Float) Set(v float64) {
	prop.SetValue(v)
}

type String struct{ *baseProperty }

func NewString(s *gio.Settings, keyName string) *String {
	return &String{newBaseProperty("s", s, keyName)}
}

func (prop *String) Get() string {
	v, _ := prop.GetValue()
	return v.(string)
}

func (prop *String) Set(v string) {
	prop.SetValue(v)
}

type Strv struct{ *baseProperty }

func NewStrv(s *gio.Settings, keyName string) *Strv {
	return &Strv{newBaseProperty("as", s, keyName)}
}

func (prop *Strv) Get() []string {
	v, _ := prop.GetValue()
	return v.([]string)
}

func (prop *Strv) Set(v []string) {
	prop.SetValue(v)
}
