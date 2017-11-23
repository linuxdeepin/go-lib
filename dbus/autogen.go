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

package dbus

import "reflect"
import "errors"
import "fmt"
import "strings"
import "log"
import "pkg.deepin.io/lib/dbus/interfaces"
import "pkg.deepin.io/lib/dbus/introspect"

func splitObjectPath(path ObjectPath) (parent ObjectPath, base string) {
	if !path.IsValid() {
		return "", ""
	}
	i := strings.LastIndex(string(path), "/")
	switch i {
	case 0:
		return ObjectPath("/"), string(path)[1:]
	default:
		return ObjectPath(string(path)[:i]), string(path)[i+1:]
	}
}

func getTypeOf(ifc interface{}) (r reflect.Type) {
	r = reflect.TypeOf(ifc)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return
}

func getValueOf(ifc interface{}) (r reflect.Value) {
	r = reflect.ValueOf(ifc)
	if r.Kind() == reflect.Ptr {
		r = r.Elem()
	}
	return
}

func BuildInterfaceInfo(ifc interface{}) *introspect.InterfaceInfo {
	ifc_info := new(introspect.InterfaceInfo)
	o_type := reflect.TypeOf(ifc)
	n := o_type.NumMethod()

	for i := 0; i < n; i++ {
		method := o_type.Method(i)
		if method.PkgPath != "" {
			continue
		}
		name := method.Name
		if (name == "GetDBusInfo" && o_type.Implements(dbusObjectInterface)) || name == "OnPropertiesChanged" {
			continue
		}
		method_info := introspect.MethodInfo{}
		method_info.Name = name

		m := method.Type
		n_in := m.NumIn()
		n_out := m.NumOut()
		args := make([]introspect.ArgInfo, 0)
		//Method's first paramter is the struct which this method bound to.
		for i := 1; i < n_in; i++ {
			t := m.In(i)
			if i == 1 && t == dbusMessageType {
				continue
			}
			args = append(args, introspect.ArgInfo{
				Type:      SignatureOfType(t).String(),
				Direction: "in",
			})
		}
		for i := 0; i < n_out; i++ {
			t := m.Out(i)
			if t.Implements(goErrorType) {
				continue
			}
			args = append(args, introspect.ArgInfo{
				Type:      SignatureOfType(t).String(),
				Direction: "out",
			})
		}
		method_info.Args = args
		ifc_info.Methods = append(ifc_info.Methods, method_info)
	}

	// generate properties if any
	if o_type.Kind() == reflect.Ptr {
		o_type = o_type.Elem()
	}
	n = o_type.NumField()
	for i := 0; i < n; i++ {
		field := o_type.Field(i)
		if !isExportedStructField(field) {
			continue
		}
		if field.Type.Kind() == reflect.Func {
			ifc_info.Signals = append(ifc_info.Signals, introspect.SignalInfo{
				Name: field.Name,
				Args: func() []introspect.ArgInfo {
					n := field.Type.NumIn()
					ret := make([]introspect.ArgInfo, n)
					for i := 0; i < n; i++ {
						arg := field.Type.In(i)
						ret[i] = introspect.ArgInfo{
							Type: SignatureOfType(arg).String(),
						}
					}
					return ret
				}(),
			})
		} else if field.PkgPath == "" {
			access := field.Tag.Get("access")
			if access != "readwrite" {
				access = "read"
			}

			v := getValueOf(ifc).Field(i)
			if p, ok := v.Interface().(Property); ok {
				if p == nil {
					log.Println("UnInit dbus property", field.Name)
				} else {
					t := p.GetType()
					if t != nil {
						ifc_info.Properties = append(ifc_info.Properties, introspect.PropertyInfo{
							Name:   field.Name,
							Type:   SignatureOfType(t).String(),
							Access: access,
						})
					}
				}
			} else {
				ifc_info.Properties = append(ifc_info.Properties, introspect.PropertyInfo{
					Name:   field.Name,
					Type:   SignatureOfType(field.Type).String(),
					Access: access,
				})
			}
		}
	}

	return ifc_info
}

func InstallOnSession(obj DBusObject, ifcs ...interfaces.DBusInterface) error {
	conn, err := SessionBus()
	if err != nil {
		return err
	}
	return InstallOnAny(conn, obj, ifcs...)
}

func InstallOnSystem(obj DBusObject, ifcs ...interfaces.DBusInterface) error {
	conn, err := SystemBus()
	if err != nil {
		return err
	}
	return InstallOnAny(conn, obj, ifcs...)
}

func InstallOnAny(conn *Conn, obj DBusObject, ifcs ...interfaces.DBusInterface) error {
	if obj == nil {
		panic("Can't install an nil DBusObject to dbus")
	}
	if reflect.TypeOf(obj).Kind() != reflect.Ptr {
		panic("DBusObject must be an ptr at this moment")
	}
	return export(conn, obj, ifcs)
}

func UnInstallObject(obj DBusObject) {
	if c := detectConnByDBusObject(obj); c != nil {
		c.handlersLck.Lock()
		delete(c.handlers, ObjectPath(obj.GetDBusInfo().ObjectPath))
		c.handlersLck.Unlock()
	}
}

func ownerName(c *Conn, v interface{}, name string) error {
	if name != "." {
		not_registered := true
		for _, _name := range c.Names() {
			if _name == name {
				not_registered = false
				break
			}

		}
		if not_registered {
			reply, err := c.RequestName(name, NameFlagDoNotQueue)
			if err != nil {
				return err
			}
			if reply != RequestNameReplyPrimaryOwner {
				return errors.New("name " + name + " already taken")
			}
		}
	}
	return nil
}

func handleSubpath(c *Conn, path ObjectPath) {
	for path.IsValid() && path != ObjectPath("/") {
		parentpath, basepath := splitObjectPath(path)
		if parent, ok := c.handlers[ObjectPath(parentpath)]; ok {
			intro := parent[InterfaceIntrospectProxy]
			if reflect.TypeOf(intro).AssignableTo(introspectProxyType) {
				intro.(*IntrospectProxy).Enable(basepath)
			}
			return
		}
		path = parentpath
	}
}

//TODO: try to remove DBusObject
func export(c *Conn, v DBusObject, interfaces []interfaces.DBusInterface) error {
	dinfo := v.GetDBusInfo()
	path := ObjectPath(dinfo.ObjectPath)
	if !path.IsValid() {
		return fmt.Errorf("ObjectPath %q is invalid", dinfo.ObjectPath)
	}

	err := ownerName(c, v, dinfo.Dest)
	if err != nil {
		return err
	}

	err = c.Export(v, path, dinfo.Interface)
	if err != nil {
		return err
	}

	c.handlersLck.Lock()
	handleSubpath(c, path)
	c.handlersLck.Unlock()

	c.handlersLck.RLock()
	ifcs := c.handlers[path]
	c.handlersLck.RUnlock()

	interfaces = append(interfaces, NewIntrospectProxy(ifcs))
	interfaces = append(interfaces, NewPropertiesProxy(ifcs))

	for _, ifc := range interfaces {
		c.Export(ifc, path, ifc.InterfaceName())
	}

	return nil
}
