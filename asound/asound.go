// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package asound

/*
#cgo pkg-config: alsa
#include <alsa/asoundlib.h>
#include <stdlib.h>
#include <errno.h>
*/
import "C"
import (
	"errors"
	"fmt"
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
	PCMFormatS24_3LE   = C.SND_PCM_FORMAT_S24_3LE
	PCMFormatS24_3BE   = C.SND_PCM_FORMAT_S24_3BE
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

// PCM sample subformat
type PCMSubformat C.snd_pcm_subformat_t

func (subformat PCMSubformat) Name() string {
	ret := C.snd_pcm_subformat_name(C.snd_pcm_subformat_t(subformat))
	return C.GoString(ret)
}

func (subformat PCMSubformat) Desc() string {
	ret := C.snd_pcm_subformat_description(C.snd_pcm_subformat_t(subformat))
	return C.GoString(ret)
}

// Internal structure for an output object.
//
// The ALSA library uses a pointer to this structure as a handle to an output object. Applications don't access its contents directly.
type Output struct {
	Ptr unsafe.Pointer
}

func (v Output) native() *C.snd_output_t {
	return (*C.snd_output_t)(v.Ptr)
}

// Creates a new output object writing to a file.
func OpenOutput(file, mode string) (Output, error) {
	file0 := C.CString(file)
	mode0 := C.CString(mode)
	var outputPtr *C.snd_output_t
	ret := C.snd_output_stdio_open(&outputPtr, file0, mode0)
	C.free(unsafe.Pointer(file0))
	C.free(unsafe.Pointer(mode0))
	if ret == 0 {
		return Output{unsafe.Pointer(outputPtr)}, nil
	}
	return Output{}, newError("snd_output_stdio_open", ret)
}

// Flushes an output handle (like fflush(3)).
func (o Output) Flush() error {
	ret := C.snd_output_flush(o.native())
	return newError("snd_output_flush", ret)
}

// Writes a string to an output handle (like fputs(3)).
func (o Output) WriteString(str string) error {
	str0 := C.CString(str)
	ret := C.snd_output_puts(o.native(), str0)
	C.free(unsafe.Pointer(str0))
	return newError("snd_output_puts", ret)
}

func (o Output) Println(a ...interface{}) error {
	str := fmt.Sprintln(a...)
	return o.WriteString(str)
}

func (o Output) Printf(format string, a ...interface{}) error {
	str := fmt.Sprintf(format, a...)
	return o.WriteString(str)
}

// Closes an output handle.
func (o Output) Close() error {
	ret := C.snd_output_close(o.native())
	return newError("snd_output_close", ret)
}

// PCM handle
type PCM struct {
	Ptr unsafe.Pointer
}

func (v PCM) native() *C.snd_pcm_t {
	return (*C.snd_pcm_t)(v.Ptr)
}

func wrapPCM(ptr *C.snd_pcm_t) PCM {
	return PCM{Ptr: unsafe.Pointer(ptr)}
}

// PCM type
type PCMType C.snd_pcm_type_t

func (typ PCMType) Name() string {
	ret := C.snd_pcm_type_name(C.snd_pcm_type_t(typ))
	return C.GoString(ret)
}

// PCM Stream
type PCMStream C.snd_pcm_stream_t

const (
	PCMStreamPlayback = C.SND_PCM_STREAM_PLAYBACK
	PCMStreamCapture  = C.SND_PCM_STREAM_CAPTURE
)

func (stream PCMStream) Name() string {
	ret := C.snd_pcm_stream_name(C.snd_pcm_stream_t(stream))
	return C.GoString(ret)
}

// Opens a PCM.
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

// get identifier of PCM handle
func (pcm PCM) Name() string {
	ret := C.snd_pcm_name(pcm.native())
	return C.GoString(ret)
}

// get type of PCM handle
func (pcm PCM) Type() PCMType {
	typ := C.snd_pcm_type(pcm.native())
	return PCMType(typ)
}

// get stream for a PCM handle
func (pcm PCM) Stream() PCMStream {
	stream := C.snd_pcm_stream(pcm.native())
	return PCMStream(stream)
}

// Obtain general (static) information for PCM handle.
func (pcm PCM) Info(info PCMInfo) error {
	ret := C.snd_pcm_info(pcm.native(), info.native())
	return newError("snd_pcm_info", ret)
}

// Stop a PCM preserving pending frames.
//
// For playback wait for all pending frames to be played and then stop the PCM. For capture stop PCM permitting to retrieve residual frames.
func (pcm PCM) Drain() error {
	ret := C.snd_pcm_drain(pcm.native())
	return newError("snd_pcm_drain", ret)
}

// Stop a PCM dropping pending frames.
//
// This function stops the PCM immediately. The pending samples on the buffer are ignored.
func (pcm PCM) Drop() error {
	ret := C.snd_pcm_drop(pcm.native())
	return newError("snd_pcm_drop", ret)
}

// Pause/resume PCM.
func (pcm PCM) Pause(enable bool) error {
	var val C.int
	if enable {
		val = 1
	}
	ret := C.snd_pcm_pause(pcm.native(), val)
	return newError("snd_pcm_pause", ret)
}

// Prepare PCM for use.
func (pcm PCM) Prepare() error {
	ret := C.snd_pcm_prepare(pcm.native())
	return newError("snd_pcm_prepare", ret)
}

// close PCM handle
//
// Closes the specified PCM handle and frees all associated resources.
func (pcm PCM) Close() error {
	ret := C.snd_pcm_close(pcm.native())
	return newError("snd_pcm_close", ret)
}

// Recover the stream state from an error or suspend.
//
// This a high-level helper function building on other functions.
//
//This functions handles -EINTR (interrupted system call), -EPIPE (overrun or underrun) and -ESTRPIPE (stream is suspended) error codes trying to prepare given stream for next I/O.
//
//Note that this function returs the original error code when it is not handled inside this function (for example -EAGAIN is returned back).
func (pcm PCM) Recover(err int, silent bool) error {
	var silent0 C.int
	if silent {
		silent0 = 1
	}
	ret := C.snd_pcm_recover(pcm.native(), C.int(err), silent0)
	return newError("snd_pcm_recover", ret)
}

// Remove PCM hardware configuration and free associated resources.
func (pcm PCM) HwFree() error {
	ret := C.snd_pcm_hw_free(pcm.native())
	return newError("snd_pcm_hw_free", ret)
}

// Start a PCM.
func (pcm PCM) Start() error {
	ret := C.snd_pcm_start(pcm.native())
	return newError("snd_pcm_start", ret)
}

// Resume from suspend, no samples are lost.
//
// This function can be used when the stream is in the suspend state to do the fine resume from this state. Not all hardware supports this feature, when an -ENOSYS error is returned, use the PCM.Prepare() function to recovery.
func (pcm PCM) Resume() error {
	ret := C.snd_pcm_resume(pcm.native())
	return newError("snd_pcm_resume", ret)
}

// Reset PCM position.
//
// Reduce PCM delay to 0.
func (pcm PCM) Reset() error {
	ret := C.snd_pcm_reset(pcm.native())
	return newError("snd_pcm_reset", ret)
}

func (pcm PCM) DumpHwSetup(out Output) error {
	ret := C.snd_pcm_dump_hw_setup(pcm.native(), out.native())
	return newError("snd_pcm_dump_hw_setup", ret)
}

func (pcm PCM) DumpSwSetup(out Output) error {
	ret := C.snd_pcm_dump_sw_setup(pcm.native(), out.native())
	return newError("snd_pcm_dump_sw_setup", ret)
}

func (pcm PCM) DumpSetup(out Output) error {
	ret := C.snd_pcm_dump_setup(pcm.native(), out.native())
	return newError("snd_pcm_dump_setup", ret)
}

func (pcm PCM) Dump(out Output) error {
	ret := C.snd_pcm_dump(pcm.native(), out.native())
	return newError("snd_pcm_dump", ret)
}

// PCM state
type PCMState C.snd_pcm_state_t

// Return PCM state.
//
// This is a faster way to obtain only the PCM state without calling PCM.Status().
func (pcm PCM) State() PCMState {
	ret := C.snd_pcm_state(pcm.native())
	return PCMState(ret)
}

// Signed frames quantity
type PCMSFrames C.snd_pcm_sframes_t

var ErrUnderrun = errors.New("underrun")

// Write interleaved frames to a PCM.
//
// If the blocking behaviour is selected and it is running, then routine waits until all requested frames are played or put to the playback ring buffer. The returned number of frames can be less only if a signal or underrun occurred.
//
//If the non-blocking behaviour is selected, then routine doesn't wait at all.
func (pcm PCM) Writei(buffer unsafe.Pointer, size PCMUFrames) (PCMUFrames, error) {
	ret := C.snd_pcm_writei(pcm.native(), buffer, C.snd_pcm_uframes_t(size))
	if ret == -C.EPIPE {
		return 0, ErrUnderrun
	} else if ret < 0 {
		return 0, newError("snd_pcm_writei", C.int(ret))
	}
	return PCMUFrames(ret), nil
}

// PCM hardware configuration space container
//
// snd_pcm_hw_params_t is an opaque structure which contains a set of possible PCM hardware configurations. For example, a given instance might include a range of buffer sizes, a range of period sizes, and a set of several sample formats. Some subset of all possible combinations these sets may be valid, but not necessarily any combination will be valid.
//
// When a parameter is set or restricted using a snd_pcm_hw_params_set* function, all of the other ranges will be updated to exclude as many impossible configurations as possible. Attempting to set a parameter outside of its acceptable range will result in the function failing and an error code being returned.
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

// Dump a PCM hardware configuration space.
func (params PCMHwParams) Dump(out Output) error {
	ret := C.snd_pcm_hw_params_dump(params.native(), out.native())
	return newError("snd_pcm_hw_params_dump", ret)
}

// PCM access type
type PCMAccess C.snd_pcm_access_t

const (
	PCMAccessMMapInterleaved    = C.SND_PCM_ACCESS_MMAP_INTERLEAVED
	PCMAccessMMapNonInterleaved = C.SND_PCM_ACCESS_MMAP_NONINTERLEAVED
	PCMAccessMMapComplex        = C.SND_PCM_ACCESS_MMAP_COMPLEX
	PCMAccessRWInterleaved      = C.SND_PCM_ACCESS_RW_INTERLEAVED
	PCMAccessRWNonInterleaved   = C.SND_PCM_ACCESS_RW_NONINTERLEAVED
)

// get name of PCM access type
func (access PCMAccess) Name() string {
	ret := C.snd_pcm_access_name(C.snd_pcm_access_t(access))
	return C.GoString(ret)
}

// Fill params with a full configuration space for a PCM.
//
// The configuration space will be filled with all possible ranges for the PCM device.
func (pcm PCM) HwParamsAny(params PCMHwParams) {
	C.snd_pcm_hw_params_any(pcm.native(), params.native())
}

// Install one PCM hardware configuration chosen from a configuration space and snd_pcm_prepare it.
//
// The configuration is chosen fixing single parameters in this order: first access, first format, first subformat, min channels, min rate, min period time, max buffer size, min tick time. If no mutually compatible set of parameters can be chosen, a negative error code will be returned.
//
//After this call, snd_pcm_prepare() is called automatically and the stream is brought to SND_PCM_STATE_PREPARED state.
//
//The hardware parameters cannot be changed when the stream is running (active). The software parameters can be changed at any time.
//
//The configuration space will be updated to reflect the chosen parameters.
func (pcm PCM) HwParams(params PCMHwParams) error {
	ret := C.snd_pcm_hw_params(pcm.native(), params.native())
	return newError("snd_pcm_hw_params", ret)
}

// Restrict a configuration space to contain only one access type.
func (pcm PCM) HwParamsSetAccess(params PCMHwParams, access PCMAccess) error {
	ret := C.snd_pcm_hw_params_set_access(pcm.native(), params.native(), C.snd_pcm_access_t(access))
	return newError("snd_pcm_hw_params_set_access", ret)
}

// Restrict a configuration space to contain only one format.
func (pcm PCM) HwParamsSetFormat(params PCMHwParams, format PCMFormat) error {
	ret := C.snd_pcm_hw_params_set_format(pcm.native(), params.native(), C.snd_pcm_format_t(format))
	return newError("snd_pcm_hw_params_set_format", ret)
}

// Unsigned frames quantity
type PCMUFrames C.snd_pcm_uframes_t

// Restrict a configuration space to contain only one buffer size.
func (pcm PCM) HwParamsSetBufferSize(params PCMHwParams, val PCMUFrames) error {
	ret := C.snd_pcm_hw_params_set_buffer_size(pcm.native(), params.native(), C.snd_pcm_uframes_t(val))
	return newError("snd_pcm_hw_params_set_buffer_size", ret)
}

// Extract maximum buffer time from a configuration space.
// val	approximate maximum buffer duration in us
func (params PCMHwParams) GetBufferTimeMax() (val uint, dir int, err error) {
	var val0 C.uint
	var dir0 C.int
	ret := C.snd_pcm_hw_params_get_buffer_time_max(params.native(), &val0, &dir0)
	err = newError("snd_pcm_hw_params_get_buffer_time_max", ret)
	if err != nil {
		return
	}
	val = uint(val0)
	dir = int(dir0)
	return
}

// Restrict a configuration space to contain only one channels count.
func (pcm PCM) HwParamsSetChannels(params PCMHwParams, val uint) error {
	ret := C.snd_pcm_hw_params_set_channels(pcm.native(), params.native(), C.uint(val))
	return newError("snd_pcm_hw_params_set_channels", ret)
}

// Restrict a configuration space to contain only one rate.
func (pcm PCM) HwParamsSetRate(params PCMHwParams, val uint, dir int) error {
	ret := C.snd_pcm_hw_params_set_rate(pcm.native(), params.native(), C.uint(val), C.int(dir))
	return newError("snd_pcm_hw_params_set_rate", ret)
}

// Restrict a configuration space to have rate nearest to a target.
func (pcm PCM) HwParamsSetRateNear(params PCMHwParams, val *uint) (int, error) {
	cVal := C.uint(*val)
	var dir C.int
	ret := C.snd_pcm_hw_params_set_rate_near(pcm.native(), params.native(),
		&cVal, &dir)
	if ret == 0 {
		*val = uint(cVal)
		return int(dir), nil
	}
	return 0, newError("snd_pcm_hw_params_set_rate_near", ret)
}

// Restrict a configuration space to have period time nearest to a target.
func (pcm PCM) HwParamsSetPeriodTimeNear(params PCMHwParams, val *uint) (int, error) {
	val0 := C.uint(*val)
	var dir0 C.int
	ret := C.snd_pcm_hw_params_set_period_time_near(pcm.native(), params.native(), &val0, &dir0)
	if ret == 0 {
		*val = uint(val0)
		return int(dir0), nil
	}
	return 0, newError("snd_pcm_hw_params_set_period_time_near", ret)
}

// Restrict a configuration space to have buffer time nearest to a target.
func (pcm PCM) HwParamsSetBufferTimeNear(params PCMHwParams, val *uint) (int, error) {
	val0 := C.uint(*val)
	var dir0 C.int
	ret := C.snd_pcm_hw_params_set_buffer_time_near(pcm.native(), params.native(), &val0, &dir0)
	if ret == 0 {
		*val = uint(val0)
		return int(dir0), nil
	}
	return 0, newError("snd_pcm_hw_params_set_buffer_time_near", ret)
}

// Restrict a configuration space to contain only one period size.
func (pcm PCM) HwParamsSetPeriodSize(params PCMHwParams, val PCMUFrames, dir int) error {
	ret := C.snd_pcm_hw_params_set_period_size(pcm.native(), params.native(),
		C.snd_pcm_uframes_t(val), C.int(dir))
	return newError("snd_pcm_hw_params_set_period_size", ret)
}

// Restrict a configuration space to have period size nearest to a target.
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

// Extract period size from a configuration space.
func (params PCMHwParams) GetPeriodSize() (PCMUFrames, int, error) {
	var val C.snd_pcm_uframes_t
	var dir C.int
	ret := C.snd_pcm_hw_params_get_period_size(params.native(), &val, &dir)
	if ret == 0 {
		return PCMUFrames(val), int(dir), nil
	}
	return 0, 0, newError("snd_pcm_hw_params_get_period_size", ret)
}

// Extract period time from a configuration space. time unit us microseconds
func (params PCMHwParams) GetPeriodTime() (uint, int, error) {
	var val C.uint
	var dir C.int
	ret := C.snd_pcm_hw_params_get_period_time(params.native(), &val, &dir)
	if ret == 0 {
		return uint(val), int(dir), nil
	}
	return 0, 0, newError("snd_pcm_hw_params_get_period_time", ret)
}

const sizeOfPointer = unsafe.Sizeof(unsafe.Pointer(nil))

// Get a set of device name hints.
func GetDeviceNameHints(card int, iface string) (DeviceNameHints, error) {
	iface0 := C.CString(iface)
	var hints *unsafe.Pointer
	ret := C.snd_device_name_hint(C.int(card), iface0, &hints)
	C.free(unsafe.Pointer(iface0))

	if ret == 0 {
		return DeviceNameHints{hints}, nil
	}
	return DeviceNameHints{}, newError("snd_device_name_hint", ret)
}

type DeviceNameHints struct {
	// it pointer to a array of unsafe.Pointer, type in C is void**
	Ptr *unsafe.Pointer
}

func (v DeviceNameHints) Free() {
	C.snd_device_name_free_hint(v.Ptr)
}

func (v DeviceNameHints) Iter() *DeviceNameHintsIter {
	return &DeviceNameHintsIter{
		hints: v.Ptr,
	}
}

type DeviceNameHintsIter struct {
	hints, n *unsafe.Pointer
}

func (iter *DeviceNameHintsIter) Next() bool {
	if iter.n == nil {
		// init iter.n
		iter.n = iter.hints
	} else {
		// same as iter.n++ in C
		iter.n = (*unsafe.Pointer)(unsafe.Pointer(
			uintptr(unsafe.Pointer(iter.n)) + sizeOfPointer,
		))
	}
	return *iter.n != nil
}

func (iter DeviceNameHintsIter) Get(id string) string {
	id0 := C.CString(id)
	ret0 := C.snd_device_name_get_hint(*iter.n, id0)
	ret := C.GoString(ret0)
	C.free(unsafe.Pointer(ret0))
	return ret
}

// Try to determine the next card.
//
// Tries to determine the next card from given card number. If card number is -1, then the first available card is returned. If the result card number is -1, no more cards are available.
func CardNext(rcard *int) error {
	rcard0 := C.int(*rcard)
	ret := C.snd_card_next(&rcard0)
	*rcard = int(rcard0)
	return newError("snd_card_next", ret)
}

// Try to load the driver for a card.
//
// return true if driver is present, false if driver is not present
func CardLoad(card int) bool {
	ret := C.snd_card_load(C.int(card))
	return ret == 1
}

// CTL handle
type CTL struct {
	Ptr unsafe.Pointer
}

func wrapCTL(p *C.snd_ctl_t) CTL {
	return CTL{
		unsafe.Pointer(p),
	}
}

func (v CTL) native() *C.snd_ctl_t {
	return (*C.snd_ctl_t)(v.Ptr)
}

// Opens a CTL.
func CTLOpen(name string, mode int) (CTL, error) {
	name0 := C.CString(name)
	var ctlp *C.snd_ctl_t
	ret := C.snd_ctl_open(&ctlp, name0, C.int(mode))
	C.free(unsafe.Pointer(name0))

	if ret == 0 {
		return wrapCTL(ctlp), nil
	}
	return CTL{}, newError("snd_ctl_open", ret)
}

// close CTL handle
//
// Closes the specified CTL handle and frees all associated resources.
func (ctl CTL) Close() error {
	ret := C.snd_ctl_close(ctl.native())
	return newError("snd_ctl_close", ret)
}

// Get card related information.
func (ctl CTL) CardInfo(info CTLCardInfo) error {
	ret := C.snd_ctl_card_info(ctl.native(), info.native())
	return newError("snd_ctl_card_info", ret)
}

// PCM generic info container
type PCMInfo struct {
	Ptr unsafe.Pointer
}

func wrapPCMInfo(ptr *C.snd_pcm_info_t) PCMInfo {
	return PCMInfo{
		Ptr: unsafe.Pointer(ptr),
	}
}

func NewPCMInfo() (PCMInfo, error) {
	var p *C.snd_pcm_info_t
	ret := C.snd_pcm_info_malloc(&p)
	if ret == 0 {
		return wrapPCMInfo(p), nil
	}
	return PCMInfo{}, newError("snd_pcm_info_malloc", ret)
}

func (v PCMInfo) native() *C.snd_pcm_info_t {
	return (*C.snd_pcm_info_t)(v.Ptr)
}

func (info PCMInfo) Free() {
	C.snd_pcm_info_free(info.native())
}

//Set wanted device inside a PCM info container
func (info PCMInfo) SetDevice(val uint) {
	C.snd_pcm_info_set_device(info.native(), C.uint(val))
}

// Set wanted stream inside a PCM info container
func (info PCMInfo) SetStream(val PCMStream) {
	C.snd_pcm_info_set_stream(info.native(), C.snd_pcm_stream_t(val))
}

//Set wanted sub device inside a PCM info container
func (info PCMInfo) SetSubdevice(val uint) {
	C.snd_pcm_info_set_subdevice(info.native(), C.uint(val))
}

// Get card from a PCM info container.
func (info PCMInfo) GetCard() int {
	ret := C.snd_pcm_info_get_card(info.native())
	return int(ret)
}

// PCM class
type PCMClass C.snd_pcm_class_t

// Get class from a PCM info container.
func (info PCMInfo) GetClass() PCMClass {
	ret := C.snd_pcm_info_get_class(info.native())
	return PCMClass(ret)
}

// Get id from a PCM info container.
func (info PCMInfo) GetID() string {
	ret := C.snd_pcm_info_get_id(info.native())
	return C.GoString(ret)
}

// Get name from a PCM info container.
func (info PCMInfo) GetName() string {
	ret := C.snd_pcm_info_get_name(info.native())
	return C.GoString(ret)
}

// Get stream (direction) from a PCM info container.
func (info PCMInfo) GetStream() PCMStream {
	ret := C.snd_pcm_info_get_stream(info.native())
	return PCMStream(ret)
}

// PCM subclass
type PCMSubclass C.snd_pcm_subclass_t

// Get subclass from a PCM info container.
func (info PCMInfo) GetSubclass() PCMSubclass {
	ret := C.snd_pcm_info_get_subclass(info.native())
	return PCMSubclass(ret)
}

// Get subdevice from a PCM info container.
func (info PCMInfo) GetSubdevice() uint {
	ret := C.snd_pcm_info_get_subdevice(info.native())
	return uint(ret)
}

// Get subdevice name from a PCM info container.
func (info PCMInfo) GetSubdeviceName() string {
	ret := C.snd_pcm_info_get_subdevice_name(info.native())
	return C.GoString(ret)
}

// Get available subdevices count from a PCM info container.
func (info PCMInfo) GetSubdevicesAvail() uint {
	ret := C.snd_pcm_info_get_subdevices_avail(info.native())
	return uint(ret)
}

// Get subdevices count from a PCM info container.
func (info PCMInfo) GetSubdevicesCount() uint {
	ret := C.snd_pcm_info_get_subdevices_count(info.native())
	return uint(ret)
}

// Get next PCM device number.
func (ctl CTL) PCMNextDevice(device *int) error {
	device0 := C.int(*device)
	ret := C.snd_ctl_pcm_next_device(ctl.native(), &device0)
	*device = int(device0)
	return newError("snd_ctl_pcm_next_device", ret)
}

// Get info about a PCM device.
func (ctl CTL) PCMInfo(info PCMInfo) error {
	ret := C.snd_ctl_pcm_info(ctl.native(), info.native())
	return newError("snd_ctl_pcm_info", ret)
}

// CTL card info container
type CTLCardInfo struct {
	Ptr unsafe.Pointer
}

func (v CTLCardInfo) native() *C.snd_ctl_card_info_t {
	return (*C.snd_ctl_card_info_t)(v.Ptr)
}

func wrapCTLCardInfo(ptr *C.snd_ctl_card_info_t) CTLCardInfo {
	return CTLCardInfo{
		Ptr: unsafe.Pointer(ptr),
	}
}

func NewCTLCardInfo() (CTLCardInfo, error) {
	var p *C.snd_ctl_card_info_t
	ret := C.snd_ctl_card_info_malloc(&p)
	if ret == 0 {
		return wrapCTLCardInfo(p), nil
	}
	return CTLCardInfo{}, newError("snd_ctl_card_info_malloc", ret)
}

func (info CTLCardInfo) Free() {
	C.snd_ctl_card_info_free(info.native())
}

// Get card identifier from a CTL card info.
func (info CTLCardInfo) GetID() string {
	ret := C.snd_ctl_card_info_get_id(info.native())
	return C.GoString(ret)
}

// Get card number from a CTL card info.
func (info CTLCardInfo) GetCard() int {
	ret := C.snd_ctl_card_info_get_card(info.native())
	return int(ret)
}

// Get card component list from a CTL card info.
func (info CTLCardInfo) GetComponents() string {
	ret := C.snd_ctl_card_info_get_components(info.native())
	return C.GoString(ret)
}

// Get card driver name from a CTL card info.
func (info CTLCardInfo) GetDriver() string {
	ret := C.snd_ctl_card_info_get_driver(info.native())
	return C.GoString(ret)
}

// Get card long name from a CTL card info.
func (info CTLCardInfo) GetLongName() string {
	ret := C.snd_ctl_card_info_get_longname(info.native())
	return C.GoString(ret)
}

// Get card mixer name from a CTL card info.
func (info CTLCardInfo) GetMixerName() string {
	ret := C.snd_ctl_card_info_get_mixername(info.native())
	return C.GoString(ret)
}

// Get card name from a CTL card info.
func (info CTLCardInfo) GetName() string {
	ret := C.snd_ctl_card_info_get_name(info.native())
	return C.GoString(ret)
}

// Mixer Interface

type Mixer struct {
	Ptr unsafe.Pointer
}

func (m Mixer) native() *C.snd_mixer_t {
	return (*C.snd_mixer_t)(m.Ptr)
}

func wrapMixer(ptr *C.snd_mixer_t) Mixer {
	return Mixer{
		Ptr: unsafe.Pointer(ptr),
	}
}

// Opens an empty mixer.
func OpenMixer(mode int) (Mixer, error) {
	var mixer *C.snd_mixer_t
	ret := C.snd_mixer_open(&mixer, C.int(mode))
	if ret == 0 {
		return wrapMixer(mixer), nil
	}
	return Mixer{}, newError("snd_mixer_open", ret)
}

// Load a mixer elements.
func (m Mixer) Load() error {
	ret := C.snd_mixer_load(m.native())
	return newError("snd_mixer_load", ret)
}

// Close a mixer and free all related resources.
func (m Mixer) Close() error {
	ret := C.snd_mixer_close(m.native())
	return newError("snd_mixer_close", ret)
}

// Unload all mixer elements and free all related resources.
func (m Mixer) Free() {
	C.snd_mixer_free(m.native())
}

// Attach an HCTL specified with the CTL device name to an opened mixer.
func (m Mixer) Attach(name string) error {
	cName := C.CString(name)
	ret := C.snd_mixer_attach(m.native(), cName)
	C.free(unsafe.Pointer(cName))
	return newError("snd_mixer_attach", ret)
}

// get first element for a mixer
func (m Mixer) FirstElem() MixerElem {
	ret := C.snd_mixer_first_elem(m.native())
	return wrapMixerElem(ret)
}

// Find a mixer simple element.
func (m Mixer) FindSelem(id MixerSelemId) MixerElem {
	ret := C.snd_mixer_find_selem(m.native(), id.native())
	return wrapMixerElem(ret)
}

type MixerElem struct {
	Ptr unsafe.Pointer
}

func wrapMixerElem(ptr *C.snd_mixer_elem_t) MixerElem {
	return MixerElem{
		Ptr: unsafe.Pointer(ptr),
	}
}

func (e MixerElem) native() *C.snd_mixer_elem_t {
	return (*C.snd_mixer_elem_t)(e.Ptr)
}

// get next mixer element
func (e MixerElem) Next() MixerElem {
	ret := C.snd_mixer_elem_next(e.native())
	return wrapMixerElem(ret)
}

// get previous mixer element
func (e MixerElem) Prev() MixerElem {
	ret := C.snd_mixer_elem_prev(e.native())
	return wrapMixerElem(ret)
}

// Return true if mixer simple element has only one volume control for both playback and capture.
func (e MixerElem) SelemHasCommonVolume() bool {
	ret := C.snd_mixer_selem_has_common_volume(e.native())
	return ret != 0
}

// Return info about playback volume control of a mixer simple element.
// return false if no control is present, true if it's present
func (e MixerElem) SelemHasPlaybackVolume() bool {
	ret := C.snd_mixer_selem_has_playback_volume(e.native())
	return ret != 0
}

// Return info about playback volume control of a mixer simple element.
// return false if control is separated per channel, true if control acts on all channels together
func (e MixerElem) SelemHasPlaybackVolumeJoined() bool {
	ret := C.snd_mixer_selem_has_playback_volume_joined(e.native())
	return ret != 0
}

// Return info about capture volume control of a mixer simple element.
// return false if no control is present, true if it's present
func (e MixerElem) SelemHasCaptureVolume() bool {
	ret := C.snd_mixer_selem_has_capture_volume(e.native())
	return ret != 0
}

// Return info about capture volume control of a mixer simple element.
// return false if control is separated per channel, true if control acts on all channels together
func (e MixerElem) SelemHasCaptureVolumeJoined() bool {
	ret := C.snd_mixer_selem_has_capture_volume_joined(e.native())
	return ret != 0
}

// Return true if mixer simple element has only one switch control for both playback and capture.
func (e MixerElem) SelemHasCommonSwitch() bool {
	ret := C.snd_mixer_selem_has_common_switch(e.native())
	return ret != 0
}

// Return info about playback switch control existence of a mixer simple element.
func (e MixerElem) SelemHasPlaybackSwitch() bool {
	ret := C.snd_mixer_selem_has_playback_switch(e.native())
	return ret != 0
}

// Return info about playback switch control of a mixer simple element.
// return false if control is separated per channel, true if control acts on all channels together
func (e MixerElem) SelemHasPlaybackSwitchJoined() bool {
	ret := C.snd_mixer_selem_has_playback_switch_joined(e.native())
	return ret != 0
}

// Return info about capture switch control existence of a mixer simple element.
// return false if no control is present, true if it's present
func (e MixerElem) SelemHasCaptureSwitch() bool {
	ret := C.snd_mixer_selem_has_capture_switch(e.native())
	return ret != 0
}

// Return info about capture switch control of a mixer simple element.
// return false if control is separated per channel, true if control acts on all channels together
func (e MixerElem) SelemHasCaptureSwitchJoined() bool {
	ret := C.snd_mixer_selem_has_capture_switch_joined(e.native())
	return ret != 0
}

// Return info about capture switch control of a mixer simple element.
// return false if control is separated per element, true if control acts on other elements too (i.e. only one active at a time inside a group)
func (e MixerElem) SelemHasCaptureSwitchExclusive() bool {
	ret := C.snd_mixer_selem_has_capture_switch_exclusive(e.native())
	return ret != 0
}

// Return true if mixer simple enumerated element belongs to the playback direction.
func (e MixerElem) SelemIsEnumPlayback() bool {
	ret := C.snd_mixer_selem_is_enum_playback(e.native())
	return ret != 0
}

// Return true if mixer simple enumerated element belongs to the capture direction.
func (e MixerElem) SelemIsEnumCapture() bool {
	ret := C.snd_mixer_selem_is_enum_capture(e.native())
	return ret != 0
}

// Return true if mixer simple element is an enumerated control.
func (e MixerElem) SelemIsEnumerated() bool {
	ret := C.snd_mixer_selem_is_enumerated(e.native())
	return ret != 0
}

// Return the number of enumerated items of the given mixer simple element.
func (e MixerElem) SelemGetEnumItems() (int, error) {
	ret := C.snd_mixer_selem_get_enum_items(e.native())
	if ret < 0 {
		return 0, newError("snd_mixer_selem_get_enum_items", ret)
	}
	return int(ret), nil
}

// get the enumerated item string for the given mixer simple element
func (e MixerElem) SelemGetEnumItemName(item uint, maxLen int) (string, error) {
	buf := (*C.char)(C.malloc(C.size_t(maxLen + 1)))
	defer C.free(unsafe.Pointer(buf))
	ret := C.snd_mixer_selem_get_enum_item_name(e.native(), C.uint(item),
		C.size_t(maxLen), buf)
	if ret == 0 {
		return C.GoString(buf), nil
	}
	return "", newError("snd_mixer_selem_get_enum_item_name", ret)
}

// Return info about capture switch control of a mixer simple element.
// return group for switch exclusivity
func (e MixerElem) SelemGetCaptureGroup() int {
	ret := C.snd_mixer_selem_get_capture_group(e.native())
	return int(ret)
}

// Get info about channels of playback stream of a mixer simple element.
// return false if not mono, true if mono
func (e MixerElem) SelemIsPlaybackMono() bool {
	ret := C.snd_mixer_selem_is_playback_mono(e.native())
	return ret != 0
}

// Get info about channels of playback stream of a mixer simple element.
// return false if channel is not present, true if present
func (e MixerElem) SelemHasPlaybackChannel(channel MixerSelemChannelId) bool {
	ret := C.snd_mixer_selem_has_playback_channel(e.native(),
		C.snd_mixer_selem_channel_id_t(channel))
	return ret != 0
}

// Get info about channels of capture stream of a mixer simple element.
// return false if not mono, true if mono
func (e MixerElem) SelemIsCaptureMono() bool {
	ret := C.snd_mixer_selem_is_capture_mono(e.native())
	return ret != 0
}

// Get info about channels of capture stream of a mixer simple element.
// return false if channel is not present, true if present
func (e MixerElem) SelemHasCaptureChannel(channel MixerSelemChannelId) bool {
	ret := C.snd_mixer_selem_has_capture_channel(e.native(),
		C.snd_mixer_selem_channel_id_t(channel))
	return ret != 0
}

// Get range for playback volume of a mixer simple element.
func (e MixerElem) SelemGetPlaybackVolumeRange() (min, max int) {
	var cMin, cMax C.long
	C.snd_mixer_selem_get_playback_volume_range(e.native(), &cMin, &cMax)
	return int(cMin), int(cMax)
}

// Get range for capture volume of a mixer simple element.
func (e MixerElem) SelemGetCaptureVolumeRange() (min, max int) {
	var cMin, cMax C.long
	C.snd_mixer_selem_get_capture_volume_range(e.native(), &cMin, &cMax)
	return int(cMin), int(cMax)
}

// Return value of playback switch control of a mixer simple element.
func (e MixerElem) SelemGetPlaybackSwitch(channel MixerSelemChannelId) (bool, error) {
	var value C.int
	ret := C.snd_mixer_selem_get_playback_switch(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), &value)
	if ret == 0 {
		return value != 0, nil
	}
	return false, newError("snd_mixer_selem_get_playback_switch", ret)
}

