// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusutil

import (
	"strings"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
	"github.com/godbus/dbus/v5/prop"
	"github.com/linuxdeepin/go-lib/strv"
)

var peerIntrospectData = introspect.Interface{
	Name: orgFreedesktopDBus + ".Peer",
	Methods: []introspect.Method{
		{
			Name: "Ping",
		},
		{
			Name: "GetMachineId",
			Args: []introspect.Arg{
				{
					Name:      "machine_uuid",
					Type:      "s",
					Direction: "out",
				},
			},
		},
	},
}

func (so *ServerObject) getChildren() (children strv.Strv) {
	var target string
	if so.path == "/" {
		target = "/"
	} else {
		target = string(so.path) + "/"
	}
	for objPath := range so.service.objMap {
		if objPath == so.path {
			continue
		}

		if strings.HasPrefix(string(objPath), target) {
			tail := string(objPath[len(target):])
			idx := strings.Index(tail, "/")
			var child string
			if idx == -1 {
				child = tail
			} else {
				child = tail[:idx]
			}

			if child == "" {
				continue
			} else {
				children, _ = children.Add(child)
			}
		}
	}
	return
}

func (so *ServerObject) introspectableIntrospect() (string, *dbus.Error) {
	so.service.DelayAutoQuit()

	node := &introspect.Node{
		Interfaces: so.getInterfaces(),
	}

	for _, child := range so.getChildren() {
		node.Children = append(node.Children, introspect.Node{Name: child})
	}

	// marshal xml
	return string(introspect.NewIntrospectable(node)), nil
}

func (so *ServerObject) getInterfaces() []introspect.Interface {
	var interfaces []introspect.Interface

	for _, impl := range so.implementers {
		implStatic := impl.getStatic(so.service)
		interfaces = append(interfaces, implStatic.introspectInterface)
	}

	interfaces = append(interfaces, introspect.IntrospectData, prop.IntrospectData, peerIntrospectData)

	return interfaces
}
