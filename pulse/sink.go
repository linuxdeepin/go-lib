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
	//flags pa_sink_flags

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
	C.pa_context_set_sink_port_by_index(GetContext().ctx, C.uint32_t(s.Index), cname, C.success_cb, nil)
}

func (s *Sink) SetMute(mute bool) {
	_mute := 0
	if mute {
		_mute = 1
	}
	C.pa_context_set_sink_mute_by_index(GetContext().ctx, C.uint32_t(s.Index), C.int(_mute), C.success_cb, nil)
}

func (s *Sink) SetVolume(v CVolume) {
	s.Volume = v
	C.pa_context_set_sink_volume_by_index(GetContext().ctx, C.uint32_t(s.Index), &v.core, C.success_cb, nil)
}