// Return value of capture switch control of a mixer simple element.
func (e MixerElem) SelemGetCaptureSwitch(channel MixerSelemChannelId) (bool, error) {
	var value C.int
	ret := C.snd_mixer_selem_get_capture_switch(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), &value)
	if ret == 0 {
		return value != 0, nil
	}
	return false, newError("snd_mixer_selem_get_capture_switch", ret)
}

// Return value of playback volume control of a mixer simple element.
func (e MixerElem) SelemGetPlaybackVolume(channel MixerSelemChannelId) (int, error) {
	var value C.long
	ret := C.snd_mixer_selem_get_playback_volume(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), &value)
	if ret == 0 {
		return int(value), nil
	}
	return 0, newError("snd_mixer_selem_get_playback_volume", ret)
}

// Return value of capture volume control of a mixer simple element.
func (e MixerElem) SelemGetCaptureVolume(channel MixerSelemChannelId) (int, error) {
	var value C.long
	ret := C.snd_mixer_selem_get_capture_volume(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), &value)
	if ret == 0 {
		return int(value), nil
	}
	return 0, newError("snd_mixer_selem_get_capture_volume", ret)
}

// Return value of playback volume in dB control of a mixer simple element.
// return value (dB * 100)
func (e MixerElem) SelemGetPlaybackDB(channel MixerSelemChannelId) (int, error) {
	var value C.long
	ret := C.snd_mixer_selem_get_playback_dB(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), &value)
	if ret == 0 {
		return int(value), nil
	}
	return 0, newError("snd_mixer_selem_get_playback_dB", ret)
}

