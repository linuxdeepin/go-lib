// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pulse

/*
#include "dde-pulse.h"
#cgo pkg-config: libpulse glib-2.0
//enable below flags in test environment
//#cgo CFLAGS: -W -Wall -fsanitize=thread
//#cgo LDFLAGS: -ltsan
*/
import "C"
import (
	"fmt"
	"sync"
	"time"
	"unsafe"
)

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

type Event struct {
	Facility int
	Type     int
	Index    uint32
}

type Context struct {
	eventChanList []chan<- *Event
	stateChanList []chan<- int
	mu            sync.Mutex

	ctx  *C.pa_context
	loop *C.pa_threaded_mainloop
}

func (c *Context) GetCardList() (r []*Card) {
	ck := newCookie()

	c.safeDo(func() {
		C._get_card_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	})

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
	c.safeDo(func() {
		C._get_card_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	})
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

	c.safeDo(func() {
		C._get_sink_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	})
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

	c.safeDo(func() {
		C._get_sink_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	})
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

	c.safeDo(func() {
		C._get_sink_input_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	})
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

	c.safeDo(func() {
		C._get_sink_input_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	})

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

	c.safeDo(func() {
		C._get_source_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	})
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
	c.safeDo(func() {
		C._get_source_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	})

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
	c.safeDo(func() {
		C._get_server_info(c.loop, c.ctx, C.int64_t(ck.id))
	})

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

	c.safeDo(func() {
		C._get_source_output_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	})
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

	c.safeDo(func() {
		C._get_source_output_info(c.loop, c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	})
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
	c.safeDo(func() {
		cname := C.CString(name)
		op := C.pa_context_set_default_sink(c.ctx, cname, C.get_success_cb(), nil)
		C.pa_operation_unref(op)
		C.free(unsafe.Pointer(cname))
	})
}
func (c *Context) SetDefaultSource(name string) {
	c.safeDo(func() {
		cname := C.CString(name)
		op := C.pa_context_set_default_source(c.ctx, cname, C.get_success_cb(), nil)
		C.pa_operation_unref(op)
		C.free(unsafe.Pointer(cname))
	})
}

// MoveSinkInputsByName move sink-inputs to the special sink name
func (c *Context) MoveSinkInputsByName(sinkInputs []uint32, sinkName string) {
	c.safeDo(func() {
		cname := C.CString(sinkName)
		for _, idx := range sinkInputs {
			op := C.pa_context_move_sink_input_by_name(c.ctx, C.uint32_t(idx), cname,
				C.get_success_cb(), nil)
			C.pa_operation_unref(op)
		}
		C.free(unsafe.Pointer(cname))
	})
}

// MoveSinkInputsByIndex move sink-inputs to the special sink index
func (c *Context) MoveSinkInputsByIndex(sinkInputs []uint32, sinkIdx uint32) {
	c.safeDo(func() {
		for _, idx := range sinkInputs {
			op := C.pa_context_move_sink_input_by_index(c.ctx, C.uint32_t(idx), C.uint32_t(sinkIdx),
				C.get_success_cb(), nil)
			C.pa_operation_unref(op)
		}
	})
}

// MoveSourceOutputsByName move source-outputs to the special source name
func (c *Context) MoveSourceOutputsByName(sourceOutputs []uint32, sourceName string) {
	c.safeDo(func() {
		cname := C.CString(sourceName)
		for _, idx := range sourceOutputs {
			op := C.pa_context_move_source_output_by_name(
				c.ctx,
				C.uint32_t(idx),
				cname,
				C.get_success_cb(),
				nil,
			)
			C.pa_operation_unref(op)
		}
		C.free(unsafe.Pointer(cname))
	})
}

// MoveSourceOutputsByIndex move source-outputs to the special source index
func (c *Context) MoveSourceOutputsByIndex(sourceOutputs []uint32, sourceIdx uint32) {
	c.safeDo(func() {
		for _, idx := range sourceOutputs {
			op := C.pa_context_move_source_output_by_index(
				c.ctx,
				C.uint32_t(idx),
				C.uint32_t(sourceIdx),
				C.get_success_cb(),
				nil,
			)
			C.pa_operation_unref(op)
		}
	})
}

// GetModuleList retrieves the list of loaded modules
func (c *Context) GetModuleList() (r []*Module) {
	ck := newCookie()

	c.safeDo(func() {
		C._get_module_info_list(c.loop, c.ctx, C.int64_t(ck.id))
	})

	for _, info := range ck.ReplyList() {
		module := info.data.(*Module)
		if module == nil {
			continue
		}
		r = append(r, module)
	}
	return
}

var (
	__context   *Context
	__ctxLocker sync.Mutex
)

func GetContextForced() *Context {
	__ctxLocker.Lock()
	ctx := __context
	if ctx != nil {
		freeContext(ctx)
		__context = nil
	}
	__ctxLocker.Unlock()
	return GetContext()
}

var PulseInitTimeout = 15

func GetContext() *Context {
	__ctxLocker.Lock()
	defer __ctxLocker.Unlock()
	if __context == nil {
		loop := C.pa_threaded_mainloop_new()

		timer := time.AfterFunc(time.Second*time.Duration(PulseInitTimeout), func() {
			C.set_connect_timeout()
			C.pa_threaded_mainloop_signal(loop, 0)
		})
		defer timer.Stop()

		ctx := C.new_pa_context(loop)
		if ctx == nil {
			//TODO: Free loop
			fmt.Println("Failed to init pulseaudio")
			return nil
		}

		__context = &Context{
			ctx:  ctx,
			loop: loop,
		}
	}
	return __context
}

func (c *Context) SuspendSinkById(idx uint32, suspend int) {
	c.safeDo(func() {
		C._suspend_sink_by_id(c.loop, c.ctx, C.uint32_t(idx), C.int(suspend))
	})
}

func (c *Context) SuspendSourceById(idx uint32, suspend int) {
	c.safeDo(func() {
		C._suspend_source_by_id(c.loop, c.ctx, C.uint32_t(idx), C.int(suspend))
	})
}

func (c *Context) SetSinkPortByIndex(sinkIndex uint32, portName string) {
	c.safeDo(func() {
		cPortName := C.CString(portName)
		op := C.pa_context_set_sink_port_by_index(c.ctx, C.uint32_t(sinkIndex),
			cPortName, C.get_success_cb(), nil)
		C.free(unsafe.Pointer(cPortName))
		C.pa_operation_unref(op)
	})
}

func (c *Context) SetSinkMuteByIndex(sinkIndex uint32, mute bool) {
	muteInt := 0
	if mute {
		muteInt = 1
	}
	c.safeDo(func() {
		op := C.pa_context_set_sink_mute_by_index(c.ctx, C.uint32_t(sinkIndex),
			C.int(muteInt), C.get_success_cb(), nil)
		C.pa_operation_unref(op)
	})
}

func (c *Context) SetSinkVolumeByIndex(sinkIndex uint32, v CVolume) {
	c.safeDo(func() {
		op := C.pa_context_set_sink_volume_by_index(c.ctx, C.uint32_t(sinkIndex),
			&v.core, C.get_success_cb(), nil)
		C.pa_operation_unref(op)
	})
}

func (c *Context) SetSourcePortByIndex(sourceIndex uint32, portName string) {
	c.safeDo(func() {
		cPortName := C.CString(portName)
		op := C.pa_context_set_source_port_by_index(c.ctx, C.uint32_t(sourceIndex),
			cPortName, C.get_success_cb(), nil)
		C.free(unsafe.Pointer(cPortName))
		C.pa_operation_unref(op)
	})
}

func (c *Context) SetSourceVolumeByIndex(sourceIndex uint32, v CVolume) {
	c.safeDo(func() {
		op := C.pa_context_set_source_volume_by_index(c.ctx, C.uint32_t(sourceIndex),
			&v.core, C.get_success_cb(), nil)
		C.pa_operation_unref(op)
	})
}

func (c *Context) SetSourceMuteByIndex(sourceIndex uint32, mute bool) {
	muteInt := 0
	if mute {
		muteInt = 1
	}

	c.safeDo(func() {
		op := C.pa_context_set_source_mute_by_index(c.ctx, C.uint32_t(sourceIndex),
			C.int(muteInt), C.get_success_cb(), nil)
		C.pa_operation_unref(op)
	})
}

func (c *Context) SetSinkInputVolume(sinkInputIndex uint32, v CVolume) {
	c.safeDo(func() {
		op := C.pa_context_set_sink_input_volume(c.ctx, C.uint32_t(sinkInputIndex),
			&v.core, C.get_success_cb(), nil)
		C.pa_operation_unref(op)
	})
}

func (c *Context) SetSinkInputMute(sinkInputIndex uint32, mute bool) {
	muteInt := 0
	if mute {
		muteInt = 1
	}
	c.safeDo(func() {
		op := C.pa_context_set_sink_input_mute(c.ctx, C.uint32_t(sinkInputIndex),
			C.int(muteInt), C.get_success_cb(), nil)
		C.pa_operation_unref(op)
	})
}

func (c *Context) LoadModule(name, argument string) {
	c.safeDo(func() {
		cName := C.CString(name)
		defer C.free(unsafe.Pointer(cName))
		var cArgument *C.char
		if argument != "" {
			cArgument = C.CString(argument)
			defer C.free(unsafe.Pointer(cArgument))
		}
		op := C.pa_context_load_module(c.ctx, cName, cArgument, C.get_index_cb(), nil)
		C.pa_operation_unref(op)
	})
}
