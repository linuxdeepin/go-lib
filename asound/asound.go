package asound

/*
#cgo pkg-config: alsa
#include <asoundlib.h>
#include <stdlib.h>
#include <errno.h>
*/
import "C"
import (
	"errors"
	"unsafe"
)

// Format is the type used for specifying sample formats.
type PCMFormat C.snd_pcm_format_t

// The range of sample formats supported by ALSA.
const (
	PCMFormatS8        = C.SND_PCM_FORMAT_S8
	PCMFormatU8        = C.SND_PCM_FORMAT_U8
	PCMFormatS16LE     = C.SND_PCM_FORMAT_S16_LE
	PCMFormatS16BE     = C.SND_PCM_FORMAT_S16_BE
	PCMFormatU16LE     = C.SND_PCM_FORMAT_U16_LE
	PCMFormatU16BE     = C.SND_PCM_FORMAT_U16_BE
	PCMFormatS24LE     = C.SND_PCM_FORMAT_S24_LE
	PCMFormatS24BE     = C.SND_PCM_FORMAT_S24_BE
	PCMFormatU24LE     = C.SND_PCM_FORMAT_U24_LE
	PCMFormatU24BE     = C.SND_PCM_FORMAT_U24_BE
	PCMFormatS32LE     = C.SND_PCM_FORMAT_S32_LE
	PCMFormatS32BE     = C.SND_PCM_FORMAT_S32_BE
	PCMFormatU32LE     = C.SND_PCM_FORMAT_U32_LE
	PCMFormatU32BE     = C.SND_PCM_FORMAT_U32_BE
	PCMFormatFloatLE   = C.SND_PCM_FORMAT_FLOAT_LE
	PCMFormatFloatBE   = C.SND_PCM_FORMAT_FLOAT_BE
	PCMFormatFloat64LE = C.SND_PCM_FORMAT_FLOAT64_LE
	PCMFormatFloat64BE = C.SND_PCM_FORMAT_FLOAT64_BE
)

// Size return bytes needed to store a quantity of PCM sample.
func (format PCMFormat) Size(samples uint) int {
	ret := C.snd_pcm_format_size(C.snd_pcm_format_t(format), C.size_t(samples))
	return int(ret)
}

func (format PCMFormat) Name() string {
	ret := C.snd_pcm_format_name(C.snd_pcm_format_t(format))
	return C.GoString(ret)
}

type PCMSubFormat C.snd_pcm_subformat_t

func (subformat PCMSubFormat) Name() string {
	ret := C.snd_pcm_subformat_name(C.snd_pcm_subformat_t(subformat))
	return C.GoString(ret)
}

func (subformat PCMSubFormat) Desc() string {
	ret := C.snd_pcm_subformat_description(C.snd_pcm_subformat_t(subformat))
	return C.GoString(ret)
}

type PCM struct {
	Ptr unsafe.Pointer
}

func (v PCM) native() *C.snd_pcm_t {
	return (*C.snd_pcm_t)(v.Ptr)
}

func wrapPCM(ptr *C.snd_pcm_t) PCM {
	return PCM{Ptr: unsafe.Pointer(ptr)}
}

type PCMType C.snd_pcm_type_t

func (typ PCMType) Name() string {
	ret := C.snd_pcm_type_name(C.snd_pcm_type_t(typ))
	return C.GoString(ret)
}

type PCMStream C.snd_pcm_stream_t

const (
	PCMStreamPlayback = C.SND_PCM_STREAM_PLAYBACK
	PCMStreamCapture  = C.SND_PCM_STREAM_CAPTURE
)

func (stream PCMStream) Name() string {
	ret := C.snd_pcm_stream_name(C.snd_pcm_stream_t(stream))
	return C.GoString(ret)
}

func OpenPCM(name string, stream PCMStream, mode int) (PCM, error) {
	var pcmp *C.snd_pcm_t
	name0 := C.CString(name)
	ret := C.snd_pcm_open(&pcmp, name0, C.snd_pcm_stream_t(stream), C.int(mode))
	C.free(unsafe.Pointer(name0))
	if ret == 0 {
		return wrapPCM(pcmp), nil
	}
	return PCM{}, newError("snd_pcm_open", ret)
}