// Return value of capture volume in dB control of a mixer simple element.
// return value (dB * 100)
func (e MixerElem) SelemGetCaptureDB(channel MixerSelemChannelId) (int, error) {
	var value C.long
	ret := C.snd_mixer_selem_get_capture_dB(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), &value)
	if ret == 0 {
		return int(value), nil
	}
	return 0, newError("snd_mixer_selem_get_capture_dB", ret)
}

// Get range in dB for playback volume of a mixer simple element.
// return minimum (dB * 100) and maximum (dB * 100)
func (e MixerElem) SelemGetPlaybackDBRange() (min, max int) {
	var cMin, cMax C.long
	C.snd_mixer_selem_get_playback_dB_range(e.native(), &cMin, &cMax)
	return int(cMin), int(cMax)
}

// Get range in dB for capture volume of a mixer simple element.
// return minimum (dB * 100) and maximum (dB * 100)
func (e MixerElem) SelemGetCaptureDBRange() (min, max int) {
	var cMin, cMax C.long
	C.snd_mixer_selem_get_capture_dB_range(e.native(), &cMin, &cMax)
	return int(cMin), int(cMax)
}

// Set value of playback volume control of a mixer simple element.
func (e MixerElem) SelemSetPlaybackVolume(channel MixerSelemChannelId, value int) error {
	ret := C.snd_mixer_selem_set_playback_volume(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), C.long(value))
	return newError("snd_mixer_selem_set_playback_volume", ret)
}

