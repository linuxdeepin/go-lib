// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pulse

/*
#include "dde-pulse.h"
*/
import "C"

type Sink struct {
	Index uint32

	Name        string
	Description string

	//sample_spec

	ChannelMap ChannelMap

	OwnerModule       uint32
	Volume            CVolume
	Mute              bool
	MonitorSource     uint32
	MonitorSourceName string
	//latency pa_usec_t

	Driver string
	Flags  int

	PropList map[string]string

	//configured_latency pa_usec_t

	BaseVolume Volume

	//state pa_sink_state_t

	NVolumeSteps uint32

	Card uint32

	Ports      []PortInfo
	ActivePort PortInfo

	//n_formats
	//formats
}

func toSinkInfo(info *C.pa_sink_info) *Sink {
	s := &Sink{}

	s.Index = uint32(info.index)

	s.Name = C.GoString(info.name)

	s.Description = C.GoString(info.description)

	//sample_spec

	s.ChannelMap = ChannelMap{info.channel_map}

	s.OwnerModule = uint32(info.owner_module)

	s.Volume = CVolume{info.volume}

	s.Mute = toBool(info.mute)

	s.MonitorSource = uint32(info.monitor_source)

	s.MonitorSourceName = C.GoString(info.monitor_source_name)

	//latency

	s.Flags = int(info.flags)

	s.Driver = C.GoString(info.driver)

	s.PropList = toProplist(info.proplist)

	//configured_latency

	s.BaseVolume = Volume{info.base_volume}

	//state

	s.NVolumeSteps = uint32(info.n_volume_steps)

	s.Card = uint32(info.card)

	s.Ports = toPorts(uint32(info.n_ports), info.ports)

	s.ActivePort = toPort(info.active_port)

	//n_formats
	//formats

	return s
}
