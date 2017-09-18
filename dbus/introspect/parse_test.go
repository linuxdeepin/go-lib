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

package introspect

import (
	"bytes"
	C "github.com/smartystreets/goconvey/convey"
	"testing"
)

var testXML = `
<node>
  <interface name="org.bluez.Device1">
    <method name="Disconnect"/>
    <method name="Connect"/>
    <method name="ConnectProfile">
      <annotation name="org.freedesktop.DBus.GLib.CSymbol" value="impl_manager_activate_connection"/>
      <annotation name="org.freedesktop.DBus.GLib.Async" value=""/>
      <arg name="UUID" type="s" direction="in">
	<annotation name="com.deepin.DBus.I18n.Dir" value="true"/>
      </arg>
    </method>
    <method name="DisconnectProfile">
      <arg name="UUID" type="s" direction="in"/>
    </method>
    <method name="Pair">
      <annotation name="org.freedesktop.DBus.Method.NoReply" value="true"/>
    </method>
    <method name="CancelPairing"/>

    <property name="Address" type="s" access="read">
      <annotation name="com.deepin.DBus.I18n.Dir" value=""/>
      <annotation name="com.deepin.DBus.I18n.Domain" value="dde-daemon"/>
    </property>
    <property name="Name" type="s" access="read">
      <annotation name="com.deepin.DBus.I18n.Domain" value="test"/>
    </property>
    <property name="Alias" type="s" access="readwrite"/>
    <property name="Class" type="u" access="read"/>
    <property name="Appearance" type="q" access="read"/>
    <property name="Icon" type="s" access="read"/>
    <property name="Paired" type="b" access="read"/>
    <property name="Trusted" type="b" access="readwrite"/>
    <property name="Blocked" type="b" access="readwrite"/>
    <property name="LegacyPairing" type="b" access="read"/>
    <property name="RSSI" type="n" access="read"/>
    <property name="Connected" type="b" access="read"/>
    <property name="UUIDs" type="as" access="read"/>
    <property name="Modalias" type="s" access="read"/>
    <property name="Adapter" type="o" access="read"/>
  </interface>
</node>
`

func TestParse(t *testing.T) {
	C.Convey("Create a reader ", t, func() {
		reader := bytes.NewBufferString(testXML)
		ninfo, err := Parse(reader)
		C.So(err, C.ShouldBeNil)
		C.Convey("Check interfaces", func() {
			C.So(len(ninfo.Interfaces), C.ShouldEqual, 1)
			ifc := ninfo.Interfaces[0]
			C.So(ifc.Name, C.ShouldEqual, "org.bluez.Device1")
			C.So(len(ifc.Methods), C.ShouldEqual, 6)
			C.So(len(ifc.Properties), C.ShouldEqual, 15)

			C.Convey("Check annotations", func() {
				m := ifc.Methods[2]
				C.So(m.Name, C.ShouldEqual, "ConnectProfile")
				C.So(m.Annotations[0].Name, C.ShouldEqual, "org.freedesktop.DBus.GLib.CSymbol")
				C.So(m.Annotations[0].Value, C.ShouldEqual, "impl_manager_activate_connection")
				dir, domain, ok := m.Args[0].I18nInfo()
				C.So(ok, C.ShouldEqual, false)

				C.So(ifc.Methods[4].Name, C.ShouldEqual, "Pair")
				C.So(ifc.Methods[4].NoReply(), C.ShouldEqual, true)

				dir, domain, ok = ifc.Methods[3].Args[0].I18nInfo()
				C.So(dir, C.ShouldEqual, "")
				C.So(domain, C.ShouldEqual, "")
				C.So(ok, C.ShouldEqual, false)

				p := ifc.Properties[0]
				C.So(p.Name, C.ShouldEqual, "Address")
				dir, domain, ok = p.I18nInfo()
				C.So(dir, C.ShouldEqual, "")
				C.So(domain, C.ShouldEqual, "dde-daemon")
				C.So(ok, C.ShouldEqual, true)

				dir, domain, ok = ifc.Properties[1].I18nInfo()
				C.So(domain, C.ShouldEqual, "test")
				C.So(ok, C.ShouldEqual, true)
				dir, domain, ok = ifc.Properties[2].I18nInfo()
				C.So(ok, C.ShouldEqual, false)
				C.So(m.NoReply(), C.ShouldEqual, false)
			})
		})
	})
}
