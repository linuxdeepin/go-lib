// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package stb_vorbis

/*
#include "stb_vorbis.h"
#include <stdlib.h>
#cgo LDFLAGS: -lm
*/
import "C"
import (
	"errors"
	"fmt"
	"reflect"
	"unsafe"
)

type Info struct {
	SampleRate uint
	Channels   int
}

type Decoder struct {
	Ptr unsafe.Pointer
}

func (v Decoder) native() *C.stb_vorbis {
	return (*C.stb_vorbis)(v.Ptr)
}

func wrapDecoder(ptr *C.stb_vorbis) Decoder {
	return Decoder{Ptr: unsafe.Pointer(ptr)}
}

func OpenFile(filename string) (Decoder, error) {
	filename0 := C.CString(filename)
	var errCode C.int
	ret := C.stb_vorbis_open_filename(filename0, &errCode, nil)
	C.free(unsafe.Pointer(filename0))

	if errCode != 0 {
		return Decoder{}, Error{Code: int(errCode)}
	}
	return wrapDecoder(ret), nil
}

func DecodeFile(filename string) (result []int16, channels int, sampleRate int, err error) {
	filename0 := C.CString(filename)
	var channels0 C.int
	var sampleRate0 C.int
	var output *C.short
	ret := C.stb_vorbis_decode_filename(filename0, &channels0, &sampleRate0, &output)
	C.free(unsafe.Pointer(filename0))
	if ret < 0 {
		return nil, 0, 0, errors.New("decode filename error")
	}

	var slice []C.short
	setSliceDataLen(unsafe.Pointer(&slice), unsafe.Pointer(output), int(ret*channels0))
	result = make([]int16, len(slice))
	for i, v := range slice {
		result[i] = int16(v)
	}
	return result, int(channels0), int(sampleRate0), nil
}

func setSliceDataLen(pSlice, pData unsafe.Pointer, len int) {
	header := (*reflect.SliceHeader)(pSlice)
	header.Cap = len
	header.Len = len
	header.Data = uintptr(pData)
}

func (d Decoder) Close() {
	C.stb_vorbis_close(d.native())
}

const (
	NoError             = C.VORBIS__no_error
	NeedMoreData        = C.VORBIS_need_more_data
	InvalidAPIMixing    = C.VORBIS_invalid_api_mixing
	OutOfMem            = C.VORBIS_outofmem
	FeatureNotSupported = C.VORBIS_feature_not_supported
	TooManyChannels     = C.VORBIS_too_many_channels
	FileOpenFailure     = C.VORBIS_file_open_failure
	SeekWithoutLength   = C.VORBIS_seek_without_length
	UnexpectedEOF       = C.VORBIS_unexpected_eof
	SeekInvalid         = C.VORBIS_seek_invalid

	// vorbis errors:
	InvalidSetup  = C.VORBIS_invalid_setup
	InvalidStream = C.VORBIS_invalid_stream

	// ogg errors:
	MissingCapturePattern         = C.VORBIS_missing_capture_pattern
	InvalidStreamStructureVersion = C.VORBIS_invalid_stream_structure_version
	InvalidContinuedPacketFlag    = C.VORBIS_continued_packet_flag_invalid
	IncorrectStreamSerialNumber   = C.VORBIS_incorrect_stream_serial_number
	InvalidFirstPage              = C.VORBIS_invalid_first_page
	BadPacketType                 = C.VORBIS_bad_packet_type
	CantFindLastPage              = C.VORBIS_cant_find_last_page
	SeekFailed                    = C.VORBIS_seek_failed
)

type Error struct {
	Code int
}

func (err Error) Error() string {
	switch err.Code {
	case NeedMoreData:
		return "need more data"
	case FeatureNotSupported:
		return "feature not supported"
	case FileOpenFailure:
		return "file open failure"
	case InvalidAPIMixing:
		return "invalid API mixing"
	case OutOfMem:
		return "out of memory"
	case SeekInvalid:
		return "seek invalid"
	case SeekWithoutLength:
		return "seek without length"
	case TooManyChannels:
		return "too many channels"
	case UnexpectedEOF:
		return "unexpected EOF"

	case InvalidSetup:
		return "invalid setup"
	case InvalidStream:
		return "invalid stream"

	case MissingCapturePattern:
		return "ogg missing capture pattern"
	case InvalidStreamStructureVersion:
		return "ogg invalid stream structure version"
	case InvalidContinuedPacketFlag:
		return "ogg invalid continued packet flag"
	case IncorrectStreamSerialNumber:
		return "ogg incorrect stream serial number"
	case InvalidFirstPage:
		return "ogg invalid first page"
	case BadPacketType:
		return "ogg bad packet type"
	case CantFindLastPage:
		return "ogg can't find last page"
	case SeekFailed:
		return "ogg seek failed"
	}
	return fmt.Sprintf("other vorbis error(%d)", err.Code)
}

func (d Decoder) GetError() error {
	ret := C.stb_vorbis_get_error(d.native())
	if ret == NoError {
		return nil
	}
	return Error{
		Code: int(ret),
	}
}

func (d Decoder) GetInfo() Info {
	ret := C.stb_vorbis_get_info(d.native())
	return Info{
		SampleRate: uint(ret.sample_rate),
		Channels:   int(ret.channels),
	}
}

func (d Decoder) StreamLengthInSamples() uint {
	ret := C.stb_vorbis_stream_length_in_samples(d.native())
	return uint(ret)
}

func (d Decoder) GetSamplesShortInterleaved(channels int, buffer []byte) int {
	buf := (*C.short)(unsafe.Pointer(&buffer[0]))
	ret := C.stb_vorbis_get_samples_short_interleaved(d.native(), C.int(channels), buf, C.int(len(buffer)/2))
	return int(ret)
}
