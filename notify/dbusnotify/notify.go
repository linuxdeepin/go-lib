/*
 * Copyright (C) 2017 ~ 2018 Deepin Technology Co., Ltd.
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

package dbusnotify

import "pkg.deepin.io/lib/dbus"
import "pkg.deepin.io/lib/dbus/property"
import "reflect"
import "sync"
import "runtime"
import "fmt"
import "errors"

/*prevent compile error*/
var _ = fmt.Println
var _ = runtime.SetFinalizer
var _ = sync.NewCond
var _ = reflect.TypeOf
var _ = property.BaseObserver{}

type Notifier struct {
	Path     dbus.ObjectPath
	DestName string
	core     *dbus.Object

	signals       map[<-chan *dbus.Signal]struct{}
	signalsLocker sync.Mutex
}

func (obj *Notifier) _createSignalChan() <-chan *dbus.Signal {
	obj.signalsLocker.Lock()
	ch := getBus().Signal()
	obj.signals[ch] = struct{}{}
	obj.signalsLocker.Unlock()
	return ch
}
func (obj *Notifier) _deleteSignalChan(ch <-chan *dbus.Signal) {
	obj.signalsLocker.Lock()
	delete(obj.signals, ch)
	getBus().DetachSignal(ch)
	obj.signalsLocker.Unlock()
}
func DestroyNotifier(obj *Notifier) {
	obj.signalsLocker.Lock()
	defer obj.signalsLocker.Unlock()
	if obj.signals == nil {
		return
	}
	for ch, _ := range obj.signals {
		getBus().DetachSignal(ch)
	}
	obj.signals = nil

	runtime.SetFinalizer(obj, nil)

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.DBus.Properties',sender='" + obj.DestName + "',member='PropertiesChanged'")
	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='PropertiesChanged'")

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='NotificationClosed'")

	dbusRemoveMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='ActionInvoked'")

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

	obj := &Notifier{Path: path, DestName: destName, core: core, signals: make(map[<-chan *dbus.Signal]struct{})}

	dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='NotificationClosed'")

	dbusAddMatch("type='signal',path='" + string(obj.Path) + "',interface='org.freedesktop.Notifications',sender='" + obj.DestName + "',member='ActionInvoked'")

	runtime.SetFinalizer(obj, func(_obj *Notifier) { DestroyNotifier(_obj) })
	return obj, nil
}
