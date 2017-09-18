/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
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

type SinkInput struct {
	Index       uint32
	Name        string
	OwnerModule uint32
	Client      uint32
	Sink        uint32

	//sample_spec

	ChannelMap ChannelMap
	Volume     CVolume

	//buffer usec
	//sink usec

	ResampleMethod string
	Driver         string

	Mute     bool
	PropList map[string]string
	Corked   int

	HasVolume      bool
	VolumeWritable bool

	//format
}

func (s *SinkInput) SetVolume(v CVolume) {
	s.Volume = v
	c := GetContext()
	c.SafeDo(func() {
		C.pa_context_set_sink_input_volume(c.ctx, C.uint32_t(s.Index), &v.core, C.get_success_cb(), nil)
	})

}
func (sink *SinkInput) SetMute(mute bool) {
	_mute := 0
	if mute {
		_mute = 1
	}
	c := GetContext()
	c.SafeDo(func() {
		C.pa_context_set_sink_input_mute(c.ctx, C.uint32_t(sink.Index), C.int(_mute), C.get_success_cb(), nil)
	})
}

func toSinkInputInfo(info *C.pa_sink_input_info) *SinkInput {
	s := &SinkInput{}
	s.Index = uint32(info.index)
	s.Name = C.GoString(info.name)
	s.OwnerModule = uint32(info.owner_module)
	s.Client = uint32(info.client)
	s.Sink = uint32(info.sink)

	//sample_spec

	s.ChannelMap = ChannelMap{info.channel_map}
	s.Volume = CVolume{info.volume}

	//buffer usec
	//sink usec

	s.ResampleMethod = C.GoString(info.resample_method)
	s.Driver = C.GoString(info.driver)

	s.Mute = toBool(info.mute)

	s.PropList = toProplist(info.proplist)
	s.Corked = int(info.corked)

	s.HasVolume = toBool(info.has_volume)
	s.VolumeWritable = toBool(info.volume_writable)

	//format

	return s
}
