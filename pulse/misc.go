// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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

const (
	AvailableTypeUnknow int = iota
	AvailableTypeNo
	AvailableTypeYes
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

type ProfileInfo2 struct {
	Name        string
	Description string
	Priority    uint32
	NSinks      uint32
	NSources    uint32
	Available   int
}
type ProfileInfos2 []ProfileInfo2

type PortInfo struct {
	Name        string
	Description string
	Priority    uint32
	Available   int // 0: Unknow, 1: No, 2: Yes
}
type PortInfos []PortInfo

func (infos PortInfos) Get(name string) *PortInfo {
	for _, info := range infos {
		if info.Name == name {
			return &info
		}
	}
	return nil
}

func (infos ProfileInfos2) Exists(name string) bool {
	for _, info := range infos {
		if info.Name == name {
			return true
		}
	}
	return false
}

func (infos ProfileInfos2) SelectProfile() string {
	if len(infos) == 0 {
		return ""
	}

	profile := infos.selectByAvailable()
	if len(profile.Name) != 0 {
		return profile.Name
	}

	return infos[0].Name
}

func (infos ProfileInfos2) Len() int {
	return len(infos)
}

func (infos ProfileInfos2) Less(i, j int) bool {
	return infos[i].Priority > infos[j].Priority
}

func (infos ProfileInfos2) Swap(i, j int) {
	infos[i], infos[j] = infos[j], infos[i]
}

func (infos ProfileInfos2) selectByAvailable() ProfileInfo2 {
	var profile ProfileInfo2
	for _, info := range infos {
		if info.Available == 0 {
			continue
		}

		if profile.Priority < info.Priority {
			profile = info
		}
	}
	return profile
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

// 音频支持单声道设置，仅在pipewire下有效果
func (cv CVolume) SetMono(volume float64, enable bool) CVolume {
	var v Volume
	v.paVolume = C.pa_volume_t(volume * C.PA_VOLUME_NORM)
	var channels uint32 = 31
	if !enable {
		channels = 32
	}
	C.pa_cvolume_set(&cv.core, C.unsigned(channels), v.paVolume)
	return cv
}
