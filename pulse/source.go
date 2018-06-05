/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package pulse

/*
#include "dde-pulse.h"
*/
import "C"

type Source struct {
	Index uint32

	Name        string
	Description string

	//sample_spec

	ChannelMap        ChannelMap
	OwnerModule       uint32
	Volume            CVolume
	Mute              bool
	MonitorOfSink     uint32
	MonitorOfSinkName string

	//latency pa_usec_t

	Driver string

	Flags int

	Proplist map[string]string

	BaseVolume Volume

	//state

	NVolumeSteps uint32
	Card         uint32
	Ports        []PortInfo
	ActivePort   PortInfo

	//n_formats
	//formats
}

func toSourceInfo(info *C.pa_source_info) *Source {
	s := &Source{}
	s.Index = uint32(info.index)
	s.Name = C.GoString(info.name)
	s.Description = C.GoString(info.description)
	s.ChannelMap = ChannelMap{info.channel_map}
	//sample_spec
	s.OwnerModule = uint32(info.owner_module)
	s.Volume = CVolume{info.volume}
	s.Mute = toBool(info.mute)
	s.MonitorOfSink = uint32(info.monitor_of_sink)
	s.MonitorOfSinkName = C.GoString(info.monitor_of_sink_name)

	//latency pa_usec_t

	s.Driver = C.GoString(info.driver)

	//flags pa_source_flags_t
	s.Flags = int(info.flags)

	s.Proplist = toProplist(info.proplist)
	s.BaseVolume = Volume{info.base_volume}

	//state

	s.NVolumeSteps = uint32(info.n_volume_steps)
	s.Card = uint32(info.card)
	s.Ports = toSourcePorts(uint32(info.n_ports), info.ports)
	s.ActivePort = toSourcePort(info.active_port)

	//n_formats
	//formats
	return s
}
