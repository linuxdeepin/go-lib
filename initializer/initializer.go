/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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

// GetError returns the first error of initializers.
func (i *Initializer) GetError() error {
	return i.e
}