// Set value of capture volume control of a mixer simple element.
func (e MixerElem) SelemSetCaptureVolume(channel MixerSelemChannelId, value int) error {
	ret := C.snd_mixer_selem_set_capture_volume(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), C.long(value))
	return newError("snd_mixer_selem_set_capture_volume", ret)
}

// Set value in dB of playback volume control of a mixer simple element.
// value: control value in dB * 100
// dir: rounding mode - rounds up if dir > 0, otherwise rounds down
func (e MixerElem) SelemSetPlaybackDB(channel MixerSelemChannelId, value, dir int) error {
	ret := C.snd_mixer_selem_set_playback_dB(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), C.long(value), C.int(dir))
	return newError("snd_mixer_selem_set_playback_dB", ret)
}

// Set value in dB of capture volume control of a mixer simple element.
// value: control value in dB * 100
// dir: rounding mode - rounds up if dir > 0, otherwise rounds down
func (e MixerElem) SelemSetCaptureDB(channel MixerSelemChannelId, value, dir int) error {
	ret := C.snd_mixer_selem_set_capture_dB(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), C.long(value), C.int(dir))
	return newError("snd_mixer_selem_set_capture_dB", ret)
}

// Set value of playback volume control for all channels of a mixer simple element.
func (e MixerElem) SelemSetPlaybackVolumeAll(value int) error {
	ret := C.snd_mixer_selem_set_playback_volume_all(e.native(), C.long(value))
	return newError("snd_mixer_selem_set_playback_volume_all", ret)
}