func (pcm PCM) Name() string {
	ret := C.snd_pcm_name(pcm.native())
	return C.GoString(ret)
}

func (pcm PCM) Drain() error {
	ret := C.snd_pcm_drain(pcm.native())
	return newError("snd_pcm_drain", ret)
}

func (pcm PCM) Drop() error {
	ret := C.snd_pcm_drop(pcm.native())
	return newError("snd_pcm_drop", ret)
}

func (pcm PCM) Pause(enable bool) error {
	var val C.int
	if enable {
		val = 1
	}
	ret := C.snd_pcm_pause(pcm.native(), val)
	return newError("snd_pcm_pause", ret)
}

func (pcm PCM) Prepare() error {
	ret := C.snd_pcm_prepare(pcm.native())
	return newError("snd_pcm_prepare", ret)
}

func (pcm PCM) Close() error {
	ret := C.snd_pcm_close(pcm.native())
	return newError("snd_pcm_close", ret)
}

func (pcm PCM) HwFree() error {
	ret := C.snd_pcm_hw_free(pcm.native())
	return newError("snd_pcm_hw_free", ret)
}

func (pcm PCM) Start() error {
	ret := C.snd_pcm_start(pcm.native())
	return newError("snd_pcm_start", ret)
}

func (pcm PCM) Resume() error {
	ret := C.snd_pcm_resume(pcm.native())
	return newError("snd_pcm_resume", ret)
}

func (pcm PCM) Reset() error {
	ret := C.snd_pcm_reset(pcm.native())
	return newError("snd_pcm_reset", ret)
}

type PCMState C.snd_pcm_state_t

func (pcm PCM) State() PCMState {
	ret := C.snd_pcm_state(pcm.native())
	return PCMState(ret)
}

type PCMSFrames C.snd_pcm_sframes_t

var ErrUnderrun = errors.New("underrun")

func (pcm PCM) Writei(buffer unsafe.Pointer, size PCMUFrames) (PCMUFrames, error) {
	ret := C.snd_pcm_writei(pcm.native(), buffer, C.snd_pcm_uframes_t(size))
	if ret == -C.EPIPE {
		return 0, ErrUnderrun
	} else if ret < 0 {
		return 0, newError("snd_pcm_writei", C.int(ret))
	}
	return PCMUFrames(ret), nil
}

//snd_pcm_hw_params_t
type PCMHwParams struct {
	Ptr unsafe.Pointer
}

func wrapPCMHwParams(ptr *C.snd_pcm_hw_params_t) PCMHwParams {
	return PCMHwParams{Ptr: unsafe.Pointer(ptr)}
}

func (v PCMHwParams) native() *C.snd_pcm_hw_params_t {
	return (*C.snd_pcm_hw_params_t)(v.Ptr)
}

func NewPCMHwParams() (PCMHwParams, error) {
	var params *C.snd_pcm_hw_params_t
	ret := C.snd_pcm_hw_params_malloc(&params)
	if ret == 0 {
		return wrapPCMHwParams(params), nil
	}
	return PCMHwParams{}, newError("snd_pcm_hw_params_malloc", ret)
}

func (params PCMHwParams) Free() {
	C.snd_pcm_hw_params_free(params.native())
}

type PCMAccess C.snd_pcm_access_t

const (
	PCMAccessMMapInterleaved    = C.SND_PCM_ACCESS_MMAP_INTERLEAVED
	PCMAccessMMapNonInterleaved = C.SND_PCM_ACCESS_MMAP_NONINTERLEAVED
	PCMAccessMMapComplex        = C.SND_PCM_ACCESS_MMAP_COMPLEX
	PCMAccessRWInterleaved      = C.SND_PCM_ACCESS_RW_INTERLEAVED
	PCMAccessRWNonInterleaved   = C.SND_PCM_ACCESS_RW_NONINTERLEAVED
)

