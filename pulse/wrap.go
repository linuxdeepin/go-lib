// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

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

func toProfiles(n uint32, c **C.pa_card_profile_info2) (r ProfileInfos2) {
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

func toCardPorts(n uint32, c **C.pa_card_port_info) (r CardPortInfos) {
	if n > 0 {
		pp := (*[1 << 10](*C.pa_card_port_info))(unsafe.Pointer(c))[:n:n]
		for _, p := range pp {
			r = append(r, toCardPort(p))
		}
	}
	return
}

func toCardPort(c *C.pa_card_port_info) CardPortInfo {
	if c == nil {
		return CardPortInfo{}
	}
	return CardPortInfo{
		PortInfo: PortInfo{
			Name:        C.GoString(c.name),
			Description: C.GoString(c.description),
			Priority:    uint32(c.priority),
			Available:   int(c.available),
		},
		Direction: int(c.direction),
		Profiles:  toProfiles(uint32(c.n_profiles), c.profiles2),
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

type Module struct {
	Index      uint32
	Name       string
	Argument   string
	NUsed      uint32
	AutoUnload int // deprecated, but keep for completeness
	PropList   map[string]string
}

func toModuleInfo(info *C.pa_module_info) *Module {
	if info == nil {
		return nil
	}
	return &Module{
		Index:      uint32(info.index),
		Name:       C.GoString(info.name),
		Argument:   C.GoString(info.argument),
		NUsed:      uint32(info.n_used),
		AutoUnload: int(info.auto_unload),
		PropList:   toProplist(info.proplist),
	}
}
