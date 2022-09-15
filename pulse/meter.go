// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package pulse

/*
#include "dde-pulse.h"
*/
import "C"
import "sync"

var sourceMeterCBs = make(map[uint32]func(float64))
var sourceMeterLock sync.RWMutex

type SourceMeter struct {
	core        *C.pa_stream
	sourceIndex uint32
	ctx         *Context
}

func NewSourceMeter(c *Context, idx uint32) *SourceMeter {
	core := C.createMonitorStreamForSource(c.loop, c.ctx, C.uint32_t(idx), 0, 0)
	return &SourceMeter{
		core:        core,
		sourceIndex: idx,
		ctx:         c,
	}
}
func (s *SourceMeter) Destroy() {
	sourceMeterLock.Lock()
	delete(sourceMeterCBs, s.sourceIndex)
	sourceMeterLock.Unlock()

	s.ctx.safeDo(func() {
		C.pa_stream_disconnect(s.core)
		C.pa_stream_unref(s.core)
	})
}

func (s *SourceMeter) ConnectChanged(cb func(v float64)) {
	sourceMeterLock.Lock()
	sourceMeterCBs[s.sourceIndex] = cb
	sourceMeterLock.Unlock()
}
