// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package simple

/*
#cgo pkg-config: libpulse-simple
#include <pulse/simple.h>
#include <pulse/error.h>
#include <stdlib.h>
*/
import "C"

import (
	"fmt"
	"unsafe"
)

type Conn struct {
	Ptr unsafe.Pointer
}

func wrapConn(ptr *C.pa_simple) Conn {
	return Conn{Ptr: unsafe.Pointer(ptr)}
}

func (v Conn) native() *C.pa_simple {
	return (*C.pa_simple)(v.Ptr)
}

type StreamDirection uint

const (
	StreamDirectionPlayback = C.PA_STREAM_PLAYBACK
)

type SampleFormat uint

const (
	SampleFormatU8        = C.PA_SAMPLE_U8
	SampleFormatS16NE     = C.PA_SAMPLE_S16NE
	SampleFormatS16LE     = C.PA_SAMPLE_S16LE
	SampleFormatS16BE     = C.PA_SAMPLE_S16BE
	SampleFormatS24LE     = C.PA_SAMPLE_S24LE
	SampleFormatS24BE     = C.PA_SAMPLE_S24BE
	SampleFormatFloat32LE = C.PA_SAMPLE_FLOAT32LE
	SampleFormatFloat32BE = C.PA_SAMPLE_FLOAT32BE
	SampleFormatS32LE     = C.PA_SAMPLE_S32LE
	SampleFormatS32BE     = C.PA_SAMPLE_S32BE
)

type SampleSpec struct {
	Format   SampleFormat
	Rate     uint32
	Channels uint8
}

func (ss *SampleSpec) native() *C.pa_sample_spec {
	if ss == nil {
		return nil
	}
	var ss0 C.pa_sample_spec
	ss0.format = C.pa_sample_format_t(ss.Format)
	ss0.rate = C.uint32_t(ss.Rate)
	ss0.channels = C.uint8_t(ss.Channels)
	return &ss0
}

// server: server name, empty for default
// name: A descriptive name for this client, application name
// dir: Open this stream for recording or playback?
// dev: sink or source name, empty for default
// streamName: A descriptive name for this stream (application name, song title, ...)
// sampleSpec: The sample type to use
func NewConn(server, name string, dir StreamDirection, dev, streamName string,
	sampleSpec *SampleSpec) (Conn, error) {
	var server0 *C.char
	if server != "" {
		server0 = C.CString(server) // or null
		defer C.free(unsafe.Pointer(server0))
	}

	name0 := C.CString(name)

	var dev0 *C.char
	if dev != "" {
		dev0 = C.CString(dev)
		defer C.free(unsafe.Pointer(dev0))
	}

	streamName0 := C.CString(streamName)

	var errCode C.int
	ret := C.pa_simple_new(server0, name0, C.pa_stream_direction_t(dir),
		dev0, streamName0, sampleSpec.native(), nil, nil, &errCode)

	// clean
	C.free(unsafe.Pointer(name0))
	C.free(unsafe.Pointer(streamName0))

	if ret == nil {
		return Conn{}, newError("pa_simple_new", errCode)
	}
	return wrapConn(ret), nil
}

func (c Conn) Write(data unsafe.Pointer, bytes uint) (int, error) {
	var errCode C.int
	ret := C.pa_simple_write(c.native(), data, C.size_t(bytes), &errCode)
	return int(ret), newError("pa_simple_write", errCode)
}

func (c Conn) Drain() (int, error) {
	var errCode C.int
	ret := C.pa_simple_drain(c.native(), &errCode)
	return int(ret), newError("pa_simple_drain", errCode)
}

// Close and free
func (c Conn) Free() {
	C.pa_simple_free(c.native())
}

func (c Conn) Flush() (int, error) {
	var errCode C.int
	ret := C.pa_simple_flush(c.native(), &errCode)
	return int(ret), newError("pa_simple_flush", errCode)
}

type Error struct {
	Code int
	Fn   string
}

func newError(fn string, errCode C.int) error {
	if errCode != 0 /*PA_OK*/ {
		return Error{
			Code: int(errCode),
			Fn:   fn,
		}
	}
	return nil
}

func (err Error) Error() string {
	errMsg := C.GoString(C.pa_strerror(C.int(err.Code)))
	return fmt.Sprintf("%s: %s", err.Fn, errMsg)
}
