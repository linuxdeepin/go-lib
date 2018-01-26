package dbusutil

import (
	"reflect"
	"testing"

	"pkg.deepin.io/lib/dbus1"
)

type exportable1 struct {
}

func (*exportable1) GetDBusExportInfo() ExportInfo {
	return ExportInfo{
		Path:      "/",
		Interface: "com.deepin.lib.Exportable1",
	}
}

func TestMakeError(t *testing.T) {
	err := MakeError(&exportable1{}, "Err1", "abc", 123)
	expectedErr := &dbus.Error{
		Name: "com.deepin.lib.Exportable1.Error.Err1",
		Body: []interface{}{"abc123"},
	}
	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("err expected %#v, but got %#v", expectedErr, err)
	}
}

func TestMakeErrorf(t *testing.T) {
	err := MakeErrorf(&exportable1{}, "Err2", "name: %s, num: %d", "abc", 123)
	expectedErr := &dbus.Error{
		Name: "com.deepin.lib.Exportable1.Error.Err2",
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
	err := MakeErrorJSON(&exportable1{}, "Err3", detail)

	expectedErr := &dbus.Error{
		Name: "com.deepin.lib.Exportable1.Error.Err3",
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
	err = MakeErrorJSON(&exportable1{}, "Err3", detail1)
	// do not panic
	if err == nil {
		t.Errorf("Expected error due to detail1 has chan type field")
	}
}

type unamedError struct{}

func (err unamedError) Error() string {
	return "xxx err msg"
}

type namedError struct{}

func (err namedError) Error() string {
	return "yyy err msg"
}

func (err namedError) Name() string {
	return "com.deepin.lib.Exportable1.Error.Err4"
}

func TestToError(t *testing.T) {
	err := ToError(unamedError{})
	expectedErr := &dbus.Error{
		Name: "com.deepin.DBus.Error.Unnamed",
		Body: []interface{}{"xxx err msg"},
	}

	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("err expected %#v, but got %#v", expectedErr, err)
	}

	err = ToError(namedError{})
	expectedErr = &dbus.Error{
		Name: "com.deepin.lib.Exportable1.Error.Err4",
		Body: []interface{}{"yyy err msg"},
	}

	if !reflect.DeepEqual(err, expectedErr) {
		t.Errorf("err expected %#v, but got %#v", expectedErr, err)
	}

}
