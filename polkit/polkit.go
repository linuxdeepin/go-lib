/*
 * Copyright (C) 2017 ~ 2017 Deepin Technology Co., Ltd.
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

package polkit

import (
	"errors"
	"pkg.deepin.io/lib/dbus"
	pk "pkg.deepin.io/lib/polkit/policykit1"
)

// dbus
const (
	dbusDest    = "org.freedesktop.PolicyKit1"
	dbusObjPath = "/org/freedesktop/PolicyKit1/Authority"
	//dbusIfc = "org.freedesktop.PolicyKit1.Authority"
)

// CheckAuthorizationFlags
const (
	CheckAuthorizationFlagsNone                 = 0
	CheckAuthorizationFlagsAllowUserInteraction = 1
)

type Subject struct {
	Kind    string
	Details map[string]dbus.Variant
}

func NewSubject(kind string) *Subject {
	return &Subject{
		Kind:    kind,
		Details: make(map[string]dbus.Variant),
	}
}

func (s *Subject) SetDetail(key string, value interface{}) {
	s.Details[key] = dbus.MakeVariant(value)
}

// SubjectKind
const (
	SubjectKindUnixProcess   = "unix-process"
	SubjectKindUnixSession   = "unix-session"
	SubjectKindSystemBusName = "system-bus-name"
)

type Identify struct {
	Kind    string
	Details map[string]dbus.Variant
}

type AuthorizationResult struct {
	IsAuthorized bool
	IsChallenge  bool
	Details      map[string]string
}

var _authority *pk.Authority
var _inited bool

func Init() {
	if _inited {
		return
	}

	_authority, _ = pk.NewAuthority(dbusDest, dbusObjPath)
}

func CheckAuthorization(subject *Subject, actionId string, details map[string]string,
	flags uint32, cancellationId string) (*AuthorizationResult, error) {

	ret, err := _authority.CheckAuthorization(subject, actionId, details, flags, cancellationId)
	if err != nil {
		return nil, err
	}

	if len(ret) != 3 {
		return nil, errors.New("length of ret is not 3")
	}

	isAuthorized, ok := ret[0].(bool)
	if !ok {
		return nil, errors.New("result field 0 type is not bool")
	}

	isChallenge, ok := ret[1].(bool)
	if !ok {
		return nil, errors.New("result field 1 type is not bool")
	}

	retDetails, ok := ret[2].(map[string]string)
	if !ok {
		return nil, errors.New("result field 2 type is not map[string]string")
	}

	return &AuthorizationResult{
		IsAuthorized: isAuthorized,
		IsChallenge:  isChallenge,
		Details:      retDetails,
	}, nil
}
