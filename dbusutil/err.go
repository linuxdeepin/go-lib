// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusutil

import (
	"encoding/json"
	"fmt"

	"github.com/godbus/dbus/v5"
)

func MakeError(v Implementer, name string, args ...interface{}) *dbus.Error {
	interfaceName := ""
	if implV20, ok := v.(ImplementerV20); ok {
		interfaceName = implV20.GetInterfaceName() + "."
	}
	errName := interfaceName + "Error." + name
	msg := fmt.Sprint(args...)
	return &dbus.Error{
		Name: errName,
		Body: []interface{}{msg},
	}
}

func MakeErrorf(v Implementer, name, format string, args ...interface{}) *dbus.Error {
	interfaceName := ""
	if implV20, ok := v.(ImplementerV20); ok {
		interfaceName = implV20.GetInterfaceName() + "."
	}
	errName := interfaceName + "Error." + name
	msg := fmt.Sprintf(format, args...)
	return &dbus.Error{
		Name: errName,
		Body: []interface{}{msg},
	}
}

func MakeErrorJSON(v Implementer, name string, detail interface{}) *dbus.Error {
	var msg string
	interfaceName := ""
	if implV20, ok := v.(ImplementerV20); ok {
		interfaceName = implV20.GetInterfaceName() + "."
	}
	errName := interfaceName + "Error." + name
	data, err := json.Marshal(detail)
	if err != nil {
		msg = "failed to marshal json"
	} else {
		msg = string(data)
	}
	return &dbus.Error{
		Name: errName,
		Body: []interface{}{msg},
	}
}

type DBusError interface {
	Error() string
	Name() string
}

func ToError(err error) *dbus.Error {
	if err == nil {
		return nil
	}
	err0, ok := err.(DBusError)
	var name, msg string
	if ok {
		name = err0.Name()
		msg = err0.Error()
	} else {
		name = "org.deepin.dde.DBus.Error.Unnamed"
		msg = err.Error()
	}
	return &dbus.Error{
		Name: name,
		Body: []interface{}{msg},
	}
}
