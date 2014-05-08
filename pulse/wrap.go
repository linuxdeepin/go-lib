package pulse

/*
#include "dde-pulse.h"
#cgo pkg-config: libpulse
*/
import "C"
import "fmt"
import "unsafe"

func (info *paInfo) ToSink() *Sink {
	switch r := info.data.(type) {
	case *Sink:
		return r
	default:
		panic(fmt.Sprintln("the type of paInfo is not Sink", r))
	}
}
func (info *paInfo) ToSinkInput() *SinkInput {
	switch r := info.data.(type) {
	case *SinkInput:
		return r
	default:
		panic(fmt.Sprintln("the type of paInfo is not SinkInput", r))
	}
}

func (info *paInfo) ToSource() *Source {
	switch r := info.data.(type) {
	case *Source:
		return r
	default:
		panic(fmt.Sprintln("the type of paInfo is not Source", r))
	}
}
func (info *paInfo) ToSourceOutput() *SourceOutput {
	switch r := info.data.(type) {
	case *SourceOutput:
		return r
	default:
		panic(fmt.Sprintln("the type of paInfo is not SourceOutput", r))
	}
}

func toSinkInfo(info *C.pa_sink_info) *Sink {
	s := &Sink{}

	s.Index = uint32(info.index)

	s.Name = C.GoString(info.name)

	s.Description = C.GoString(info.description)

	//sample_spec

	s.ChannelMap = toChannelMap(info.channel_map)

	s.OwnerModule = uint32(info.owner_module)

	s.Volume = toCVolume(info.volume)

	if int(info.mute) == 0 {
		s.Mute = false
	} else {
		s.Mute = true
	}

	s.MonitorSource = uint32(info.monitor_source)

	s.MonitorSourceName = C.GoString(info.monitor_source_name)

	//latency
	//flags

	s.Driver = C.GoString(info.driver)

	s.PropList = toProplist(info.proplist)

	//configured_latency

	s.BaseVolume = Volume(info.base_volume)

	//state

	s.NVolumeSteps = uint32(info.n_volume_steps)

	s.Card = uint32(info.card)

	s.Ports = toPorts(uint32(info.n_ports), info.ports)

	s.ActivePort = toPort(info.active_port)

	//n_formats
	//formats

	return s
}

func toPorts(n uint32, c **C.pa_sink_port_info) (r []PortInfo) {
	pp := (*[1 << 30](*C.pa_sink_port_info))(unsafe.Pointer(c))[:n:n]
	for _, p := range pp {
		r = append(r, toPort(p))
	}
	return
}
func toPort(c *C.pa_sink_port_info) PortInfo {
	return PortInfo{
		Name:        C.GoString(c.name),
		Description: C.GoString(c.description),
		Priority:    uint32(c.priority),
		Available:   int(c.available),
	}
}

type SampleFormat int32
type SampleSpec struct {
	Format   SampleFormat
	Rate     uint32
	Channels uint8
}

func toSampleSpec(c *C.pa_sample_spec) *SampleSpec {
	return nil
}

func toSinkInputInfo(info *C.pa_sink_input_info) *SinkInput {
	s := &SinkInput{}
	s.Index = uint32(info.index)
	s.Name = C.GoString(info.name)
	s.OwnerModule = uint32(info.owner_module)
	s.Client = uint32(info.client)
	s.Sink = uint32(info.sink)

	//sample_spec

	s.ChannelMap = toChannelMap(info.channel_map)
	s.Volume = toCVolume(info.volume)

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
func toSourceInfo(info *C.pa_source_info) *Source {
	return nil
}

func toSourceOutputInfo(info *C.pa_source_output_info) *SourceOutput {
	return nil
}
