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
	c.lock()
	defer c.unlock()
	C.pa_context_set_sink_input_volume(c.ctx, C.uint32_t(s.Index), &v.core, C.success_cb, nil)
}
func (sink *SinkInput) SetMute(mute bool) {
	_mute := 0
	if mute {
		_mute = 1
	}
	c := GetContext()
	c.lock()
	defer c.unlock()
	C.pa_context_set_sink_input_mute(c.ctx, C.uint32_t(sink.Index), C.int(_mute), C.success_cb, nil)
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
