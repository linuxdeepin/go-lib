package pulse

/*
#include "dde-pulse.h"
*/
import "C"
import "unsafe"

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

func (s *Sink) SetPort(name string) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))

	c := GetContext()
	c.lock()
	defer c.unlock()

	C.pa_context_set_sink_port_by_index(c.ctx, C.uint32_t(s.Index), cname, C.success_cb, nil)
}

func (s *Sink) SetMute(mute bool) {
	_mute := 0
	if mute {
		_mute = 1
	}
	c := GetContext()
	c.lock()
	defer c.unlock()
	C.pa_context_set_sink_mute_by_index(c.ctx, C.uint32_t(s.Index), C.int(_mute), C.success_cb, nil)
}

func (s *Sink) SetVolume(v CVolume) {
	s.Volume = v

	c := GetContext()
	c.lock()
	defer c.unlock()
	C.pa_context_set_sink_volume_by_index(c.ctx, C.uint32_t(s.Index), &v.core, C.success_cb, nil)
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
