// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
// 
// SPDX-License-Identifier: GPL-3.0-or-later
package dbusutilv1

import (
	"reflect"
	"testing"

	"github.com/godbus/dbus/v5"
)

type impl1 struct {
}

func (*impl1) GetExportedMethods() ExportedMethods {
	return nil
}

func (*impl1) GetInterfaceName() string {
	return "org.deepin.dde.lib.Exportable1"
}

func TestMakeError(t *testing.T) {
	err := MakeError(&impl1{}, "Err1", "abc", 123)
	expectedErr := &dbus.Error{
		Name: "Error.Err1",
		Body: []interface{}{"abc123"},
	}
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("err expected %#v, but got %#v", expectedErr, err)
	}
}

func TestMakeErrorf(t *testing.T) {
	err := MakeErrorf(&impl1{}, "Err2", "name: %s, num: %d", "abc", 123)
	expectedErr := &dbus.Error{
		Name: "Error.Err2",
		Body: []interface{}{"name: abc, num: 123"},
	}
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("err expected %#v, but got %#v", expectedErr, err)
	}
}

func TestMakeErrorJSON(t *testing.T) {
	detail := &struct {
		Name string `json:"name"`
		Num  int    `json:"num"`
	}{"abc", 123}
	err := MakeErrorJSON(&impl1{}, "Err3", detail)

	expectedErr := &dbus.Error{
		Name: "Error.Err3",
		Body: []interface{}{`{"name":"abc","num":123}`},
	}

	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("err expected %#v, but got %#v", expectedErr, err)
	}

	detail1 := &struct {
		Chan chan int
	}{
		nil,
	}
	err = MakeErrorJSON(&impl1{}, "Err3", detail1)
	// do not panic
	if err == nil {
		t.Errorf("Expected error due to detail1 has chan type field")
	}
}

type unnamedError struct{}

func (err unnamedError) Error() string {
	return "xxx err msg"
}

type namedError struct{}

func (err namedError) Error() string {
	return "yyy err msg"
}

func (err namedError) Name() string {
	return "Error.Err4"
}

func TestToError(t *testing.T) {
	err := ToError(unnamedError{})
	expectedErr := &dbus.Error{
		Name: "org.deepin.dde.DBus.Error.Unnamed",
		Body: []interface{}{"xxx err msg"},
	}

	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("err expected %#v, but got %#v", expectedErr, err)
	}

	err = ToError(namedError{})
	expectedErr = &dbus.Error{
		Name: "Error.Err4",
		Body: []interface{}{"yyy err msg"},
	}

	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("err expected %#v, but got %#v", expectedErr, err)
	}

	err = ToError(nil)
	if err != nil {
		t.Errorf("err expected nil, but got %#v", err)
	}
}
