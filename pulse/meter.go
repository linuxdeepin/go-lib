/*
 * Copyright (C) 2014 ~ 2017 Deepin Technology Co., Ltd.
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
