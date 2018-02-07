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

import "unsafe"
import "sync"

/*
#include "dde-pulse.h"
#cgo pkg-config: libpulse
*/
import "C"

type paInfo struct {
	data interface{}
	Type int
}

func NewPaInfo(data unsafe.Pointer, Type int) *paInfo {
	info := &paInfo{Type: Type}
	switch Type {
	case C.PA_SUBSCRIPTION_EVENT_SINK:
		info.data = toSinkInfo((*C.pa_sink_info)(data))
	case C.PA_SUBSCRIPTION_EVENT_SINK_INPUT:
		info.data = toSinkInputInfo((*C.pa_sink_input_info)(data))
	case C.PA_SUBSCRIPTION_EVENT_SOURCE:
		info.data = toSourceInfo((*C.pa_source_info)(data))
	case C.PA_SUBSCRIPTION_EVENT_SOURCE_OUTPUT:
		info.data = toSourceOutputInfo((*C.pa_source_output_info)(data))
	case C.PA_SUBSCRIPTION_EVENT_SERVER:
		info.data = toServerInfo((*C.pa_server_info)(data))
	case C.PA_SUBSCRIPTION_EVENT_CARD:
		info.data = toCardInfo((*C.pa_card_info)(data))
	default:
		// current didn't support this type
		return nil
	}
	return info
}

type cookie struct {
	id   int64
	data chan *paInfo
}

func NewCookie(id int64) *cookie {
	return &cookie{int64(id), make(chan *paInfo)}
}
func (c *cookie) Reply() *paInfo {
	defer deleteCookie(c.id)
	return <-c.data
}
func (c *cookie) ReplyList() []*paInfo {
	defer deleteCookie(c.id)
	var infos []*paInfo
	for info := range c.data {
		infos = append(infos, info)
	}
	return infos
}

func (c *cookie) Feed(infoType int, info unsafe.Pointer) {
	paInfo := NewPaInfo(info, infoType)
	if paInfo == nil {
		return
	}
	c.data <- paInfo
}
func (c *cookie) EndOfList() {
	close(c.data)
	deleteCookie(c.id)
}

func (c *cookie) Failed() {
	close(c.data)
	deleteCookie(c.id)
}

var newCookie, fetchCookie, deleteCookie = func() (func() *cookie,
	func(int64) *cookie,
	func(int64)) {

	cookies := make(map[int64]*cookie)
	id := int64(0)
	var locker sync.Mutex
	return func() *cookie {
			locker.Lock()
			id++
			c := NewCookie(id)
			cookies[c.id] = c
			locker.Unlock()
			return c
		}, func(i int64) *cookie {
			locker.Lock()
			c := cookies[i]
			locker.Unlock()

			return c
		}, func(i int64) {
			locker.Lock()
			delete(cookies, i)
			locker.Unlock()
		}
}()
