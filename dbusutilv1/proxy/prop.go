// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
// 
// SPDX-License-Identifier: GPL-3.0-or-later
package proxy

import (
	"errors"
	"fmt"

	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/mock"
)

var errNilCallback = errors.New("nil callback")

type PropBool interface {
	Get(flags dbus.Flags) (value bool, err error)
	Set(flags dbus.Flags, value bool) error
	ConnectChanged(cb func(hasValue bool, value bool)) error
}

type ImplPropBool struct {
	Impl Implementer
	Name string
}

func (p ImplPropBool) Get(flags dbus.Flags) (value bool, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropBool) Set(flags dbus.Flags, value bool) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropBool) ConnectChanged(cb func(hasValue bool, value bool)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(bool)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, false)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropBool struct {
	mock.Mock
}

func (p *MockPropBool) Get(flags dbus.Flags) (value bool, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropBool) Set(flags dbus.Flags, value bool) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropBool) ConnectChanged(cb func(hasValue bool, value bool)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropString interface {
	Get(flags dbus.Flags) (value string, err error)
	Set(flags dbus.Flags, value string) error
	ConnectChanged(cb func(hasValue bool, value string)) error
}

type ImplPropString struct {
	Impl Implementer
	Name string
}

func (p ImplPropString) Get(flags dbus.Flags) (value string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropString) Set(flags dbus.Flags, value string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropString) ConnectChanged(cb func(hasValue bool, value string)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(string)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, "")
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropString struct {
	mock.Mock
}

func (p *MockPropString) Get(flags dbus.Flags) (value string, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropString) Set(flags dbus.Flags, value string) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropString) ConnectChanged(cb func(hasValue bool, value string)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropObjectPath interface {
	Get(flags dbus.Flags) (value dbus.ObjectPath, err error)
	Set(flags dbus.Flags, value dbus.ObjectPath) error
	ConnectChanged(cb func(hasValue bool, value dbus.ObjectPath)) error
}

type ImplPropObjectPath struct {
	Impl Implementer
	Name string
}

func (p ImplPropObjectPath) Get(flags dbus.Flags) (value dbus.ObjectPath, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropObjectPath) Set(flags dbus.Flags, value dbus.ObjectPath) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropObjectPath) ConnectChanged(cb func(hasValue bool, value dbus.ObjectPath)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(dbus.ObjectPath)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, "")
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropObjectPath struct {
	mock.Mock
}

func (p *MockPropObjectPath) Get(flags dbus.Flags) (value dbus.ObjectPath, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropObjectPath) Set(flags dbus.Flags, value dbus.ObjectPath) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropObjectPath) ConnectChanged(cb func(hasValue bool, value dbus.ObjectPath)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropDouble interface {
	Get(flags dbus.Flags) (value float64, err error)
	Set(flags dbus.Flags, value float64) error
	ConnectChanged(cb func(hasValue bool, value float64)) error
}

type ImplPropDouble struct {
	Impl Implementer
	Name string
}

func (p ImplPropDouble) Get(flags dbus.Flags) (value float64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropDouble) Set(flags dbus.Flags, value float64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropDouble) ConnectChanged(cb func(hasValue bool, value float64)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(float64)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, 0)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropDouble struct {
	mock.Mock
}

func (p *MockPropDouble) Get(flags dbus.Flags) (value float64, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(float64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropDouble) Set(flags dbus.Flags, value float64) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropDouble) ConnectChanged(cb func(hasValue bool, value float64)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropByte interface {
	Get(flags dbus.Flags) (value byte, err error)
	Set(flags dbus.Flags, value byte) error
	ConnectChanged(cb func(hasValue bool, value byte)) error
}

type ImplPropByte struct {
	Impl Implementer
	Name string
}

func (p ImplPropByte) Get(flags dbus.Flags) (value byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropByte) Set(flags dbus.Flags, value byte) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropByte) ConnectChanged(cb func(hasValue bool, value byte)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(byte)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, 0)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropByte struct {
	mock.Mock
}

func (p *MockPropByte) Get(flags dbus.Flags) (value byte, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(byte)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropByte) Set(flags dbus.Flags, value byte) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropByte) ConnectChanged(cb func(hasValue bool, value byte)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropInt16 interface {
	Get(flags dbus.Flags) (value int16, err error)
	Set(flags dbus.Flags, value int16) error
	ConnectChanged(cb func(hasValue bool, value int16)) error
}

type ImplPropInt16 struct {
	Impl Implementer
	Name string
}

func (p ImplPropInt16) Get(flags dbus.Flags) (value int16, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropInt16) Set(flags dbus.Flags, value int16) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropInt16) ConnectChanged(cb func(hasValue bool, value int16)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(int16)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, 0)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropInt16 struct {
	mock.Mock
}

