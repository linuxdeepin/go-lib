// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

// fork from https://github.com/pocke/goevent
package event

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

type Event struct {
	listeners []reflect.Value
	lmu       sync.RWMutex
	argTypes  []reflect.Type
}

// New
func New(f interface{}) *Event {
	fn := reflect.ValueOf(f)
	if fn.Kind() != reflect.Func {
		panic("not a function")
	}
	fnType := fn.Type()
	fnNum := fnType.NumIn()
	argTypes := make([]reflect.Type, fnNum)
	for i := 0; i < fnNum; i++ {
		argTypes[i] = fnType.In(i)
	}
	return &Event{
		argTypes: argTypes,
	}
}

// On
func (e *Event) On(f interface{}) error {
	fn, err := e.checkFuncSignature(f)
	if err != nil {
		return err
	}

	e.lmu.Lock()
	e.listeners = append(e.listeners, *fn)
	e.lmu.Unlock()
	return nil
}

// Off
func (e *Event) Off(f interface{}) error {
	fn := reflect.ValueOf(f)

	e.lmu.Lock()
	l := len(e.listeners)
	m := l // for error check
	for i := 0; i < l; i++ {
		if fn == e.listeners[i] {
			// XXX: GC Ref: http://jxck.hatenablog.com/entry/golang-slice-internals
			e.listeners = append(e.listeners[:i], e.listeners[i+1:]...)
			l--
			i--
		}
	}
	e.lmu.Unlock()

	if l == m {
		return fmt.Errorf("Listener does't exists")
	}
	return nil
}

// Trigger
func (e *Event) Trigger(args ...interface{}) error {
	// check args len
	if len(e.argTypes) != len(args) {
		return fmt.Errorf("argument length expected %d, but got %d", len(e.argTypes), len(args))
	}
	// check each arg type
	for i, t := range e.argTypes {
		t0 := reflect.TypeOf(args[i])
		if t != t0 {
			return fmt.Errorf("args[%d] type expected %s, but got %s", i, t, t0)
		}
	}

	arguments := make([]reflect.Value, len(args))
	for i, v := range args {
		arguments[i] = reflect.ValueOf(v)
	}

	e.lmu.RLock()
	defer e.lmu.RUnlock()

	wg := sync.WaitGroup{}
	wg.Add(len(e.listeners))
	for _, fn := range e.listeners {
		go func(f reflect.Value) {
			defer wg.Done()
			f.Call(arguments)
		}(fn)
	}
	wg.Wait()
	return nil
}

func (e *Event) checkFuncSignature(f interface{}) (*reflect.Value, error) {
	fn := reflect.ValueOf(f)
	if fn.Kind() != reflect.Func {
		return nil, errors.New("not a function")
	}
	fnType := fn.Type()
	fnNum := fnType.NumIn()
	// check args len
	if len(e.argTypes) != fnNum {
		return nil, fmt.Errorf("argument length expected %d, but got %d", len(e.argTypes), fnNum)
	}

	// check each arg type
	for i, t := range e.argTypes {
		t0 := fnType.In(i)
		if t != t0 {
			return nil, fmt.Errorf("args[%d] type expected %s, but got %s", i, t, t0)
		}
	}
	return &fn, nil
}
