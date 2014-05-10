package pulse

/*
#include "dde-pulse.h"
*/
import "C"
import "unsafe"

func toProplist(c *C.pa_proplist) map[string]string {
	var ret = make(map[string]string)
	var state unsafe.Pointer
	for key := C.pa_proplist_iterate(c, &state); key != nil; {
		ret[C.GoString(key)] = C.GoString(C.pa_proplist_gets(c, key))
		key = C.pa_proplist_iterate(c, &state)
	}
	return ret
}

type PortInfo struct {
	Name        string
	Description string
	Priority    uint32
	Available   int
}

func toBool(v C.int) bool {
	if int(v) == 0 {
		return false
	} else {
		return true
	}
}

type Volume uint32

func (v Volume) ToLiner() float64 {
	return float64(C.pa_sw_volume_to_linear(C.pa_volume_t(v)))
}
func (v Volume) TodB() float64 {
	return float64(C.pa_sw_volume_to_dB(C.pa_volume_t(v)))
}

type CVolume struct {
	core C.pa_cvolume
}

func (cv *CVolume) Avg() float64 {
	return float64(C.pa_cvolume_max(&cv.core)) / C.PA_VOLUME_NORM
}
func (cv *CVolume) SetAvg(v float64, cm ChannelMap) {
	C.pa_cvolume_scale(&cv.core, C.pa_volume_t((C.double(v) * C.PA_VOLUME_NORM)))
}

func (c *CVolume) Balance(cmap ChannelMap) float64 {
	return float64(C.pa_cvolume_get_balance(&c.core, &cmap.core))
}

func (c *CVolume) SetBalance(cm ChannelMap, balance float64) {
	C.pa_cvolume_set_balance(&c.core, &cm.core, C.float(balance))
}

func toCWVolume(v float64) C.pa_volume_t {
	return C.pa_sw_volume_from_linear(C.double(v))
}

type ChannelPosition int32
type ChannelMap struct {
	core C.pa_channel_map
}
