package pulse

/*
#include "dde-pulse.h"
#cgo pkg-config: libpulse glib-2.0 libpulse-mainloop-glib
*/
import "C"
import "fmt"
import "unsafe"

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

type Context struct {
	cbs map[int][]Callback

	ctx  *C.pa_context
	loop *C.pa_mainloop
}

func (c *Context) GetSinkList() (r []*Sink) {
	ck := newCookie()

	C.get_sink_info_list(c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		r = append(r, info.ToSink())
	}
	return
}

func (c *Context) GetSink(index uint32) *Sink {
	ck := newCookie()
	C.get_sink_info(c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	info := ck.Reply()
	if info == nil {
		return nil
	}
	return info.ToSink()
}

func (c *Context) GetSinkInputList() (r []*SinkInput) {
	ck := newCookie()

	C.get_sink_input_info_list(c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		r = append(r, info.ToSinkInput())
	}
	return
}

func (c *Context) GetSinkInput(index uint32) *SinkInput {
	ck := newCookie()
	C.get_sink_input_info(c.ctx, C.int64_t(ck.id), C.uint32_t(index))

	info := ck.Reply()
	if info == nil {
		return nil
	}
	return info.ToSinkInput()
}

func (c *Context) GetSourceList() (r []*Source) {
	ck := newCookie()

	C.get_source_info_list(c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		r = append(r, info.ToSource())
	}
	return
}

func (c *Context) GetSource(index uint32) *Source {
	ck := newCookie()
	C.get_source_info(c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	return ck.Reply().ToSource()
}

func (c *Context) GetServer() *Server {
	ck := newCookie()
	C.get_server_info(c.ctx, C.int64_t(ck.id))
	return ck.Reply().ToServer()
}

func (c *Context) GetSourceOutputList() (r []*SourceOutput) {
	ck := newCookie()

	C.get_source_output_info_list(c.ctx, C.int64_t(ck.id))
	for _, info := range ck.ReplyList() {
		r = append(r, info.ToSourceOutput())
	}
	return
}

func (c *Context) GetSourceOutput(index uint32) *SourceOutput {
	ck := newCookie()
	C.get_source_output_info(c.ctx, C.int64_t(ck.id), C.uint32_t(index))
	return ck.Reply().ToSourceOutput()
}

func (c *Context) GetDefaultSource() string {
	return ""
}
func (c *Context) GetDefaultSink() string {
	return ""
}
func (c *Context) SetDefaultSink(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.pa_context_set_default_sink(c.ctx, cname, C.success_cb, nil)
}
func (c *Context) SetDefaultSource(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	C.pa_context_set_default_source(c.ctx, cname, C.success_cb, nil)
}

var __context *Context

func GetContext() *Context {
	if __context == nil {
		loop := C.pa_mainloop_new()
		ctx := C.pa_init(loop)
		__context = &Context{
			cbs:  make(map[int][]Callback),
			ctx:  ctx,
			loop: loop,
		}
		go __context.runLoop()
	}
	return __context
}

func (c *Context) runLoop() {
	C.pa_mainloop_run(c.loop, nil)
}

//export receive_some_info
func receive_some_info(cookie int64, infoType int, info unsafe.Pointer, end bool) {
	c := fetchCookie(cookie)
	if c == nil {
		//NOTE: didn't know why the c will be nil, but it happened at sometimes
		fmt.Println("Warning: recieve_some_info with nil cookie", cookie, infoType, info, end)
		return
	}
	if end {
		c.EndOfList()
	} else {
		c.Feed(infoType, info)
	}
}

func (c *Context) ConnectPeekDetect(cb func(idx int, v float64)) {
}

func (c *Context) Connect(facility int, cb func(eventType int, idx uint32)) {
	// sink sinkinput source sourceoutput
	c.cbs[facility] = append(c.cbs[facility], cb)
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
