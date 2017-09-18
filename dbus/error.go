/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package dbus

import "reflect"

// Error represents a D-Bus message of type Error.
type dbusError struct {
	Name string
	Body []interface{}
}

var (
	goErrorType   = reflect.TypeOf((*error)(nil)).Elem()
	dbusErrorType = reflect.TypeOf((*dbusError)(nil))
)

func (e dbusError) Error() string {
	if len(e.Body) >= 1 {
		s, ok := e.Body[0].(string)
		if ok {
			return e.Name + ":" + s
		}
	}
	return e.Name
}

const (
	NoObjectError = iota
	UnknowInterfaceError
	UnknowMethodError
	UnknowPropertyError
	OtherError
)

func NewCustomError(name string, args ...interface{}) dbusError {
	return dbusError{
		name,
		args,
	}
}
func NewNoObjectError(path ObjectPath) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.NoSuchObject",
		[]interface{}{"No such object" + string(path)},
	}
}
func newError(errType int, args ...interface{}) dbusError {
	//TODO: complete this
	name := "UnknowError"
	switch errType {
	case NoObjectError:
		name = "org.freedesktop.DBus.Error.NoSuchObject"
	}
	return dbusError{
		name,
		args,
	}
}

func NewPropertyNotWritableError(name string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.NoWritable",
		[]interface{}{"Can't write this property."},
	}
}

func NewUnknowInterfaceError(ifcName string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.NoSuchInterface",
		[]interface{}{"No such interface"},
	}
}
func NewUnknowPropertyError(name string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.UnknownProperty",
		[]interface{}{"Unknown / invalid Property"},
	}
}

func NewOtherError(body interface{}) dbusError {
	return dbusError{
		"com.deepin.DBus.Error.UnknowError",
		[]interface{}{body},
	}
}
func newInternalError(body interface{}) dbusError {
	return dbusError{
		"com.deepin.DBus.Error.InternalError",
		[]interface{}{body},
	}
}
func NewUnknowMethod(path ObjectPath, ifc, name string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.UnknownMethod",
		[]interface{}{"Cant find the method of " + name},
	}
}
func NewInvalidArg(content string) dbusError {
	return dbusError{
		"org.freedesktop.DBus.Error.InvalidArgs",
		[]interface{}{"Invalid type / number of args" + content},
	}
}
