// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package proxy

import (
	"errors"
	"sync"

	"github.com/godbus/dbus/v5"
)

var globalRuleCounter = &ruleCounter{
	connRuleCount: make(map[string]map[string]uint),
}

type ruleCounter struct {
	connRuleCount map[string]map[string]uint
	//                ^unique name  ^rule
	mu sync.Mutex
}

func getConnName(conn *dbus.Conn) string {
	return conn.Names()[0]
}

func (rc *ruleCounter) addMatch(conn *dbus.Conn, rule string) error {
	name := getConnName(conn)
	if name == "" {
		return errors.New("conn name is empty")
	}
	rc.mu.Lock()
	defer rc.mu.Unlock()

	var count uint
	ruleCount, ok := rc.connRuleCount[name]
	if ok {
		count = ruleCount[rule]
	} else {
		ruleCount = make(map[string]uint)
		rc.connRuleCount[name] = ruleCount
	}
	if count == 0 {
		err := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0, rule).Err
		if err != nil {
			return err
		}
	}
	ruleCount[rule] = count + 1
	return nil
}

func (rc *ruleCounter) removeMatch(conn *dbus.Conn, rule string) error {
	name := getConnName(conn)
	if name == "" {
		return errors.New("conn name is empty")
	}

	rc.mu.Lock()
	defer rc.mu.Unlock()

	var count uint
	ruleCount, ok := rc.connRuleCount[name]
	if !ok {
		return nil
	}

	count = ruleCount[rule]
	if count == 1 {
		err := conn.BusObject().Call("org.freedesktop.DBus.RemoveMatch", 0, rule).Err
		if err != nil {
			return err
		}
	}
	ruleCount[rule] = count - 1
	return nil
}

func (rc *ruleCounter) deleteConn(conn *dbus.Conn) {
	rc.mu.Lock()
	delete(rc.connRuleCount, getConnName(conn))
	rc.mu.Unlock()
}

// DeleteConn 删除连接相关的数据，可以在你关闭这个连接之后调用。
func DeleteConn(conn *dbus.Conn) {
	globalRuleCounter.deleteConn(conn)
}