func (access PCMAccess) Name() string {
	ret := C.snd_pcm_access_name(C.snd_pcm_access_t(access))
	return C.GoString(ret)
}

func (pcm PCM) HwParamsAny(params PCMHwParams) {
	C.snd_pcm_hw_params_any(pcm.native(), params.native())
}

//snd_pcm_hw_params
func (pcm PCM) HwParams(params PCMHwParams) error {
	ret := C.snd_pcm_hw_params(pcm.native(), params.native())
	return newError("snd_pcm_hw_params", ret)
}

func (pcm PCM) HwParamsSetAccess(params PCMHwParams, access PCMAccess) error {
	ret := C.snd_pcm_hw_params_set_access(pcm.native(), params.native(), C.snd_pcm_access_t(access))
	return newError("snd_pcm_hw_params_set_access", ret)
}

func (pcm PCM) HwParamsSetFormat(params PCMHwParams, format PCMFormat) error {
	ret := C.snd_pcm_hw_params_set_format(pcm.native(), params.native(), C.snd_pcm_format_t(format))
	return newError("snd_pcm_hw_params_set_format", ret)
}

type PCMUFrames C.snd_pcm_uframes_t

func (pcm PCM) HwParamsSetBufferSize(params PCMHwParams, val PCMUFrames) error {
	ret := C.snd_pcm_hw_params_set_buffer_size(pcm.native(), params.native(), C.snd_pcm_uframes_t(val))
	return newError("snd_pcm_hw_params_set_buffer_size", ret)
}

func (pcm PCM) HwParamsSetChannels(params PCMHwParams, val uint) error {
	ret := C.snd_pcm_hw_params_set_channels(pcm.native(), params.native(), C.uint(val))
	return newError("snd_pcm_hw_params_set_channels", ret)
}

func (pcm PCM) HwParamsSetRate(params PCMHwParams, val uint, dir int) error {
	ret := C.snd_pcm_hw_params_set_rate(pcm.native(), params.native(), C.uint(val), C.int(dir))
	return newError("snd_pcm_hw_params_set_rate", ret)
}

func (pcm PCM) HwParamsSetPeriodSize(params PCMHwParams, val PCMUFrames, dir int) error {
	ret := C.snd_pcm_hw_params_set_period_size(pcm.native(), params.native(),
		C.snd_pcm_uframes_t(val), C.int(dir))
	return newError("snd_pcm_hw_params_set_period_size", ret)
}

func (pcm PCM) HwParamsSetPeriodSizeNear(params PCMHwParams, val PCMUFrames) (PCMUFrames, int, error) {
	cval := C.snd_pcm_uframes_t(val)
	var dir C.int
	ret := C.snd_pcm_hw_params_set_period_size_near(pcm.native(), params.native(),
		&cval, &dir)
	if ret == 0 {
		return PCMUFrames(cval), int(dir), nil
	}
	return 0, 0, newError("snd_pcm_hw_params_set_period_size_near", ret)
}

func (params PCMHwParams) GetPeriodSize() (PCMUFrames, int, error) {
	var val C.snd_pcm_uframes_t
	var dir C.int
	ret := C.snd_pcm_hw_params_get_period_size(params.native(), &val, &dir)
	if ret == 0 {
		return PCMUFrames(val), int(dir), nil
	}
	return 0, 0, newError("snd_pcm_hw_params_get_period_size", ret)
}

// time unit us microseconds
func (params PCMHwParams) GetPeriodTime() (uint, int, error) {
	var val C.uint
	var dir C.int
	ret := C.snd_pcm_hw_params_get_period_time(params.native(), &val, &dir)
	if ret == 0 {
		return uint(val), int(dir), nil
	}
	return 0, 0, newError("snd_pcm_hw_params_get_period_time", ret)
}

type Error struct {
	code C.int
	fn   string
}

func newError(fn string, code C.int) error {
	if code == 0 {
		return nil
	}
	return &Error{
		fn:   fn,
		code: code,
	}
}

func (err *Error) Error() string {
	cmsg := C.GoString(C.snd_strerror(err.code))
	return err.fn + ": " + cmsg
}
