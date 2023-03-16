// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
// 
// SPDX-License-Identifier: GPL-3.0-or-later
package dbusutilv1

import (
	"bytes"
	"fmt"

	"github.com/godbus/dbus/v5"
)

type MatchRule struct {
	Str string
}

func (mr MatchRule) AddTo(conn *dbus.Conn) error {
	return conn.BusObject().Call(orgFreedesktopDBus+".AddMatch", 0, mr.Str).Err
}

func (mr MatchRule) RemoveFrom(conn *dbus.Conn) error {
	return conn.BusObject().Call(orgFreedesktopDBus+".RemoveMatch", 0, mr.Str).Err
}

type matchRuleItem struct {
	key, value string
}

type MatchRuleBuilder struct {
	items []matchRuleItem
}

func (b *MatchRuleBuilder) addItem(k, v string) {
	b.items = append(b.items, matchRuleItem{k, v})
}

func NewMatchRuleBuilder() *MatchRuleBuilder {
	return &MatchRuleBuilder{}
}

func (b *MatchRuleBuilder) Type(type0 string) *MatchRuleBuilder {
	b.addItem("type", type0)
	return b
}

func (b *MatchRuleBuilder) Path(path string) *MatchRuleBuilder {
	b.addItem("path", path)
	return b
}

func (b *MatchRuleBuilder) Sender(sender string) *MatchRuleBuilder {
	b.addItem("sender", sender)
	return b
}

func (b *MatchRuleBuilder) Interface(interface0 string) *MatchRuleBuilder {
	b.addItem("interface", interface0)
	return b
}

func (b *MatchRuleBuilder) Member(member string) *MatchRuleBuilder {
	b.addItem("member", member)
	return b
}

func (b *MatchRuleBuilder) PathNamespace(pathNamespace string) *MatchRuleBuilder {
	b.addItem("path_namespace", pathNamespace)
	return b
}

func (b *MatchRuleBuilder) Destination(destination string) *MatchRuleBuilder {
	b.addItem("destination", destination)
	return b
}

func (b *MatchRuleBuilder) Eavesdrop(eavesdrop bool) *MatchRuleBuilder {
	val := "false"
	if eavesdrop {
		val = "true"
	}
	b.addItem("eavesdrop", val)
	return b
}

func (b *MatchRuleBuilder) Arg(idx int, value string) *MatchRuleBuilder {
	b.addItem(fmt.Sprintf("arg%d", idx), value)
	return b
}

func (b *MatchRuleBuilder) ArgPath(idx int, path string) *MatchRuleBuilder {
	b.addItem(fmt.Sprintf("arg%dpath", idx), path)
	return b
}

func (b *MatchRuleBuilder) ArgNamespace(idx int, namespace string) *MatchRuleBuilder {
	b.addItem(fmt.Sprintf("arg%dnamespace", idx), namespace)
	return b
}

func (b *MatchRuleBuilder) ExtPropertiesChanged(path, interfaceName string) *MatchRuleBuilder {
	return b.Type("signal").
		Path(path).
		Interface(orgFreedesktopDBus+".Properties").
		Member("PropertiesChanged").
		ArgNamespace(0, interfaceName)
}

func (b *MatchRuleBuilder) ExtSignal(path, interfaceName, name string) *MatchRuleBuilder {
	return b.Type("signal").
		Path(path).
		Interface(interfaceName).
		Member(name)
}

func (b *MatchRuleBuilder) BuildStr() string {
	var buf bytes.Buffer
	for _, item := range b.items {
		buf.WriteString(item.key)
		buf.Write([]byte("='"))
		buf.WriteString(item.value)
		buf.Write([]byte("',"))
	}
	data := buf.Bytes()
	return string(data[:len(data)-1]) // remove tailing comma
}

func (b *MatchRuleBuilder) Build() MatchRule {
	return MatchRule{Str: b.BuildStr()}
}
