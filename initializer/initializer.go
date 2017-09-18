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

package initializer

import (
	"pkg.deepin.io/lib/dbus"
)

// TODO:
// 1. a Promise like initializer might be better.
// 2. using reflect to support different initializer.

// Initializer is a chainable initializer. Init/InitOnSystemBus/InitOnSystemBus will accept a initializer,
// and then pass the successful return value to the next initializer. If error occurs, the rest initializers
// won't be executed any more. GetError is used to access the error.
type Initializer struct {
	v interface{}
	e error
}

// NewInitializer creates a new Initializer.
func NewInitializer() *Initializer {
	return new(Initializer)
}

func installOnSessionBus(v interface{}) (interface{}, error) {
	return v, dbus.InstallOnSession(v.(dbus.DBusObject))
}

func installOnSystemBus(v interface{}) (interface{}, error) {
	return v, dbus.InstallOnSystem(v.(dbus.DBusObject))
}

func noop(v interface{}) (interface{}, error) {
	return v, nil
}

func (i *Initializer) initWithHandler(fn func(interface{}) (interface{}, error), handler func(interface{}) (interface{}, error)) *Initializer {
	if i.e != nil {
		return i
	}

	var err error
	var v interface{}

	v, err = fn(i.v)
	if err != nil {
		i.e = err
		return i
	}

	v, err = handler(v)
	if err != nil {
		i.e = err
	}

	i.v = v

	return i
}

// Init accepts a initializer function, and pass the successful return value to next initializer.
func (i *Initializer) Init(fn func(interface{}) (interface{}, error)) *Initializer {
	return i.initWithHandler(fn, noop)
}

// InitOnSessionBus accepts a initializer function which must return a dbus.DBusObject
// as successful value, and then install this dbus.DBusObject on session bus.
func (i *Initializer) InitOnSessionBus(fn func(interface{}) (interface{}, error)) *Initializer {
	return i.initWithHandler(fn, installOnSessionBus)
}

// InitOnSystemBus accepts a initializer function which must return a dbus.DBusObject
// as successful value, and then install this dbus.DBusObject on system bus.
func (i *Initializer) InitOnSystemBus(fn func(interface{}) (interface{}, error)) *Initializer {
	return i.initWithHandler(fn, installOnSystemBus)
}

// GetError returns the first error of initializers.
func (i *Initializer) GetError() error {
	return i.e
}
