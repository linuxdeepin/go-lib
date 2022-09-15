// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package tasker

import (
	"fmt"
	"reflect"
	"sync"
	"time"
)

type DelayTask struct {
	timer    *time.Timer
	handler  reflect.Value
	argsType []reflect.Type

	duration time.Duration

	termination chan struct{}
	locker      sync.Mutex
}

func NewDelayTask(duration time.Duration,
	cb interface{}) (*DelayTask, error) {
	handler, argsType, err := handleFuncSignature(cb)
	if err != nil {
		return nil, err
	}
	return &DelayTask{
		handler:  *handler,
		argsType: argsType,
		duration: duration,
	}, nil
}

func (task *DelayTask) Start(args ...interface{}) error {
	// if started, stop it, else nothing to do
	task.Stop()

	task.locker.Lock()
	defer task.locker.Unlock()
	if len(task.argsType) != len(args) {
		return fmt.Errorf("argument length expected %d, but got %d",
			len(task.argsType), len(args))
	}
	// check each arg type
	for i, t := range task.argsType {
		t0 := reflect.TypeOf(args[i])
		if t != t0 {
			return fmt.Errorf("args[%d] type expected %s, but got %s",
				i, t, t0)
		}
	}
	var values = make([]reflect.Value, len(args))
	for i, v := range args {
		values[i] = reflect.ValueOf(v)
	}

	task.doStart(values)
	return nil
}

func (task *DelayTask) doStart(values []reflect.Value) {
	task.termination = make(chan struct{})
	task.timer = time.NewTimer(task.duration)
	go func() {
		exit := task.termination
		timer := task.timer
		select {
		case <-exit:
			return
		case <-timer.C:
			task.handler.Call(values)
			return
		}
	}()
}

func (task *DelayTask) Stop() {
	task.locker.Lock()
	defer task.locker.Unlock()
	if task.termination != nil {
		close(task.termination)
		task.termination = nil
		// TODO: delete
		time.Sleep(time.Millisecond * 50)
	}
	if task.timer != nil {
		task.timer.Stop()
	}
}

func handleFuncSignature(f interface{}) (*reflect.Value,
	[]reflect.Type, error) {
	fn := reflect.ValueOf(f)
	if fn.Kind() != reflect.Func {
		return nil, nil, fmt.Errorf("not a function")
	}

	fnType := fn.Type()
	fnNum := fnType.NumIn()
	argTypes := make([]reflect.Type, fnNum)
	for i := 0; i < fnNum; i++ {
		argTypes[i] = fnType.In(i)
	}

	return &fn, argTypes, nil
}
