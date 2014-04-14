/**
 * Copyright (c) 2011 ~ 2014 Deepin, Inc.
 *               2013 ~ 2014 jouyouyun
 *
 * Author:      jouyouyun <jouyouwen717@gmail.com>
 * Maintainer:  jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software; you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation; either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program; if not, see <http://www.gnu.org/licenses/>.
 **/

package utils

import (
        libpolkit1 "dbus/org/freedesktop/policykit1"
        "dlib/dbus"
)

type polkitSubject struct {
        /*
         * The following kinds of subjects are known:
         * Unix Process: should be set to unix-process with keys
         *                  pid (of type uint32) and
         *                  start-time (of type uint64)
         * Unix Session: should be set to unix-session with the key
         *                  session-id (of type string)
         * System Bus Name: should be set to system-bus-name with the key
         *                  name (of type string)
         */
        SubjectKind    string
        SubjectDetails map[string]dbus.Variant
}

func (op *Manager) PolkitAuthWithPid(actionId string, pid uint32) bool {
        objPolkit, err := libpolkit1.NewAuthority("org.freedesktop.PolicyKit1",
                "/org/freedesktop/PolicyKit1/Authority")
        if err != nil {
                logger.Warning("New PolicyKit Object Failed: ", err)
                return false
        }

        subject := polkitSubject{}
        subject.SubjectKind = "unix-process"
        subject.SubjectDetails = make(map[string]dbus.Variant)
        subject.SubjectDetails["pid"] = dbus.MakeVariant(uint32(pid))
        subject.SubjectDetails["start-time"] = dbus.MakeVariant(uint64(0))
        details := make(map[string]string)
        details[""] = ""
        flags := uint32(1)
        cancelId := ""

        _, err = objPolkit.CheckAuthorization(subject, actionId, details,
                flags, cancelId)
        if err != nil {
                logger.Warning("CheckAuthorization Failed: ", err)
                return false
        }

        return true
}
