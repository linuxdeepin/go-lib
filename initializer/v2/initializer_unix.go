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

package initializer

import (
	"pkg.deepin.io/lib/dbus"
)

// DoWithSessionBus starts the initialization and install dbus object to session bus.
func (i *Initializer) DoWithSessionBus(fn func() (dbus.DBusObject, error)) *Initializer {
	return i.init(func() error {
		dbusObject, err := fn()
		if err != nil {
			return err
		}

		err = dbus.InstallOnSession(dbusObject)
		if err != nil {
			return err
		}
		return nil
	})
}

// DoWithSystemBus starts the initialization and install dbus object to system bus.
func (i *Initializer) DoWithSystemBus(fn func() (dbus.DBusObject, error)) *Initializer {
	return i.init(func() error {
		dbusObject, err := fn()
		if err != nil {
			return err
		}

		err = dbus.InstallOnSystem(dbusObject)
		if err != nil {
			return err
		}
		return nil
	})
}

// DoWithSessionBus starts the initialization and install dbus object to sesison bus.
func DoWithSessionBus(fn func() (dbus.DBusObject, error)) *Initializer {
	i := new(Initializer)
	return i.DoWithSessionBus(fn)
}

// DoWithSystemBus starts the initialization and install dbus object to system bus.
func DoWithSystemBus(fn func() (dbus.DBusObject, error)) *Initializer {
	i := new(Initializer)
	return i.DoWithSystemBus(fn)
}
