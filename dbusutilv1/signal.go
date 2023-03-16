// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
// 
// SPDX-License-Identifier: GPL-3.0-or-later
package dbusutilv1

import (
	"sync"

	"github.com/godbus/dbus/v5"
)

type SignalHandlerId int

type SignalLoop struct {
	conn          *dbus.Conn
	bufSize       int
	ch            chan *dbus.Signal
	mu            sync.Mutex
	handlers      map[SignalHandlerId]SignalHandler
	nextHandlerId SignalHandlerId
}

func NewSignalLoop(conn *dbus.Conn, bufSize int) *SignalLoop {
	pl := &SignalLoop{
		conn:          conn,
		bufSize:       bufSize,
		handlers:      make(map[SignalHandlerId]SignalHandler),
		nextHandlerId: 1,
	}
	return pl
}

func (sl *SignalLoop) Conn() *dbus.Conn {
	return sl.conn
}

func (sl *SignalLoop) Start() {
	ch := make(chan *dbus.Signal, sl.bufSize)
	sl.conn.Signal(ch)
	sl.ch = ch

	go func() {
		for signal := range ch {
			sl.process(signal)
		}
	}()
}

func (sl *SignalLoop) process(signal *dbus.Signal) {
	var cbs []SignalHandlerFunc
	sl.mu.Lock()
	for _, handler := range sl.handlers {
		if handler.rule.match(signal) {
			if handler.cb != nil {
				cbs = append(cbs, handler.cb)
			}
		}
	}
	sl.mu.Unlock()

	for _, cb := range cbs {
		cb(signal)
	}
}

func (sl *SignalLoop) Stop() {
	sl.conn.RemoveSignal(sl.ch)
	close(sl.ch)
}

func (sl *SignalLoop) AddHandler(rule *SignalRule, cb SignalHandlerFunc) SignalHandlerId {
	sl.mu.Lock()
	defer sl.mu.Unlock()

	id := sl.nextHandlerId
	sl.nextHandlerId++

	sl.handlers[id] = SignalHandler{
		rule: rule,
		cb:   cb,
	}
	return id
}

func (sl *SignalLoop) RemoveHandler(id SignalHandlerId) {
	if id < 0 {
		return
	}
	sl.mu.Lock()
	delete(sl.handlers, id)
	sl.mu.Unlock()
}

type SignalHandler struct {
	rule *SignalRule
	cb   SignalHandlerFunc
}

type SignalRule struct {
	Sender string
	Path   dbus.ObjectPath
	Name   string
}

func (m *SignalRule) match(sig *dbus.Signal) bool {
	if m == nil {
		return true
	}

	if m.Sender != "" && m.Sender != sig.Sender {
		// define sender, but not equal
		return false
	}

	if m.Path != "" && m.Path != sig.Path {
		// define path, but not equal
		return false
	}

	if m.Name != "" && m.Name != sig.Name {
		// define name, but not equal
		return false
	}
	return true
}

type SignalHandlerFunc func(sig *dbus.Signal)