func (p *MockPropInt16) Get(flags dbus.Flags) (value int16, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(int16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropInt16) Set(flags dbus.Flags, value int16) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropInt16) ConnectChanged(cb func(hasValue bool, value int16)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropUint16 interface {
	Get(flags dbus.Flags) (value uint16, err error)
	Set(flags dbus.Flags, value uint16) error
	ConnectChanged(cb func(hasValue bool, value uint16)) error
}

type ImplPropUint16 struct {
	Impl Implementer
	Name string
}

func (p ImplPropUint16) Get(flags dbus.Flags) (value uint16, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropUint16) Set(flags dbus.Flags, value uint16) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropUint16) ConnectChanged(cb func(hasValue bool, value uint16)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(uint16)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, 0)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropUint16 struct {
	mock.Mock
}

func (p *MockPropUint16) Get(flags dbus.Flags) (value uint16, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(uint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropUint16) Set(flags dbus.Flags, value uint16) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropUint16) ConnectChanged(cb func(hasValue bool, value uint16)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropInt32 interface {
	Get(flags dbus.Flags) (value int32, err error)
	Set(flags dbus.Flags, value int32) error
	ConnectChanged(cb func(hasValue bool, value int32)) error
}

type ImplPropInt32 struct {
	Impl Implementer
	Name string
}

func (p ImplPropInt32) Get(flags dbus.Flags) (value int32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropInt32) Set(flags dbus.Flags, value int32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropInt32) ConnectChanged(cb func(hasValue bool, value int32)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(int32)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, 0)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropInt32 struct {
	mock.Mock
}

func (p *MockPropInt32) Get(flags dbus.Flags) (value int32, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropInt32) Set(flags dbus.Flags, value int32) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropInt32) ConnectChanged(cb func(hasValue bool, value int32)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropUint32 interface {
	Get(flags dbus.Flags) (value uint32, err error)
	Set(flags dbus.Flags, value uint32) error
	ConnectChanged(cb func(hasValue bool, value uint32)) error
}

type ImplPropUint32 struct {
	Impl Implementer
	Name string
}

func (p ImplPropUint32) Get(flags dbus.Flags) (value uint32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropUint32) Set(flags dbus.Flags, value uint32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropUint32) ConnectChanged(cb func(hasValue bool, value uint32)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(uint32)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, 0)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropUint32 struct {
	mock.Mock
}

func (p *MockPropUint32) Get(flags dbus.Flags) (value uint32, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropUint32) Set(flags dbus.Flags, value uint32) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropUint32) ConnectChanged(cb func(hasValue bool, value uint32)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropInt64 interface {
	Get(flags dbus.Flags) (value int64, err error)
	Set(flags dbus.Flags, value int64) error
	ConnectChanged(cb func(hasValue bool, value int64)) error
}

type ImplPropInt64 struct {
	Impl Implementer
	Name string
}

func (p ImplPropInt64) Get(flags dbus.Flags) (value int64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropInt64) Set(flags dbus.Flags, value int64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropInt64) ConnectChanged(cb func(hasValue bool, value int64)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(int64)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, 0)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropInt64 struct {
	mock.Mock
}

func (p *MockPropInt64) Get(flags dbus.Flags) (value int64, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(int64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropInt64) Set(flags dbus.Flags, value int64) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropInt64) ConnectChanged(cb func(hasValue bool, value int64)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropUint64 interface {
	Get(flags dbus.Flags) (value uint64, err error)
	Set(flags dbus.Flags, value uint64) error
	ConnectChanged(cb func(hasValue bool, value uint64)) error
}

type ImplPropUint64 struct {
	Impl Implementer
	Name string
}

func (p ImplPropUint64) Get(flags dbus.Flags) (value uint64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropUint64) Set(flags dbus.Flags, value uint64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropUint64) ConnectChanged(cb func(hasValue bool, value uint64)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.(uint64)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, 0)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropUint64 struct {
	mock.Mock
}

func (p *MockPropUint64) Get(flags dbus.Flags) (value uint64, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).(uint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropUint64) Set(flags dbus.Flags, value uint64) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropUint64) ConnectChanged(cb func(hasValue bool, value uint64)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropBoolArray interface {
	Get(flags dbus.Flags) (value []bool, err error)
	Set(flags dbus.Flags, value []bool) error
	ConnectChanged(cb func(hasValue bool, value []bool)) error
}

type ImplPropBoolArray struct {
	Impl Implementer
	Name string
}

func (p ImplPropBoolArray) Get(flags dbus.Flags) (value []bool, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropBoolArray) Set(flags dbus.Flags, value []bool) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropBoolArray) ConnectChanged(cb func(hasValue bool, value []bool)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]bool)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropBoolArray struct {
	mock.Mock
}

