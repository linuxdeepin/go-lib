/*This file is auto generate by dlib/dbus/proxyer. Don't edit it*/
package logger

import "dlib/dbus"
import "dlib/dbus/property"
import "reflect"
import "sync"
import "runtime"
import "errors"
import "strings"
import "fmt"

/*prevent compile error*/
var _ = runtime.SetFinalizer
var _ = sync.NewCond
var _ = reflect.TypeOf
var _ = property.BaseObserver{}

var __conn *dbus.Conn = nil

func getBus() *dbus.Conn {
	if __conn == nil {
		var err error
		__conn, err = dbus.SystemBus()
		if err != nil {
			panic(err)
		}
	}
	return __conn
}

type Logapi struct {
	Path dbus.ObjectPath
	core *dbus.Object
}

func (obj Logapi) Debug(arg0 uint64, arg1 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Debug", 0, arg0, arg1).Store()
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj Logapi) Error(arg0 uint64, arg1 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Error", 0, arg0, arg1).Store()
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj Logapi) Fatal(arg0 uint64, arg1 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Fatal", 0, arg0, arg1).Store()
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj Logapi) Info(arg0 uint64, arg1 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Info", 0, arg0, arg1).Store()
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj Logapi) NewLogger(arg0 string) (arg1 uint64, _err error) {
	_err = obj.core.Call("com.deepin.api.Logger.NewLogger", 0, arg0).Store(&arg1)
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj Logapi) Warning(arg0 uint64, arg1 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Warning", 0, arg0, arg1).Store()
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func newLogapi(path dbus.ObjectPath) (*Logapi, error) {
	if !path.IsValid() {
		return nil, errors.New("The path of '" + string(path) + "' is invalid.")
	}

	core := getBus().Object("com.deepin.api.Logger", path)
	var v string
	core.Call("org.freedesktop.DBus.Introspectable.Introspect", 0).Store(&v)
	if strings.Index(v, "com.deepin.api.Logger") == -1 {
		return nil, errors.New("'" + string(path) + "' hasn't interface 'com.deepin.api.Logger'.")
	}

	obj := &Logapi{Path: path, core: core}

	return obj, nil
}
