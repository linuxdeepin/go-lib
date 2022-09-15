// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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
