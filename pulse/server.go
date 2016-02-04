/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

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
