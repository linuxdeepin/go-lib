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

package pulse

/*
#include "dde-pulse.h"
#cgo pkg-config: libpulse glib-2.0
//enable below flags in test environment
//#cgo CFLAGS: -fsanitize=thread
//#cgo LDFLAGS: -ltsan
*/
import "C"
import "fmt"
import "unsafe"
import "runtime"
import "sync"
import "time"

type Callback func(eventType int, idx uint32)

const (
	EventTypeNew    = C.PA_SUBSCRIPTION_EVENT_NEW
	EventTypeChange = C.PA_SUBSCRIPTION_EVENT_CHANGE
	EventTypeRemove = C.PA_SUBSCRIPTION_EVENT_REMOVE
)
const (
	FacilityServer       = C.PA_SUBSCRIPTION_EVENT_SERVER
	FacilitySink         = C.PA_SUBSCRIPTION_EVENT_SINK
	FacilitySource       = C.PA_SUBSCRIPTION_EVENT_SOURCE
	FacilitySinkInput    = C.PA_SUBSCRIPTION_EVENT_SINK_INPUT
	FacilitySourceOutput = C.PA_SUBSCRIPTION_EVENT_SOURCE_OUTPUT
	FacilityCard         = C.PA_SUBSCRIPTION_EVENT_CARD
	FacilityClient       = C.PA_SUBSCRIPTION_EVENT_CLIENT
	FacilityModule       = C.PA_SUBSCRIPTION_EVENT_MODULE
	FacilitySampleCache  = C.PA_SUBSCRIPTION_EVENT_SAMPLE_CACHE
)

const (
	ContextStateUnconnected = C.PA_CONTEXT_UNCONNECTED
	ContextStateConnecting  = C.PA_CONTEXT_CONNECTING
	ContextStateAuthorizing = C.PA_CONTEXT_AUTHORIZING
	ContextStateSettingName = C.PA_CONTEXT_SETTING_NAME
	ContextStateReady       = C.PA_CONTEXT_READY
	ContextStateFailed      = C.PA_CONTEXT_FAILED
	ContextStateTerminated  = C.PA_CONTEXT_TERMINATED
)

type Context struct {
	cbs      map[int][]Callback
	stateCbs map[int][]func()

	ctx  *C.pa_context
	loop *C.pa_threaded_mainloop
}

func (c *Context) GetCardList() (r []*Card) {
	ck := newCookie()

	C.get_card_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		card := info.ToCard()
		if card == nil {
			continue
		}
		r = append(r, card)
	}
	return
}

func (c *Context) GetCard(index uint32) (*Card, error) {
	ck := newCookie()
	C.get_card_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	info := ck.Reply()
	if info == nil {
		return nil, fmt.Errorf("Can't obtain this instance for: %v", index)
	}

	card := info.ToCard()
	if card == nil {
		return nil, fmt.Errorf("'%d' not a valid card index", index)
	}
	return card, nil
}

func (c *Context) GetSinkList() (r []*Sink) {
	ck := newCookie()

	C.get_sink_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		sink := info.ToSink()
		if sink == nil {
			continue
		}
		r = append(r, sink)
	}
	return
}

func (c *Context) GetSink(index uint32) (*Sink, error) {
	ck := newCookie()
	C.get_sink_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	info := ck.Reply()
	if info == nil {
		return nil, fmt.Errorf("Can't obtain this instance for: %v", index)
	}

	sink := info.ToSink()
	if sink == nil {
		return nil, fmt.Errorf("'%d' not a valid sink index", index)
	}
	return sink, nil
}

func (c *Context) GetSinkInputList() (r []*SinkInput) {
	ck := newCookie()

	C.get_sink_input_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		si := info.ToSinkInput()
		if si == nil {
			continue
		}
		r = append(r, si)
	}
	return
}

