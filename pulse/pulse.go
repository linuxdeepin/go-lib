package pulse

/*
#include "dde-pulse.h"
#cgo pkg-config: libpulse glib-2.0 libpulse-mainloop-glib
*/
import "C"
import "unsafe"
import "fmt"

type Context struct {
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
	return ck.Reply().ToSink()
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
	return ck.Reply().ToSinkInput()
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

var __context *Context

func GetContext() *Context {
	if __context == nil {
		loop := C.pa_mainloop_new()
		ctx := C.pa_init(loop)
		__context = &Context{ctx, loop}
		go __context.runLoop()
	}
	return __context
}

func (c *Context) runLoop() {
	C.pa_mainloop_run(c.loop, nil)
}

func (*Context) Connect(c, t int, cb func()) {
}

//export receive_some_info
func receive_some_info(cookie int64, infoType int, info unsafe.Pointer, end bool) {
	c := fetchCookie(cookie)
	if end {
		c.EndOfList()
	} else {
		c.Feed(infoType, info)
	}
}

//export go_handle_changed
func go_handle_changed(facility int, event_type int, idx uint32) {
	switch facility {
	case C.PA_SUBSCRIPTION_EVENT_CARD:
		if event_type == C.PA_SUBSCRIPTION_EVENT_NEW {
			fmt.Printf("DEBUG card %d new\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_CHANGE {
			fmt.Printf("DEBUG card %d state changed\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_REMOVE {
			fmt.Printf("DEBUG card %d removed\n", idx)
		}
	case C.PA_SUBSCRIPTION_EVENT_SINK:
		if event_type == C.PA_SUBSCRIPTION_EVENT_NEW {
			fmt.Printf("DEBUG sink %d new\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_CHANGE {
			fmt.Printf("DEBUG sink %d state changed\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_REMOVE {
			fmt.Printf("DEBUG sink %d removed\n", idx)
		}
	case C.PA_SUBSCRIPTION_EVENT_SOURCE:
		if event_type == C.PA_SUBSCRIPTION_EVENT_NEW {
			fmt.Printf("DEBUG source %d new\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_CHANGE {
			fmt.Printf("DEBUG source %d changed\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_REMOVE {
			fmt.Printf("DEBUG source %d removed\n", idx)
		}
	case C.PA_SUBSCRIPTION_EVENT_SINK_INPUT:
		if event_type == C.PA_SUBSCRIPTION_EVENT_NEW {
			fmt.Printf("DEBUG sink input %d new\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_CHANGE {
			fmt.Printf("DEBUG sink input %d changed\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_REMOVE {
			fmt.Printf("DEBUG sink input %d removed\n", idx)
		}
	case C.PA_SUBSCRIPTION_EVENT_SOURCE_OUTPUT:
		if event_type == C.PA_SUBSCRIPTION_EVENT_NEW {
			fmt.Printf("DEBUG source output %d new\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_CHANGE {
			fmt.Printf("DEBUG source output %d changed\n", idx)
		} else if event_type == C.PA_SUBSCRIPTION_EVENT_REMOVE {
			fmt.Printf("DEBUG source output %d removed\n", idx)
		}
	case C.PA_SUBSCRIPTION_EVENT_SAMPLE_CACHE:
	}
}
