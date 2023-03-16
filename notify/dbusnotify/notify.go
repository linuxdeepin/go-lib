// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusnotify

import (
	"errors"
	"fmt"
	"reflect"
	"runtime"
	"sync"

	"github.com/godbus/dbus/v5"
)

/*prevent compile error*/
var _ = fmt.Println
var _ = runtime.SetFinalizer
var _ = sync.NewCond
var _ = reflect.TypeOf

type Notifier struct {
	Path     dbus.ObjectPath
	DestName string
	core     dbus.BusObject

	signals       map[chan *dbus.Signal]struct{}
	signalsLocker sync.Mutex
}

func (obj *Notifier) _createSignalChan() chan *dbus.Signal {
	obj.signalsLocker.Lock()
	ch := make(chan *dbus.Signal, 10)
	getBus().Signal(ch)
	obj.signals[ch] = struct{}{}
	obj.signalsLocker.Unlock()
	return ch
}
func (obj *Notifier) _deleteSignalChan(ch chan *dbus.Signal) {
	obj.signalsLocker.Lock()
	delete(obj.signals, ch)
	getBus().RemoveSignal(ch)
	obj.signalsLocker.Unlock()
}
func DestroyNotifier(obj *Notifier) {
	obj.signalsLocker.Lock()
	defer obj.signalsLocker.Unlock()
	if obj.signals == nil {
		return
	}
	for ch := range obj.signals {
		getBus().RemoveSignal(ch)
	}
	obj.signals = nil

	runtime.SetFinalizer(obj, nil)

	_ = dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.DBus.Properties',sender='" + obj.DestName + "',member='PropertiesChanged'")
	_ = dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='PropertiesChanged'")

	_ = dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='NotificationClosed'")

	_ = dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='ActionInvoked'")

}

func (obj *Notifier) CloseNotification(id uint32) (_err error) {
	_err = obj.core.Call("org.freedesktop.Notifications.CloseNotification", 0, id).Store()
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj *Notifier) GetCapabilities() (caps []string, _err error) {
	_err = obj.core.Call("org.freedesktop.Notifications.GetCapabilities", 0).Store(&caps)
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj *Notifier) GetServerInformation() (name string, vendor string, version string, spec_version string, _err error) {
	_err = obj.core.Call("org.freedesktop.Notifications.GetServerInformation", 0).Store(&name, &vendor, &version, &spec_version)
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj *Notifier) Notify(app_name string, id uint32, icon string, summary string, body string, actions []string, hints map[string]dbus.Variant, timeout int32) (id_ uint32, _err error) {
	_err = obj.core.Call("org.freedesktop.Notifications.Notify", 0, app_name, id, icon, summary, body, actions, hints, timeout).Store(&id_)
	if _err != nil {
		fmt.Println(_err)
	}
	return
}

func (obj *Notifier) ConnectNotificationClosed(callback func(id uint32, reason uint32)) func() {
	sigChan := obj._createSignalChan()
	go func() {
		for v := range sigChan {
			if v.Path != obj.Path || v.Name != "org.freedesktop.Notifications.NotificationClosed" || 2 != len(v.Body) {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*uint32)(nil)).Elem() {
				continue
			}
			if reflect.TypeOf(v.Body[1]) != reflect.TypeOf((*uint32)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(uint32), v.Body[1].(uint32))
		}
	}()
	return func() {
		obj._deleteSignalChan(sigChan)
	}
}

func (obj *Notifier) ConnectActionInvoked(callback func(id uint32, action_key string)) func() {
	sigChan := obj._createSignalChan()
	go func() {
		for v := range sigChan {
			if v.Path != obj.Path || v.Name != "org.freedesktop.Notifications.ActionInvoked" || 2 != len(v.Body) {
				continue
			}
			if reflect.TypeOf(v.Body[0]) != reflect.TypeOf((*uint32)(nil)).Elem() {
				continue
			}
			if reflect.TypeOf(v.Body[1]) != reflect.TypeOf((*string)(nil)).Elem() {
				continue
			}

			callback(v.Body[0].(uint32), v.Body[1].(string))
		}
	}()
	return func() {
		obj._deleteSignalChan(sigChan)
	}
}

func NewNotifier(destName string, path dbus.ObjectPath) (*Notifier, error) {
	if !path.IsValid() {
		return nil, errors.New("The path of '" + string(path) + "' is invalid.")
	}

	core := getBus().Object(destName, path)

	obj := &Notifier{Path: path, DestName: destName, core: core, signals: make(map[chan *dbus.Signal]struct{})}

	_ = dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='NotificationClosed'")

	_ = dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='ActionInvoked'")

	runtime.SetFinalizer(obj, func(_obj *Notifier) { DestroyNotifier(_obj) })
	return obj, nil
}