func (c *Context) GetSinkInput(index uint32) (*SinkInput, error) {
	ck := newCookie()
	C.get_sink_input_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))

	info := ck.Reply()
	if info == nil {
		return nil, fmt.Errorf("Can't obtain this instance for: %v", index)
	}

	si := info.ToSinkInput()
	if si == nil {
		return nil, fmt.Errorf("'%d' not a valid sinkinput index", index)
	}
	return si, nil
}

func (c *Context) GetSourceList() (r []*Source) {
	ck := newCookie()

	C.get_source_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		source := info.ToSource()
		if source == nil {
			continue
		}
		r = append(r, source)
	}
	return
}

func (c *Context) GetSource(index uint32) (*Source, error) {
	ck := newCookie()
	C.get_source_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))

	info := ck.Reply()
	if info == nil {
		return nil, fmt.Errorf("Can't obtain this instance for: %v", index)
	}

	source := info.ToSource()
	if source == nil {
		return nil, fmt.Errorf("'%d' not a valid source index", index)
	}
	return source, nil
}

func (c *Context) GetServer() (*Server, error) {
	ck := newCookie()
	C.get_server_info(c.loop, c.ctx, C.int64_t(ck.id))

	info := ck.Reply()
	if info == nil {
		return nil, fmt.Errorf("Can't obtain the server instance.")
	}

	s := info.ToServer()
	if s == nil {
		return nil, fmt.Errorf("Not found valid server")
	}
	return s, nil
}

func (c *Context) GetSourceOutputList() (r []*SourceOutput) {
	ck := newCookie()

	C.get_source_output_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		so := info.ToSourceOutput()
		if so == nil {
			continue
		}
		r = append(r, so)
	}
	return
}

func (c *Context) GetSourceOutput(index uint32) (*SourceOutput, error) {
	ck := newCookie()
	C.get_source_output_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	info := ck.Reply()
	if info == nil {
		return nil, fmt.Errorf("Can't obtain the this instance for: %v", index)
	}

	so := info.ToSourceOutput()
	if so == nil {
		return nil, fmt.Errorf("'%d' not a valid sourceoutput index", index)
	}
	return so, nil
}

func (c *Context) GetDefaultSource() string {
	s, _ := c.GetServer()
	if s != nil {
		return s.DefaultSourceName
	}
	return ""
}
func (c *Context) GetDefaultSink() string {
	s, _ := c.GetServer()
	if s != nil {
		return s.DefaultSinkName
	}
	return ""
}
func (c *Context) SetDefaultSink(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	c.safeDo(func() {
		C.pa_context_set_default_sink(c.ctx, cname, C.get_success_cb(), nil)
	})
}
func (c *Context) SetDefaultSource(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	c.safeDo(func() {
		C.pa_context_set_default_source(c.ctx, cname, C.get_success_cb(), nil)
	})
}

// MoveSinkInputsByName move sink-inputs to the special sink name
func (c *Context) MoveSinkInputsByName(sinkInputs []uint32, sinkName string) {
	cname := C.CString(sinkName)
	defer C.free(unsafe.Pointer(cname))

	c.safeDo(func() {
		for _, idx := range sinkInputs {
			C.pa_context_move_sink_input_by_name(c.ctx, C.uint32_t(idx), cname,
				C.get_success_cb(), nil)
		}
	})
}

// MoveSinkInputsByIndex move sink-inputs to the special sink index
func (c *Context) MoveSinkInputsByIndex(sinkInputs []uint32, sinkIdx uint32) {
	c.safeDo(func() {
		for _, idx := range sinkInputs {
			C.pa_context_move_sink_input_by_index(c.ctx, C.uint32_t(idx), C.uint32_t(sinkIdx),
				C.get_success_cb(), nil)
		}
	})
}

// MoveSourceOutputsByName move source-outputs to the special source name
func (c *Context) MoveSourceOutputsByName(sourceOutputs []uint32, sourceName string) {
	cname := C.CString(sourceName)
	defer C.free(unsafe.Pointer(cname))

	c.safeDo(func() {
		for _, idx := range sourceOutputs {
			C.pa_context_move_source_output_by_name(c.ctx, C.uint32_t(idx), cname,
				C.get_success_cb(), nil)
		}
	})
}

