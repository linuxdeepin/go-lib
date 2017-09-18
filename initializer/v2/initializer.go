/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
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

// Initializer is a chainable initializer. Init/InitOnSystemBus/InitOnSystemBus will accept a initializer,
// and then pass the successful return value to the next initializer. If error occurs, the rest initializers
// won't be executed any more. GetError is used to access the error.
type Initializer struct {
	e error
}

func (i *Initializer) init(fn func() error) *Initializer {
	if i.e != nil {
		return i
	}

	if err := fn(); err != nil {
		i.e = err
	}

	return i
}

// Do accepts a initializer function, stop other Do if any error occurs.
func (i *Initializer) Do(fn func() error) *Initializer {
	return i.init(func() error {
		err := fn()
		if err != nil {
			return err
		}
		return nil
	})
}

// GetError returns the first error of initializers.
func (i *Initializer) GetError() error {
	if i.e != nil {
		return i.e
	}
	return nil
}

// Do starts the initialization.
func Do(fn func() error) *Initializer {
	i := new(Initializer)
	return i.Do(fn)
}