func (p *MockPropBoolArray) Get(flags dbus.Flags) (value []bool, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]bool)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropBoolArray) Set(flags dbus.Flags, value []bool) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropBoolArray) ConnectChanged(cb func(hasValue bool, value []bool)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropStringArray interface {
	Get(flags dbus.Flags) (value []string, err error)
	Set(flags dbus.Flags, value []string) error
	ConnectChanged(cb func(hasValue bool, value []string)) error
}

type ImplPropStringArray struct {
	Impl Implementer
	Name string
}

func (p ImplPropStringArray) Get(flags dbus.Flags) (value []string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropStringArray) Set(flags dbus.Flags, value []string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropStringArray) ConnectChanged(cb func(hasValue bool, value []string)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]string)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropStringArray struct {
	mock.Mock
}

func (p *MockPropStringArray) Get(flags dbus.Flags) (value []string, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]string)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropStringArray) Set(flags dbus.Flags, value []string) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropStringArray) ConnectChanged(cb func(hasValue bool, value []string)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropObjectPathArray interface {
	Get(flags dbus.Flags) (value []dbus.ObjectPath, err error)
	Set(flags dbus.Flags, value []dbus.ObjectPath) error
	ConnectChanged(cb func(hasValue bool, value []dbus.ObjectPath)) error
}

type ImplPropObjectPathArray struct {
	Impl Implementer
	Name string
}

func (p ImplPropObjectPathArray) Get(flags dbus.Flags) (value []dbus.ObjectPath, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropObjectPathArray) Set(flags dbus.Flags, value []dbus.ObjectPath) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropObjectPathArray) ConnectChanged(cb func(hasValue bool, value []dbus.ObjectPath)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]dbus.ObjectPath)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropObjectPathArray struct {
	mock.Mock
}

func (p *MockPropObjectPathArray) Get(flags dbus.Flags) (value []dbus.ObjectPath, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]dbus.ObjectPath)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropObjectPathArray) Set(flags dbus.Flags, value []dbus.ObjectPath) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropObjectPathArray) ConnectChanged(cb func(hasValue bool, value []dbus.ObjectPath)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropDoubleArray interface {
	Get(flags dbus.Flags) (value []float64, err error)
	Set(flags dbus.Flags, value []float64) error
	ConnectChanged(cb func(hasValue bool, value []float64)) error
}

type ImplPropDoubleArray struct {
	Impl Implementer
	Name string
}

func (p ImplPropDoubleArray) Get(flags dbus.Flags) (value []float64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropDoubleArray) Set(flags dbus.Flags, value []float64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropDoubleArray) ConnectChanged(cb func(hasValue bool, value []float64)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]float64)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropDoubleArray struct {
	mock.Mock
}

func (p *MockPropDoubleArray) Get(flags dbus.Flags) (value []float64, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]float64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropDoubleArray) Set(flags dbus.Flags, value []float64) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropDoubleArray) ConnectChanged(cb func(hasValue bool, value []float64)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropByteArray interface {
	Get(flags dbus.Flags) (value []byte, err error)
	Set(flags dbus.Flags, value []byte) error
	ConnectChanged(cb func(hasValue bool, value []byte)) error
}

type ImplPropByteArray struct {
	Impl Implementer
	Name string
}

func (p ImplPropByteArray) Get(flags dbus.Flags) (value []byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropByteArray) Set(flags dbus.Flags, value []byte) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropByteArray) ConnectChanged(cb func(hasValue bool, value []byte)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]byte)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropByteArray struct {
	mock.Mock
}

func (p *MockPropByteArray) Get(flags dbus.Flags) (value []byte, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]byte)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropByteArray) Set(flags dbus.Flags, value []byte) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropByteArray) ConnectChanged(cb func(hasValue bool, value []byte)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropInt16Array interface {
	Get(flags dbus.Flags) (value []int16, err error)
	Set(flags dbus.Flags, value []int16) error
	ConnectChanged(cb func(hasValue bool, value []int16)) error
}

type ImplPropInt16Array struct {
	Impl Implementer
	Name string
}

func (p ImplPropInt16Array) Get(flags dbus.Flags) (value []int16, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropInt16Array) Set(flags dbus.Flags, value []int16) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropInt16Array) ConnectChanged(cb func(hasValue bool, value []int16)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]int16)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropInt16Array struct {
	mock.Mock
}

func (p *MockPropInt16Array) Get(flags dbus.Flags) (value []int16, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]int16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropInt16Array) Set(flags dbus.Flags, value []int16) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropInt16Array) ConnectChanged(cb func(hasValue bool, value []int16)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropUint16Array interface {
	Get(flags dbus.Flags) (value []uint16, err error)
	Set(flags dbus.Flags, value []uint16) error
	ConnectChanged(cb func(hasValue bool, value []uint16)) error
}

type ImplPropUint16Array struct {
	Impl Implementer
	Name string
}

func (p ImplPropUint16Array) Get(flags dbus.Flags) (value []uint16, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropUint16Array) Set(flags dbus.Flags, value []uint16) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropUint16Array) ConnectChanged(cb func(hasValue bool, value []uint16)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]uint16)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropUint16Array struct {
	mock.Mock
}