// Set value of capture volume control for all channels of a mixer simple element.
func (e MixerElem) SelemSetCaptureVolumeAll(value int) error {
	ret := C.snd_mixer_selem_set_capture_volume_all(e.native(), C.long(value))
	return newError("snd_mixer_selem_set_capture_volume_all", ret)
}

// Set value in dB of playback volume control for all channels of a mixer simple element.
// value: control value in dB * 100
// dir: rounding mode - rounds up if dir > 0, otherwise rounds down
func (e MixerElem) SelemSetPlaybackDBAll(value, dir int) error {
	ret := C.snd_mixer_selem_set_playback_dB_all(e.native(), C.long(value), C.int(dir))
	return newError("snd_mixer_selem_set_playback_dB_all", ret)
}

// Set value in dB of capture volume control for all channels of a mixer simple element.
// value: control value in dB * 100
// dir: rounding mode - rounds up if dir > 0, otherwise rounds down
func (e MixerElem) SelemSetCaptureDBAll(value, dir int) error {
	ret := C.snd_mixer_selem_set_capture_dB_all(e.native(), C.long(value), C.int(dir))
	return newError("snd_mixer_selem_set_capture_dB_all", ret)
}

// Set value of playback switch control of a mixer simple element.
func (e MixerElem) SelemSetPlaybackSwitch(channel MixerSelemChannelId, value bool) error {
	var cValue C.int
	if value {
		cValue = 1
	}
	ret := C.snd_mixer_selem_set_playback_switch(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), cValue)
	return newError("snd_mixer_selem_set_playback_switch", ret)
}

