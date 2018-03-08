package gsprop

import (
	"errors"
	"reflect"

	"gir/gio-2.0"
	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
	"pkg.deepin.io/lib/gsettings"
)

type base struct {
	gs            *gio.Settings
	key           string
	notifyChanged func(val interface{})
}

func (b *base) bind(gs *gio.Settings, keyName string,
	getFn func() (interface{}, *dbus.Error)) {

	b.gs = gs
	b.key = keyName

	var propPath string
	gs.GetProperty("path", &propPath)
	gsettings.ConnectChanged(propPath, keyName, func(_ string) {
		val, _ := getFn()
		b.notifyChanged(val)
	})
}

func (b *base) SetNotifyChangedFunc(fn func(val interface{})) {
	b.notifyChanged = fn
}

func checkSet(setOk bool) *dbus.Error {
	if setOk {
		return nil
	}
	return dbusutil.ToError(errors.New("write failed"))
}

type Bool struct {
	base
}

func (b *Bool) Bind(gs *gio.Settings, key string) {
	b.bind(gs, key, b.GetValue)
}

func (b *Bool) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	err = checkSet(b.Set(val.(bool)))
	return
}

func (b *Bool) GetValue() (val interface{}, err *dbus.Error) {
	val = b.Get()
	return
}

func (b *Bool) Get() bool {
	return b.gs.GetBoolean(b.key)
}

func (b *Bool) Set(val bool) bool {
	if b.Get() != val {
		return b.gs.SetBoolean(b.key, val)
	}
	return true
}

func (*Bool) GetType() reflect.Type {
	return reflect.TypeOf(false)
}

type Int struct {
	base
}

func (i *Int) Bind(gs *gio.Settings, key string) {
	i.bind(gs, key, i.GetValue)
}

func (i *Int) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	err = checkSet(i.Set(val.(int32)))
	return
}

func (i *Int) GetValue() (val interface{}, err *dbus.Error) {
	val = i.Get()
	return
}

func (*Int) GetType() reflect.Type {
	return reflect.TypeOf(int32(0))
}

func (i *Int) Set(val int32) bool {
	if i.Get() != val {
		return i.gs.SetInt(i.key, val)
	}
	return true
}

func (i *Int) Get() int32 {
	return i.gs.GetInt(i.key)
}

type Enum struct {
	base
}

func (e *Enum) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	err = checkSet(e.Set(val.(int32)))
	return
}

func (e *Enum) GetValue() (val interface{}, err *dbus.Error) {
	val = e.Get()
	return
}

func (e *Enum) GetType() reflect.Type {
	return reflect.TypeOf(int32(0))
}

func (e *Enum) Bind(gs *gio.Settings, key string) {
	e.bind(gs, key, e.GetValue)
}

func (e *Enum) Get() int32 {
	return e.gs.GetEnum(e.key)
}

func (e *Enum) Set(val int32) bool {
	if e.Get() != val {
		return e.gs.SetEnum(e.key, val)
	}
	return true
}

type Uint struct {
	base
}

func (u *Uint) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	err = checkSet(u.Set(val.(uint32)))
	return
}

func (u *Uint) GetValue() (val interface{}, err *dbus.Error) {
	val = u.Get()
	return
}

func (*Uint) GetType() reflect.Type {
	return reflect.TypeOf(uint32(0))
}

func (u *Uint) Get() uint32 {
	return u.gs.GetUint(u.key)
}

func (u *Uint) Set(val uint32) bool {
	if u.Get() != val {
		return u.gs.SetUint(u.key, val)
	}
	return true
}

func (u *Uint) Bind(gs *gio.Settings, key string) {
	u.bind(gs, key, u.GetValue)
}

type Double struct {
	base
}

func (d *Double) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	err = checkSet(d.Set(val.(float64)))
	return
}

func (d *Double) GetValue() (val interface{}, err *dbus.Error) {
	val = d.Get()
	return
}

func (*Double) GetType() reflect.Type {
	return reflect.TypeOf(float64(0))
}

func (d *Double) Get() float64 {
	return d.gs.GetDouble(d.key)
}

func (d *Double) Set(val float64) bool {
	if d.Get() != val {
		return d.gs.SetDouble(d.key, val)
	}
	return true
}

func (d *Double) Bind(gs *gio.Settings, key string) {
	d.bind(gs, key, d.GetValue)
}

type String struct {
	base
}

func (s *String) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	err = checkSet(s.Set(val.(string)))
	return
}

func (s *String) GetValue() (val interface{}, err *dbus.Error) {
	val = s.Get()
	return
}

func (*String) GetType() reflect.Type {
	return reflect.TypeOf("")
}

func (s *String) Get() string {
	return s.gs.GetString(s.key)
}

func (s *String) Set(val string) bool {
	if s.Get() != val {
		return s.gs.SetString(s.key, val)
	}
	return true
}

func (s *String) Bind(gs *gio.Settings, key string) {
	s.bind(gs, key, s.GetValue)
}

type Strv struct {
	base
}

func (s *Strv) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	err = checkSet(s.Set(val.([]string)))
	return
}

func (s *Strv) GetValue() (val interface{}, err *dbus.Error) {
	val = s.Get()
	return
}

func (*Strv) GetType() reflect.Type {
	return reflect.TypeOf([]string{})
}

func (s *Strv) Get() []string {
	return s.gs.GetStrv(s.key)
}

func (s *Strv) Set(val []string) bool {
	if !strvEqual(s.Get(), val) {
		return s.gs.SetStrv(s.key, val)
	}
	return true
}

func (s *Strv) Bind(gs *gio.Settings, key string) {
	s.bind(gs, key, s.GetValue)
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
