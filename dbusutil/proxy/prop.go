package proxy

import "errors"
import "github.com/godbus/dbus"

var errNilCallback = errors.New("nil callback")

type PropBool struct {
	Impl Implementer
	Name string
}

func (p PropBool) Get(flags dbus.Flags) (value bool, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropBool) Set(flags dbus.Flags, value bool) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropBool) ConnectChanged(cb func(hasValue bool, value bool)) error {
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

type PropString struct {
	Impl Implementer
	Name string
}

func (p PropString) Get(flags dbus.Flags) (value string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropString) Set(flags dbus.Flags, value string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropString) ConnectChanged(cb func(hasValue bool, value string)) error {
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

type PropObjectPath struct {
	Impl Implementer
	Name string
}

func (p PropObjectPath) Get(flags dbus.Flags) (value dbus.ObjectPath, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropObjectPath) Set(flags dbus.Flags, value dbus.ObjectPath) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropObjectPath) ConnectChanged(cb func(hasValue bool, value dbus.ObjectPath)) error {
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

type PropDouble struct {
	Impl Implementer
	Name string
}

func (p PropDouble) Get(flags dbus.Flags) (value float64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropDouble) Set(flags dbus.Flags, value float64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropDouble) ConnectChanged(cb func(hasValue bool, value float64)) error {
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

type PropByte struct {
	Impl Implementer
	Name string
}

func (p PropByte) Get(flags dbus.Flags) (value byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropByte) Set(flags dbus.Flags, value byte) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropByte) ConnectChanged(cb func(hasValue bool, value byte)) error {
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

type PropInt16 struct {
	Impl Implementer
	Name string
}

func (p PropInt16) Get(flags dbus.Flags) (value int16, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropInt16) Set(flags dbus.Flags, value int16) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropInt16) ConnectChanged(cb func(hasValue bool, value int16)) error {
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

type PropUint16 struct {
	Impl Implementer
	Name string
}

func (p PropUint16) Get(flags dbus.Flags) (value uint16, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropUint16) Set(flags dbus.Flags, value uint16) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropUint16) ConnectChanged(cb func(hasValue bool, value uint16)) error {
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

type PropInt32 struct {
	Impl Implementer
	Name string
}

func (p PropInt32) Get(flags dbus.Flags) (value int32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropInt32) Set(flags dbus.Flags, value int32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropInt32) ConnectChanged(cb func(hasValue bool, value int32)) error {
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

type PropUint32 struct {
	Impl Implementer
	Name string
}

func (p PropUint32) Get(flags dbus.Flags) (value uint32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropUint32) Set(flags dbus.Flags, value uint32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropUint32) ConnectChanged(cb func(hasValue bool, value uint32)) error {
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

type PropInt64 struct {
	Impl Implementer
	Name string
}

func (p PropInt64) Get(flags dbus.Flags) (value int64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropInt64) Set(flags dbus.Flags, value int64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropInt64) ConnectChanged(cb func(hasValue bool, value int64)) error {
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

type PropUint64 struct {
	Impl Implementer
	Name string
}

func (p PropUint64) Get(flags dbus.Flags) (value uint64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropUint64) Set(flags dbus.Flags, value uint64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropUint64) ConnectChanged(cb func(hasValue bool, value uint64)) error {
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

type PropBoolArray struct {
	Impl Implementer
	Name string
}

func (p PropBoolArray) Get(flags dbus.Flags) (value []bool, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropBoolArray) Set(flags dbus.Flags, value []bool) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropBoolArray) ConnectChanged(cb func(hasValue bool, value []bool)) error {
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

type PropStringArray struct {
	Impl Implementer
	Name string
}

func (p PropStringArray) Get(flags dbus.Flags) (value []string, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropStringArray) Set(flags dbus.Flags, value []string) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropStringArray) ConnectChanged(cb func(hasValue bool, value []string)) error {
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

type PropObjectPathArray struct {
	Impl Implementer
	Name string
}

func (p PropObjectPathArray) Get(flags dbus.Flags) (value []dbus.ObjectPath, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropObjectPathArray) Set(flags dbus.Flags, value []dbus.ObjectPath) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropObjectPathArray) ConnectChanged(cb func(hasValue bool, value []dbus.ObjectPath)) error {
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

type PropDoubleArray struct {
	Impl Implementer
	Name string
}

func (p PropDoubleArray) Get(flags dbus.Flags) (value []float64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropDoubleArray) Set(flags dbus.Flags, value []float64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropDoubleArray) ConnectChanged(cb func(hasValue bool, value []float64)) error {
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

type PropByteArray struct {
	Impl Implementer
	Name string
}

func (p PropByteArray) Get(flags dbus.Flags) (value []byte, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropByteArray) Set(flags dbus.Flags, value []byte) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropByteArray) ConnectChanged(cb func(hasValue bool, value []byte)) error {
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

type PropInt16Array struct {
	Impl Implementer
	Name string
}

func (p PropInt16Array) Get(flags dbus.Flags) (value []int16, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropInt16Array) Set(flags dbus.Flags, value []int16) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropInt16Array) ConnectChanged(cb func(hasValue bool, value []int16)) error {
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

type PropUint16Array struct {
	Impl Implementer
	Name string
}

func (p PropUint16Array) Get(flags dbus.Flags) (value []uint16, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropUint16Array) Set(flags dbus.Flags, value []uint16) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropUint16Array) ConnectChanged(cb func(hasValue bool, value []uint16)) error {
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

type PropInt32Array struct {
	Impl Implementer
	Name string
}

func (p PropInt32Array) Get(flags dbus.Flags) (value []int32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropInt32Array) Set(flags dbus.Flags, value []int32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropInt32Array) ConnectChanged(cb func(hasValue bool, value []int32)) error {
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

type PropUint32Array struct {
	Impl Implementer
	Name string
}

func (p PropUint32Array) Get(flags dbus.Flags) (value []uint32, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropUint32Array) Set(flags dbus.Flags, value []uint32) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropUint32Array) ConnectChanged(cb func(hasValue bool, value []uint32)) error {
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

type PropInt64Array struct {
	Impl Implementer
	Name string
}

func (p PropInt64Array) Get(flags dbus.Flags) (value []int64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropInt64Array) Set(flags dbus.Flags, value []int64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropInt64Array) ConnectChanged(cb func(hasValue bool, value []int64)) error {
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

type PropUint64Array struct {
	Impl Implementer
	Name string
}

func (p PropUint64Array) Get(flags dbus.Flags) (value []uint64, err error) {
	err = p.Impl.GetObject_().GetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, &value)
	return
}

func (p PropUint64Array) Set(flags dbus.Flags, value []uint64) error {
	return p.Impl.GetObject_().SetProperty_(flags, p.Impl.GetInterfaceName_(), p.Name, value)
}

func (p PropUint64Array) ConnectChanged(cb func(hasValue bool, value []uint64)) error {
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
