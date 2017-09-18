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

type Server struct {
	UserName      string
	HostName      string
	ServerVersion string
	ServerName    string
	//sample_spec
	DefaultSinkName   string
	DefaultSourceName string
	Cookie            uint32
	ChannelMap        ChannelMap
}

func toServerInfo(info *C.pa_server_info) *Server {
	s := &Server{}
	s.UserName = C.GoString(info.user_name)
	s.HostName = C.GoString(info.host_name)
	s.ServerVersion = C.GoString(info.server_version)
	s.ServerName = C.GoString(info.server_name)
	//sample_spec
	s.DefaultSinkName = C.GoString(info.default_sink_name)
	s.DefaultSourceName = C.GoString(info.default_source_name)
	s.Cookie = uint32(info.cookie)
	s.ChannelMap = ChannelMap{info.channel_map}
	return s
}
