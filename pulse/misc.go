package pulse

/*
#include "dde-pulse.h"
*/
import "C"
import "unsafe"

const (
	VolumeMax   = C.PA_VOLUME_MAX * 1.0 / C.PA_VOLUME_NORM
	VolumeUIMax = 99957.0 / C.PA_VOLUME_NORM // C.pa_sw_volume_from_dB(11.0)
)

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

type Volume struct {
	paVolume C.pa_volume_t
}

func (v Volume) ToPercent() float64 {
	return float64(v.paVolume) / C.PA_VOLUME_NORM
}
func (v Volume) ToLiner() float64 {
	return float64(C.pa_sw_volume_to_linear(C.pa_volume_t(v.paVolume)))
}
func (v Volume) TodB() float64 {
	return float64(C.pa_sw_volume_to_dB(C.pa_volume_t(v.paVolume)))
}

type CVolume struct {
	core C.pa_cvolume
}

func (cv CVolume) Avg() float64 {
	return float64(C.pa_cvolume_max(&cv.core)) / C.PA_VOLUME_NORM
}
func (cv CVolume) Balance(cmap ChannelMap) float64 {
	return float64(C.pa_cvolume_get_balance(&cv.core, &cmap.core))
}
func (cv CVolume) Fade(cmap ChannelMap) float64 {
	return float64(C.pa_cvolume_get_fade(&cv.core, &cmap.core))
}

func (cv CVolume) SetAvg(v float64) CVolume {
	C.pa_cvolume_scale(&cv.core, C.pa_volume_t((C.double(v) * C.PA_VOLUME_NORM)))
	return cv
}
func (cv CVolume) SetBalance(cm ChannelMap, balance float64) CVolume {
	C.pa_cvolume_set_balance(&cv.core, &cm.core, C.float(balance))
	return cv
}
func (cv CVolume) SetFade(cm ChannelMap, fade float64) CVolume {
	C.pa_cvolume_set_fade(&cv.core, &cm.core, C.float(fade))
	return cv
}

type ChannelPosition int32
type ChannelMap struct {
	core C.pa_channel_map
}

func (cm ChannelMap) CanBalance() bool {
	if C.pa_channel_map_can_balance(&cm.core) == 0 {
		return false
	} else {
		return true
	}
}
func (cm ChannelMap) CanFade() bool {
	if C.pa_channel_map_can_fade(&cm.core) == 0 {
		return false
	} else {
		return true
	}
}
