package dbusutil

import (
	"encoding/json"
	"fmt"

	"pkg.deepin.io/lib/dbus1"
)

func MakeError(v Exportable, name string, args ...interface{}) *dbus.Error {
	errName := v.GetDBusExportInfo().Interface + ".Error." + name
	msg := fmt.Sprint(args...)
	return &dbus.Error{
		Name: errName,
		Body: []interface{}{msg},
	}
}

func MakeErrorf(v Exportable, name, format string, args ...interface{}) *dbus.Error {
	errName := v.GetDBusExportInfo().Interface + ".Error." + name
	msg := fmt.Sprintf(format, args...)
	return &dbus.Error{
		Name: errName,
		Body: []interface{}{msg},
	}
}

func MakeErrorJSON(v Exportable, name string, detail interface{}) *dbus.Error {
	var msg string
	errName := v.GetDBusExportInfo().Interface + ".Error." + name
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
	err0, ok := err.(DBusError)
	var name, msg string
	if ok {
		name = err0.Name()
		msg = err0.Error()
	} else {
		name = "com.deepin.DBus.Error.Unnamed"
		msg = err.Error()
	}
	return &dbus.Error{
		Name: name,
		Body: []interface{}{msg},
	}
}
