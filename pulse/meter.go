/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
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
