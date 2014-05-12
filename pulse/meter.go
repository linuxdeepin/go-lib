package pulse

/*
#include "dde-pulse.h"
*/
import "C"

var sourceMeterCBs = make(map[uint32]func(float64))

type SourceMeter struct {
	core        *C.pa_stream
	sourceIndex uint32
	cb          func()
}

func NewSourceMeter(c *Context, idx uint32) *SourceMeter {
	core := C.createMonitorStreamForSource(c.ctx, C.uint32_t(idx), 0, 0)
	return &SourceMeter{core, idx, nil}
}
func (s *SourceMeter) Destroy() {
	delete(sourceMeterCBs, s.sourceIndex)
	C.pa_stream_disconnect(s.core)
	C.pa_stream_unref(s.core)
}
func (s *SourceMeter) ConnectChanged(cb func(v float64)) {
	sourceMeterCBs[s.sourceIndex] = cb
}

//export go_update_volume_meter
func go_update_volume_meter(source_index uint32, sink_index uint32, v float64) {
	if cb, ok := sourceMeterCBs[source_index]; ok && cb != nil {
		cb(v)
	}
}
