// +build 386 amd64

package dbus

import (
	"reflect"
)

func setupSignalHandler(c *Conn, v interface{}, path ObjectPath, iface string) {
	value := reflect.ValueOf(v).Elem()
	types := reflect.TypeOf(v).Elem()
	n := value.NumField()
	for i := 0; i < n; i++ {
		fn := value.Field(i)
		t := types.Field(i)
		if isExportedStructField(t) && fn.Type().Kind() == reflect.Func {
			fn.Set(reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
				panic("This style of emit dbus signal has been disabled. please using dbus.Emit instead")
			}))
		}
	}
}
