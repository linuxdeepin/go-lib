/*This file is auto generate by pkg.linuxdeepin.com/dbus-generator. Don't edit it*/
package log

// TODO: remove

import "pkg.linuxdeepin.com/lib/dbus"
import "pkg.linuxdeepin.com/lib/dbus/property"
import "reflect"
import "sync"
import "runtime"
import "errors"
import "strings"
import golog "log"

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
			golog.Println("[INFO] dbus unavailable,", err)
		}
	}
	return __conn
}

type Logapi struct {
	Path dbus.ObjectPath
	core *dbus.Object
}

func (obj Logapi) Debug(arg0, arg1, arg2 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Debug", 0, arg0, arg1, arg2).Store()
	if _err != nil {
		golog.Println(_err)
	}
	return
}

func (obj Logapi) Error(arg0, arg1, arg2 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Error", 0, arg0, arg1, arg2).Store()
	if _err != nil {
		golog.Println(_err)
	}
	return
}

func (obj Logapi) Fatal(arg0, arg1, arg2 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Fatal", 0, arg0, arg1, arg2).Store()
	if _err != nil {
		golog.Println(_err)
	}
	return
}

func (obj Logapi) Info(arg0, arg1, arg2 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Info", 0, arg0, arg1, arg2).Store()
	if _err != nil {
		golog.Println(_err)
	}
	return
}

func (obj Logapi) NewLogger(arg0 string) (arg1 string, _err error) {
	_err = obj.core.Call("com.deepin.api.Logger.NewLogger", 0, arg0).Store(&arg1)
	if _err != nil {
		golog.Println(_err)
	}
	return
}

func (obj Logapi) Warning(arg0, arg1, arg2 string) (_err error) {
	_err = obj.core.Call("com.deepin.api.Logger.Warning", 0, arg0, arg1, arg2).Store()
	if _err != nil {
		golog.Println(_err)
	}
	return
}

func (obj Logapi) GetLog(arg0 string) (arg1 string, _err error) {
	_err = obj.core.Call("com.deepin.api.Logger.GetLog", 0, arg0).Store(&arg1)
	if _err != nil {
		golog.Println(_err)
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