// Set value of capture switch control of a mixer simple element.
func (e MixerElem) SelemSetCaptureSwitch(channel MixerSelemChannelId, value bool) error {
	var cValue C.int
	if value {
		cValue = 1
	}
	ret := C.snd_mixer_selem_set_capture_switch(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), cValue)
	return newError("snd_mixer_selem_set_capture_switch", ret)
}

// Set value of playback switch control for all channels of a mixer simple element.
func (e MixerElem) SelemSetPlaybackSwitchAll(value bool) error {
	var cValue C.int
	if value {
		cValue = 1
	}
	ret := C.snd_mixer_selem_set_playback_switch_all(e.native(), cValue)
	return newError("snd_mixer_selem_set_playback_switch_all", ret)
}

// Set value of capture switch control for all channels of a mixer simple element.
func (e MixerElem) SelemSetCaptureSwitchAll(value bool) error {
	var cValue C.int
	if value {
		cValue = 1
	}
	ret := C.snd_mixer_selem_set_capture_switch_all(e.native(), cValue)
	return newError("snd_mixer_selem_set_capture_switch_all", ret)
}

// Mixer simple element channel identifier
type MixerSelemChannelId C.snd_mixer_selem_channel_id_t

const (
	MixerSChanUnknown = C.SND_MIXER_SCHN_UNKNOWN
	/* Front left */
	MixerSChanFrontLeft = C.SND_MIXER_SCHN_FRONT_LEFT
	/* Front right */
	MixerSChanFrontRight = C.SND_MIXER_SCHN_FRONT_RIGHT
	/* Rear left */
	MixerSChanRearLeft = C.SND_MIXER_SCHN_REAR_LEFT
	/* Rear right */
	MixerSChanRearRight = C.SND_MIXER_SCHN_REAR_RIGHT
	/* Front center */
	MixerSChanFrontCenter = C.SND_MIXER_SCHN_FRONT_CENTER
	/* Woofer */
	MixerSChanWoofer = C.SND_MIXER_SCHN_WOOFER
	/* Side Left */
	MixerSChanSideLeft = C.SND_MIXER_SCHN_SIDE_LEFT
	/* Side Right */
	MixerSChanSideRight = C.SND_MIXER_SCHN_SIDE_RIGHT
	/* Rear Center */
	MixerSChanRearCenter = C.SND_MIXER_SCHN_REAR_CENTER
	/* Mono (Front left alias) */
	MixerSChanMono = C.SND_MIXER_SCHN_MONO
	MixerSChanLast = C.SND_MIXER_SCHN_LAST
)

