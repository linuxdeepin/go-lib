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

type PCMSubformat C.snd_pcm_subformat_t

func (subformat PCMSubformat) Name() string {
	ret := C.snd_pcm_subformat_name(C.snd_pcm_subformat_t(subformat))
	return C.GoString(ret)
}

func (subformat PCMSubformat) Desc() string {
	ret := C.snd_pcm_subformat_description(C.snd_pcm_subformat_t(subformat))
	return C.GoString(ret)
}

type Output struct {
	Ptr unsafe.Pointer
}

func (v Output) native() *C.snd_output_t {
	return (*C.snd_output_t)(v.Ptr)
}

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

func (o Output) Flush() error {
	ret := C.snd_output_flush(o.native())
	return newError("snd_output_flush", ret)
}

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

func (o Output) Close() error {
	ret := C.snd_output_close(o.native())
	return newError("snd_output_close", ret)
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

func (pcm PCM) Recover(err int, silent bool) error {
	var silent0 C.int
	if silent {
		silent0 = 1
	}
	ret := C.snd_pcm_recover(pcm.native(), C.int(err), silent0)
	return newError("snd_pcm_recover", ret)
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

func (params PCMHwParams) Dump(out Output) error {
	ret := C.snd_pcm_hw_params_dump(params.native(), out.native())
	return newError("snd_pcm_hw_params_dump", ret)
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

func (pcm PCM) HwParamsSetChannels(params PCMHwParams, val uint) error {
	ret := C.snd_pcm_hw_params_set_channels(pcm.native(), params.native(), C.uint(val))
	return newError("snd_pcm_hw_params_set_channels", ret)
}

func (pcm PCM) HwParamsSetRate(params PCMHwParams, val uint, dir int) error {
	ret := C.snd_pcm_hw_params_set_rate(pcm.native(), params.native(), C.uint(val), C.int(dir))
	return newError("snd_pcm_hw_params_set_rate", ret)
}

func (pcm PCM) HwParamsSetRateNear(params PCMHwParams, val uint) (uint, int, error) {
	cval := C.uint(val)
	var dir C.int
	ret := C.snd_pcm_hw_params_set_rate_near(pcm.native(), params.native(),
		&cval, &dir)
	if ret == 0 {
		return uint(cval), int(dir), nil
	}
	return 0, 0, newError("snd_pcm_hw_params_set_rate_near", ret)
}

func (pcm PCM) HwParamsSetPeriodTimeNear(params PCMHwParams, val uint) (uint, int, error) {
	val0 := C.uint(val)
	var dir0 C.int
	ret := C.snd_pcm_hw_params_set_period_time_near(pcm.native(), params.native(), &val0, &dir0)
	if ret == 0 {
		return uint(val0), int(dir0), nil
	}
	return 0, 0, newError("snd_pcm_hw_params_set_period_time_near", ret)
}

func (pcm PCM) HwParamsSetBufferTimeNear(params PCMHwParams, val uint) (uint, int, error) {
	val0 := C.uint(val)
	var dir0 C.int
	ret := C.snd_pcm_hw_params_set_buffer_time_near(pcm.native(), params.native(), &val0, &dir0)
	if ret == 0 {
		return uint(val0), int(dir0), nil
	}
	return 0, 0, newError("snd_pcm_hw_params_set_buffer_time_near", ret)
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

const sizeOfPointer = unsafe.Sizeof(unsafe.Pointer(nil))

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

func CardNext(rcard *int) error {
	rcard0 := C.int(*rcard)
	ret := C.snd_card_next(&rcard0)
	*rcard = int(rcard0)
	return newError("snd_card_next", ret)
}

func CardLoad(card int) bool {
	ret := C.snd_card_load(C.int(card))
	if ret == 1 {
		//driver is pressent
		return true
	}
	// else 0, driver not pressent
	return false
}

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

func (ctl CTL) Close() error {
	ret := C.snd_ctl_close(ctl.native())
	return newError("snd_ctl_close", ret)
}

func (ctl CTL) CardInfo(info CTLCardInfo) error {
	ret := C.snd_ctl_card_info(ctl.native(), info.native())
	return newError("snd_ctl_card_info", ret)
}

type PCMInfo struct {
	Ptr unsafe.Pointer
}

func NewPCMInfo() PCMInfo {
	p := make([]byte, C.snd_pcm_info_sizeof())
	return PCMInfo{unsafe.Pointer(&p[0])}
}

func (v PCMInfo) native() *C.snd_pcm_info_t {
	return (*C.snd_pcm_info_t)(v.Ptr)
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

func (info PCMInfo) GetID() string {
	ret := C.snd_pcm_info_get_id(info.native())
	return C.GoString(ret)
}

func (info PCMInfo) GetName() string {
	ret := C.snd_pcm_info_get_name(info.native())
	return C.GoString(ret)
}

func (info PCMInfo) GetStream() PCMStream {
	ret := C.snd_pcm_info_get_stream(info.native())
	return PCMStream(ret)
}

type PCMSubclass C.snd_pcm_subclass_t

func (info PCMInfo) GetSubclass() PCMSubclass {
	ret := C.snd_pcm_info_get_subclass(info.native())
	return PCMSubclass(ret)
}

func (info PCMInfo) GetSubdevice() uint {
	ret := C.snd_pcm_info_get_subdevice(info.native())
	return uint(ret)
}

func (info PCMInfo) GetSubdeviceName() string {
	ret := C.snd_pcm_info_get_subdevice_name(info.native())
	return C.GoString(ret)
}

// Get available subdevices count from a PCM info container.
func (info PCMInfo) GetSubdevicesAvail() uint {
	ret := C.snd_pcm_info_get_subdevices_avail(info.native())
	return uint(ret)
}

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

func NewCTLCardInfo() CTLCardInfo {
	p := make([]byte, C.snd_ctl_card_info_sizeof())
	return CTLCardInfo{unsafe.Pointer(&p[0])}
}

func (info CTLCardInfo) GetID() string {
	ret := C.snd_ctl_card_info_get_id(info.native())
	return C.GoString(ret)
}

func (info CTLCardInfo) GetCard() int {
	ret := C.snd_ctl_card_info_get_card(info.native())
	return int(ret)
}

func (info CTLCardInfo) GetComponents() string {
	ret := C.snd_ctl_card_info_get_components(info.native())
	return C.GoString(ret)
}

func (info CTLCardInfo) GetDriver() string {
	ret := C.snd_ctl_card_info_get_driver(info.native())
	return C.GoString(ret)
}

func (info CTLCardInfo) GetLongName() string {
	ret := C.snd_ctl_card_info_get_longname(info.native())
	return C.GoString(ret)
}

func (info CTLCardInfo) GetMixerName() string {
	ret := C.snd_ctl_card_info_get_mixername(info.native())
	return C.GoString(ret)
}

func (info CTLCardInfo) GetName() string {
	ret := C.snd_ctl_card_info_get_name(info.native())
	return C.GoString(ret)
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
