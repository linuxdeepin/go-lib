package pulse

/*
#include "dde-pulse.h"
#cgo pkg-config: libpulse
*/
import "C"
import "fmt"
import "unsafe"

func (info *paInfo) ToServer() *Server {
	switch r := info.data.(type) {
	case *Server:
		return r
	default:
		panic(fmt.Sprintln("the type of paInfo is not Sink", r))
	}
}

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
