// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package gsprop

import (
	"errors"
	"reflect"
	"sync"

	"github.com/godbus/dbus/v5"
	gio "github.com/linuxdeepin/go-gir/gio-2.0"
	"github.com/linuxdeepin/go-lib/dbusutil"
	"github.com/linuxdeepin/go-lib/gsettings"
)

type base struct {
	mu                sync.Mutex
	gs                *gio.Settings
	key               string
	notifyChangedList []func(val interface{})
}

func (b *base) bind(gs *gio.Settings, keyName string,
	getFn func() (interface{}, *dbus.Error)) {

	b.gs = gs
	b.key = keyName

	var propPath string
	gs.GetProperty("path", &propPath)
	gsettings.ConnectChanged(propPath, keyName, func(_ string) {
		b.mu.Lock()
		notifyChangedList := b.notifyChangedList
		b.mu.Unlock()
		if notifyChangedList != nil {
			val, _ := getFn()
			for _, notifyChanged := range notifyChangedList {
				if notifyChanged != nil {
					notifyChanged(val)
				}
			}
		}
	})
}

func (b *base) SetNotifyChangedFunc(fn func(val interface{})) {
	b.mu.Lock()
	b.notifyChangedList = append(b.notifyChangedList, fn)
	b.mu.Unlock()
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
	b.mu.Lock()
	v := b.gs.GetBoolean(b.key)
	b.mu.Unlock()
	return v
}

func (b *Bool) Set(val bool) bool {
	if b.Get() != val {
		b.mu.Lock()
		ok := b.gs.SetBoolean(b.key, val)
		b.mu.Unlock()
		return ok
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
		i.mu.Lock()
		ok := i.gs.SetInt(i.key, val)
		i.mu.Unlock()
		return ok
	}
	return true
}

func (i *Int) Get() int32 {
	i.mu.Lock()
	v := i.gs.GetInt(i.key)
	i.mu.Unlock()
	return v
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
	e.mu.Lock()
	v := e.gs.GetEnum(e.key)
	e.mu.Unlock()
	return v
}

func (e *Enum) GetString() string {
	e.mu.Lock()
	v := e.gs.GetString(e.key)
	e.mu.Unlock()
	return v
}

func (e *Enum) Set(val int32) bool {
	if e.Get() != val {
		e.mu.Lock()
		ok := e.gs.SetEnum(e.key, val)
		e.mu.Unlock()
		return ok
	}
	return true
}

func (e *Enum) SetString(val string) bool {
	if e.GetString() != val {
		e.mu.Lock()
		ok := e.gs.SetString(e.key, val)
		e.mu.Unlock()
		return ok
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
	u.mu.Lock()
	v := u.gs.GetUint(u.key)
	u.mu.Unlock()
	return v
}

func (u *Uint) Set(val uint32) bool {
	if u.Get() != val {
		u.mu.Lock()
		ok := u.gs.SetUint(u.key, val)
		u.mu.Unlock()
		return ok
	}
	return true
}

func (u *Uint) Bind(gs *gio.Settings, key string) {
	u.bind(gs, key, u.GetValue)
}

type Uint64 struct {
	base
}

func (u *Uint64) SetValue(val interface{}) (changed bool, err *dbus.Error) {
	err = checkSet(u.Set(val.(uint64)))
	return
}

func (u *Uint64) GetValue() (val interface{}, err *dbus.Error) {
	val = u.Get()
	return
}

func (*Uint64) GetType() reflect.Type {
	return reflect.TypeOf(uint64(0))
}

func (u *Uint64) Get() uint64 {
	u.mu.Lock()
	v := u.gs.GetUint64(u.key)
	u.mu.Unlock()
	return v
}

func (u *Uint64) Set(val uint64) bool {
	if u.Get() != val {
		u.mu.Lock()
		ok := u.gs.SetUint64(u.key, val)
		u.mu.Unlock()
		return ok
	}
	return true
}

func (u *Uint64) Bind(gs *gio.Settings, key string) {
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
	d.mu.Lock()
	v := d.gs.GetDouble(d.key)
	d.mu.Unlock()
	return v
}

func (d *Double) Set(val float64) bool {
	if d.Get() != val {
		d.mu.Lock()
		ok := d.gs.SetDouble(d.key, val)
		d.mu.Unlock()
		return ok
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
	s.mu.Lock()
	v := s.gs.GetString(s.key)
	s.mu.Unlock()
	return v
}

func (s *String) Set(val string) bool {
	if s.Get() != val {
		s.mu.Lock()
		ok := s.gs.SetString(s.key, val)
		s.mu.Unlock()
		return ok
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
	s.mu.Lock()
	v := s.gs.GetStrv(s.key)
	s.mu.Unlock()
	return v
}

func (s *Strv) Set(val []string) bool {
	if !strvEqual(s.Get(), val) {
		s.mu.Lock()
		ok := s.gs.SetStrv(s.key, val)
		s.mu.Unlock()
		return ok
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
