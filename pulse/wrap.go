package pulse

/*
#include "dde-pulse.h"
#cgo pkg-config: libpulse
*/
import "C"
import "unsafe"

func (info *paInfo) ToServer() *Server {
	switch r := info.data.(type) {
	case *Server:
		return r
	}
	return nil
}

func (info *paInfo) ToCard() *Card {
	switch r := info.data.(type) {
	case *Card:
		return r
	}
	return nil
}

func (info *paInfo) ToSink() *Sink {
	switch r := info.data.(type) {
	case *Sink:
		return r
	}
	return nil
}
func (info *paInfo) ToSinkInput() *SinkInput {
	switch r := info.data.(type) {
	case *SinkInput:
		return r
	}
	return nil
}

func (info *paInfo) ToSource() *Source {
	switch r := info.data.(type) {
	case *Source:
		return r
	}
	return nil
}
func (info *paInfo) ToSourceOutput() *SourceOutput {
	switch r := info.data.(type) {
	case *SourceOutput:
		return r
	}
	return nil
}

func toProfiles(n uint32, c **C.pa_card_profile_info2) (r []ProfileInfo2) {
	if n > 0 {
		pp := (*[1 << 10](*C.pa_card_profile_info2))(unsafe.Pointer(c))[:n:n]
		for _, p := range pp {
			r = append(r, toProfile(p))
		}
	}
	return
}
func toProfile(c *C.pa_card_profile_info2) ProfileInfo2 {
	if nil == c {
		return ProfileInfo2{}
	}
	return ProfileInfo2{
		Name:        C.GoString(c.name),
		Description: C.GoString(c.description),
		Priority:    uint32(c.priority),
		NSinks:      uint32(c.n_sinks),
		NSources:    uint32(c.n_sources),
		Available:   int(c.available),
	}
}

func toPorts(n uint32, c **C.pa_sink_port_info) (r []PortInfo) {
	if n > 0 {
		pp := (*[1 << 10](*C.pa_sink_port_info))(unsafe.Pointer(c))[:n:n]
		for _, p := range pp {
			r = append(r, toPort(p))
		}
	}
	return
}
func toPort(c *C.pa_sink_port_info) PortInfo {
	if c == nil {
		return PortInfo{}
	}
	return PortInfo{
		Name:        C.GoString(c.name),
		Description: C.GoString(c.description),
		Priority:    uint32(c.priority),
		Available:   int(c.available),
	}
}

func toSourcePorts(n uint32, c **C.pa_source_port_info) (r []PortInfo) {
	if n > 0 {
		pp := (*[1 << 10](*C.pa_source_port_info))(unsafe.Pointer(c))[:n:n]
		for _, p := range pp {
			r = append(r, toSourcePort(p))
		}
	}
	return
}
func toSourcePort(c *C.pa_source_port_info) PortInfo {
	if c == nil {
		return PortInfo{}
	}
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
