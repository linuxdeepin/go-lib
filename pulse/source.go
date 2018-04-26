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
import "unsafe"

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

func (s *Source) SetPort(name string) {
	c := GetContext()
	c.safeDo(func() {
		cname := C.CString(name)
		C.pa_context_set_source_port_by_index(c.ctx, C.uint32_t(s.Index), cname, C.get_success_cb(), nil)
		C.free(unsafe.Pointer(cname))
	})
}

func (s *Source) SetVolume(v CVolume) {
	s.Volume = v
	c := GetContext()

	c.safeDo(func() {
		C.pa_context_set_source_volume_by_index(c.ctx, C.uint32_t(s.Index), &s.Volume.core, C.get_success_cb(), nil)
	})
}

func (s *Source) SetMute(mute bool) {
	_mute := 0
	if mute {
		_mute = 1
	}

	c := GetContext()

	c.safeDo(func() {
		C.pa_context_set_source_mute_by_index(c.ctx, C.uint32_t(s.Index), C.int(_mute), C.get_success_cb(), nil)
	})
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
