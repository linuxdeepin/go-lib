// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pulse

//#include "dde-pulse.h"
import "C"

import (
	"fmt"
	"runtime"
	"unsafe"
)

// ONLY C file and below function can use C.pa_threaded_mainloop_lock/unlock

type safeDoCtx struct {
	fn   func()
	loop *C.pa_threaded_mainloop
}

var pendingSafeDo = make(chan safeDoCtx)

func startHandleSafeDo() {
	// Move all safeDo fn to here to avoid creating too many OS Thread.
	// Because GOMAXPROC only apply to go-runtime.
	for c := range pendingSafeDo {
		if c.fn != nil && c.loop != nil {
			runtime.LockOSThread()
			C.pa_threaded_mainloop_lock(c.loop)
			c.fn()
			C.pa_threaded_mainloop_unlock(c.loop)
			runtime.UnlockOSThread()
		}
	}
}

func init() {
	// TODO: encapsulate the logic to Context and adding Start/Stop logic.
	go startHandleSafeDo()
}

func freeContext(ctx *Context) {
	runtime.LockOSThread()
	C.pa_threaded_mainloop_lock(ctx.loop)
	// The pa_context_unref must be protected.
	// This operation will cancel all pending operations which will touch  mainloop object.
	C.pa_context_unref(ctx.ctx)
	C.pa_threaded_mainloop_unlock(ctx.loop)
	// There no need to call pa_threaded_mainloop_stop.
	C.pa_threaded_mainloop_free(ctx.loop)
	runtime.UnlockOSThread()

	ctx.loop = nil
	ctx.ctx = nil
}

// safeDo invoke an function with lock
func (c *Context) safeDo(fn func()) {
	pendingSafeDo <- safeDoCtx{fn, c.loop}
}

func (c *Context) AddEventChan(ch chan<- *Event) {
	c.mu.Lock()
	c.eventChanList = append(c.eventChanList, ch)
	c.mu.Unlock()
}

func (c *Context) RemoveEventChan(ch chan<- *Event) {
	c.mu.Lock()
	var newList []chan<- *Event
	for _, ch0 := range c.eventChanList {
		if ch0 != ch {
			newList = append(newList, ch0)
		}
	}
	c.eventChanList = newList
	c.mu.Unlock()
}

func (c *Context) AddStateChan(ch chan<- int) {
	c.mu.Lock()
	c.stateChanList = append(c.stateChanList, ch)
	c.mu.Unlock()
}

func (c *Context) RemoveStateChan(ch chan<- int) {
	c.mu.Lock()
	var newList []chan<- int
	for _, ch0 := range c.stateChanList {
		if ch0 != ch {
			newList = append(newList, ch0)
		}
	}
	c.stateChanList = newList
	c.mu.Unlock()
}

// ALL of below functions are invoked from pulse's mainloop thread.
// So
//   1. Don't hold go-runtime lock in current thread.
//   2. Don't hold pa_threaded_mainloop in current thread.
//
// NOTE:
// In pa_threaded_mainloop, _prepare_ and _dispatch_ is running
// under lock. Only phrase _poll_ is release the lock. So callback
// must not try to hold lock.

//export go_handle_changed
func go_handle_changed(facility int, event_type int, idx uint32) {

	event := &Event{
		Facility: facility,
		Type:     event_type,
		Index:    idx,
	}

	ctx := GetContext()
	ctx.mu.Lock()

	for _, eventChan := range ctx.eventChanList {
		select {
		case eventChan <- event:
		default:
		}
	}

	ctx.mu.Unlock()
}

//export go_handle_state_changed
func go_handle_state_changed(state int) {
	ctx := GetContext()
	ctx.mu.Lock()

	for _, stateChan := range ctx.stateChanList {
		select {
		case stateChan <- state:
		default:
		}
	}

	ctx.mu.Unlock()
}

//export go_update_volume_meter
func go_update_volume_meter(source_index uint32, sink_index uint32, v float64) {
	sourceMeterLock.RLock()
	cb, ok := sourceMeterCBs[source_index]
	sourceMeterLock.RUnlock()

	if ok && cb != nil {
		cb(v)
	}
}

//export go_receive_some_info
func go_receive_some_info(cookie int64, infoType int, info unsafe.Pointer, eol int) {
	ck := fetchCookie(cookie)
	if ck == nil {
		if eol == 0 {
			// Sometimes pulseaudio will send eol=1 to pa_context_get_*_by_index/name.
			fmt.Println("Warning: recieve_some_info with nil cookie", cookie, infoType, info)
		}
		return
	}

	switch {
	case eol == 1:
		ck.EndOfList()
	case eol == 0:
		// 1. the Feed will be blocked until the ck creator received.
		// 2. the creator of ck may be invoked under PendingCallback
		// so this must not be moved to PendingCallback.
		paInfo := NewPaInfo(info, infoType)
		ck.Feed(paInfo)
	case eol < 0:
		ck.Failed()
	}
}