// MoveSourceOutputsByIndex move source-outputs to the special source index
func (c *Context) MoveSourceOutputsByIndex(sourceOutputs []uint32, sourceIdx uint32) {
	c.safeDo(func() {
		for _, idx := range sourceOutputs {
			C.pa_context_move_source_output_by_index(c.ctx, C.uint32_t(idx), C.uint32_t(sourceIdx),
				C.get_success_cb(), nil)
		}
	})
}

// safeDo invoke an function with lock
func (c *Context) safeDo(fn func()) {
	runtime.LockOSThread()
	C.pa_threaded_mainloop_lock(c.loop)
	fn()
	C.pa_threaded_mainloop_unlock(c.loop)
	runtime.UnlockOSThread()
}

var (
	__context   *Context
	__ctxLocker sync.Mutex
)

// PulseInitTimeout set pa_init timeout, if timeout return nil
var PulseInitTimeout = time.Duration(20)

func GetContextForced() *Context {
	__ctxLocker.Lock()
	ctx := __context
	if ctx != nil {
		C.pa_threaded_mainloop_stop(ctx.loop)
		C.pa_context_unref(ctx.ctx)
		C.pa_threaded_mainloop_free(ctx.loop)
		ctx.loop = nil
		ctx.ctx = nil
		__context = nil
	}
	__ctxLocker.Unlock()
	return GetContext()
}

func GetContext() *Context {
	__ctxLocker.Lock()
	defer __ctxLocker.Unlock()
	if __context == nil {
		loop := C.pa_threaded_mainloop_new()
		time.AfterFunc(time.Second*PulseInitTimeout, func() {
			fmt.Println("Connect to pulseaudio timeout, terminate")
			C.pa_finalize()
		})
		ctx := C.new_pa_context(loop)
		if ctx == nil {
			fmt.Println("Failed to init pulseaudio")
			return nil
		}

		__context = &Context{
			cbs:      make(map[int][]Callback),
			stateCbs: make(map[int][]func()),
			ctx:      ctx,
			loop:     loop,
		}
	}
	return __context
}

//export receive_some_info
func receive_some_info(cookie int64, infoType int, info unsafe.Pointer, status int) {
	c := fetchCookie(cookie)
	if c == nil {
		fmt.Println("Warning: recieve_some_info with nil cookie", cookie, infoType, info, status)
		return
	}

	switch {
	case status == 1:
		c.EndOfList()
	case status == 0:
		c.Feed(infoType, info)
	case status < 0:
		c.Failed()
	}
}

func (c *Context) Connect(facility int, cb func(eventType int, idx uint32)) {
	// sink sinkinput source sourceoutput
	c.cbs[facility] = append(c.cbs[facility], cb)
}

func (c *Context) ConnectStateChanged(state int, cb func()) {
	c.stateCbs[state] = append(c.stateCbs[state], cb)
}

func (c *Context) SuspendSinkById(idx uint32, suspend int) {
	C.suspend_sink_by_id(c.ctx, C.uint32_t(idx), C.int(suspend))
}

func (c *Context) SuspendSourceById(idx uint32, suspend int) {
	C.suspend_source_by_id(c.ctx, C.uint32_t(idx), C.int(suspend))
}

func (c *Context) handlePAEvent(facility, eventType int, idx uint32) {
	if cb, ok := c.cbs[facility]; ok {
		for _, c := range cb {
			go c(eventType, idx)
		}
	} else {
		fmt.Println("unknow event", facility, eventType, idx)
	}
}

//export go_handle_changed
func go_handle_changed(facility int, event_type int, idx uint32) {
	GetContext().handlePAEvent(facility, event_type, idx)
}

//export go_handle_state_changed
func go_handle_state_changed(state int) {
	cbs, ok := GetContext().stateCbs[state]
	if !ok {
		fmt.Println("Unregiste state:", state)
	}

	for _, cb := range cbs {
		go cb()
	}
}