// Return name of mixer simple element channel.
func (id MixerSelemChannelId) Name() string {
	return C.GoString(C.snd_mixer_selem_channel_name(C.snd_mixer_selem_channel_id_t(id)))
}

// get the current selected enumerated item for the given mixer simple element
func (e MixerElem) SelemGetEnumItem(channel MixerSelemChannelId) (uint, error) {
	var item C.uint
	ret := C.snd_mixer_selem_get_enum_item(e.native(),
		C.snd_mixer_selem_channel_id_t(channel), &item)
	if ret == 0 {
		return uint(item), nil
	}
	return 0, newError("snd_mixer_selem_get_enum_item", ret)
}

type MixerSelemRegopt struct {
	c      C.struct_snd_mixer_selem_regopt
	device *C.char
}

func (opt *MixerSelemRegopt) Free() {
	if opt.device != nil {
		C.free(unsafe.Pointer(opt.device))
		opt.device = nil
	}
}

func (opt *MixerSelemRegopt) SetDevice(device string) {
	if opt.device != nil {
		C.free(unsafe.Pointer(opt.device))
	}
	opt.device = C.CString(device)
	opt.c.device = opt.device
}

func (opt *MixerSelemRegopt) SetVer(ver int) {
	opt.c.ver = C.int(ver)
}

