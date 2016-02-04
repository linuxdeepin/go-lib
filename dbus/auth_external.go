/**
 * Copyright (C) 2014 Deepin Technology Co., Ltd.
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 **/

package dbus

import (
	"encoding/hex"
)

// AuthExternal returns an Auth that authenticates as the given user with the
// EXTERNAL mechanism.
func AuthExternal(user string) Auth {
	return authExternal{user}
}

// AuthExternal implements the EXTERNAL authentication mechanism.
type authExternal struct {
	user string
}

func (a authExternal) FirstData() ([]byte, []byte, AuthStatus) {
	b := make([]byte, 2*len(a.user))
	hex.Encode(b, []byte(a.user))
	return []byte("EXTERNAL"), b, AuthOk
}

func (a authExternal) HandleData(b []byte) ([]byte, AuthStatus) {
	return nil, AuthError
}