func (p *MockPropUint16Array) Get(flags dbus.Flags) (value []uint16, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]uint16)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropUint16Array) Set(flags dbus.Flags, value []uint16) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropUint16Array) ConnectChanged(cb func(hasValue bool, value []uint16)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropInt32Array interface {
	Get(flags dbus.Flags) (value []int32, err error)
	Set(flags dbus.Flags, value []int32) error
	ConnectChanged(cb func(hasValue bool, value []int32)) error
}

type ImplPropInt32Array struct {
	Impl Implementer
	Name string
}

func (p ImplPropInt32Array) Get(flags dbus.Flags) (value []int32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropInt32Array) Set(flags dbus.Flags, value []int32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropInt32Array) ConnectChanged(cb func(hasValue bool, value []int32)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]int32)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropInt32Array struct {
	mock.Mock
}

func (p *MockPropInt32Array) Get(flags dbus.Flags) (value []int32, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]int32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropInt32Array) Set(flags dbus.Flags, value []int32) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropInt32Array) ConnectChanged(cb func(hasValue bool, value []int32)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropUint32Array interface {
	Get(flags dbus.Flags) (value []uint32, err error)
	Set(flags dbus.Flags, value []uint32) error
	ConnectChanged(cb func(hasValue bool, value []uint32)) error
}

type ImplPropUint32Array struct {
	Impl Implementer
	Name string
}

func (p ImplPropUint32Array) Get(flags dbus.Flags) (value []uint32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropUint32Array) Set(flags dbus.Flags, value []uint32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropUint32Array) ConnectChanged(cb func(hasValue bool, value []uint32)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]uint32)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropUint32Array struct {
	mock.Mock
}

func (p *MockPropUint32Array) Get(flags dbus.Flags) (value []uint32, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]uint32)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropUint32Array) Set(flags dbus.Flags, value []uint32) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropUint32Array) ConnectChanged(cb func(hasValue bool, value []uint32)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropInt64Array interface {
	Get(flags dbus.Flags) (value []int64, err error)
	Set(flags dbus.Flags, value []int64) error
	ConnectChanged(cb func(hasValue bool, value []int64)) error
}

type ImplPropInt64Array struct {
	Impl Implementer
	Name string
}

func (p ImplPropInt64Array) Get(flags dbus.Flags) (value []int64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropInt64Array) Set(flags dbus.Flags, value []int64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropInt64Array) ConnectChanged(cb func(hasValue bool, value []int64)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]int64)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropInt64Array struct {
	mock.Mock
}

func (p *MockPropInt64Array) Get(flags dbus.Flags) (value []int64, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]int64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropInt64Array) Set(flags dbus.Flags, value []int64) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropInt64Array) ConnectChanged(cb func(hasValue bool, value []int64)) error {
	args := p.Called(cb)

	return args.Error(0)
}

type PropUint64Array interface {
	Get(flags dbus.Flags) (value []uint64, err error)
	Set(flags dbus.Flags, value []uint64) error
	ConnectChanged(cb func(hasValue bool, value []uint64)) error
}

type ImplPropUint64Array struct {
	Impl Implementer
	Name string
}

func (p ImplPropUint64Array) Get(flags dbus.Flags) (value []uint64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p ImplPropUint64Array) Set(flags dbus.Flags, value []uint64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p ImplPropUint64Array) ConnectChanged(cb func(hasValue bool, value []uint64)) error {
	if cb == nil {
		return errNilCallback
	}
	cb0 := func(hasValue bool, value interface{}) {
		if hasValue {
			val, ok := value.([]uint64)
			if ok {
				cb(true, val)
			}
		} else {
			cb(false, nil)
		}
	}
	return p.Impl.GetObject_().ConnectPropertyChanged_(p.Impl.GetInterfaceName_(), p.Name, cb0)
}

type MockPropUint64Array struct {
	mock.Mock
}

func (p *MockPropUint64Array) Get(flags dbus.Flags) (value []uint64, err error) {
	args := p.Called(flags)

	var ok bool
	value, ok = args.Get(0).([]uint64)
	if !ok {
		panic(fmt.Sprintf("assert: arguments: %d failed because object wasn't correct type: %v", 0, args.Get(0)))
	}

	err = args.Error(1)

	return
}

func (p *MockPropUint64Array) Set(flags dbus.Flags, value []uint64) error {
	args := p.Called(flags, value)

	return args.Error(0)
}

func (p *MockPropUint64Array) ConnectChanged(cb func(hasValue bool, value []uint64)) error {
	args := p.Called(cb)

	return args.Error(0)
}