func (opt *MixerSelemRegopt) SetPlaybackPCM(pcm PCM) {
	opt.c.playback_pcm = pcm.native()
}

func (opt *MixerSelemRegopt) SetCapturePCM(pcm PCM) {
	opt.c.capture_pcm = pcm.native()
}

const (
	MixerSAbstractNone  = C.SND_MIXER_SABSTRACT_NONE
	MixerSAbstractBasic = C.SND_MIXER_SABSTRACT_BASIC
)

type MixerClass struct {
	Ptr unsafe.Pointer
}

// Register mixer simple element class.
func (m Mixer) SelemRegister(options *MixerSelemRegopt, class *MixerClass) error {
	var cOptions *C.struct_snd_mixer_selem_regopt
	if options != nil {
		cOptions = &options.c
	}
	var cClass **C.snd_mixer_class_t
	if class != nil {
		cClass = (**C.snd_mixer_class_t)(unsafe.Pointer(&class.Ptr))
	}
	ret := C.snd_mixer_selem_register(m.native(),
		cOptions, cClass)
	return newError("snd_mixer_selem_register", ret)
}

type MixerSelem struct {
	Ptr unsafe.Pointer
}

// Get mixer simple element identifier.
func (e MixerElem) GetSelemId(id MixerSelemId) {
	C.snd_mixer_selem_get_id(e.native(), id.native())
}

// Get info about the active state of a mixer simple element.
func (e MixerElem) SelemIsActive() bool {
	ret := C.snd_mixer_selem_is_active(e.native())
	return ret != 0
}

type MixerSelemId struct {
	Ptr unsafe.Pointer
}

func wrapMixerSelemId(ptr *C.snd_mixer_selem_id_t) MixerSelemId {
	return MixerSelemId{
		Ptr: unsafe.Pointer(ptr),
	}
}

func (id MixerSelemId) native() *C.snd_mixer_selem_id_t {
	return (*C.snd_mixer_selem_id_t)(id.Ptr)
}

func NewMixerSelemId() (MixerSelemId, error) {
	var ptr *C.snd_mixer_selem_id_t
	ret := C.snd_mixer_selem_id_malloc(&ptr)
	if ret == 0 {
		return wrapMixerSelemId(ptr), nil
	}
	return MixerSelemId{}, newError("snd_mixer_selem_id_malloc", ret)
}

func (id MixerSelemId) Free() {
	C.snd_mixer_selem_id_free(id.native())
}

// Get name part of mixer simple element identifier.
func (id MixerSelemId) GetName() string {
	ret := C.snd_mixer_selem_id_get_name(id.native())
	return C.GoString(ret)
}

// Get index part of a mixer simple element identifier.
func (id MixerSelemId) GetIndex() uint {
	ret := C.snd_mixer_selem_id_get_index(id.native())
	return uint(ret)
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
	errStr := C.GoString(C.snd_strerror(err.code))
	return err.fn + ": " + errStr
}
