package dbus

import "reflect"
import "errors"
import "strings"

var (
	dbusErrorType = reflect.TypeOf(Error{})
)

func splitObjectPath(path ObjectPath) (parent, base string) {
	i := strings.LastIndex(string(path), "/")
	if i != -1 && i < len(string(path))-2 {
		return string(path)[:i], string(path)[i+1:]
	}
	return
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

func genInterfaceInfo(ifc interface{}) *InterfaceInfo {
	ifc_info := new(InterfaceInfo)
	o_type := reflect.TypeOf(ifc)
	n := o_type.NumMethod()

	for i := 0; i < n; i++ {
		name := o_type.Method(i).Name
		if (name == "GetDBusInfo" && o_type.Implements(dbusObjectInterface)) || name == "OnPropertiesChanged" {
			continue
		}
		method := MethodInfo{}
		method.Name = name

		m := o_type.Method(i).Type
		n_in := m.NumIn()
		n_out := m.NumOut()
		args := make([]ArgInfo, 0)
		//Method's first paramter is the struct which this method bound to.
		for i := 1; i < n_in; i++ {
			t := m.In(i)
			args = append(args, ArgInfo{
				Type:      SignatureOfType(t).String(),
				Direction: "in",
			})
		}
		for i := 0; i < n_out; i++ {
			t := m.Out(i)
			if t == reflect.TypeOf(dbusErrorType) {
				continue
			}
			args = append(args, ArgInfo{
				Type:      SignatureOfType(t).String(),
				Direction: "out",
			})
		}
		method.Args = args
		ifc_info.Methods = append(ifc_info.Methods, method)
	}

	// generate properties if any
	if o_type.Kind() == reflect.Ptr {
		o_type = o_type.Elem()
	}
	n = o_type.NumField()
	for i := 0; i < n; i++ {
		field := o_type.Field(i)
		if field.Type.Kind() == reflect.Func {
			ifc_info.Signals = append(ifc_info.Signals, SignalInfo{
				Name: field.Name,
				Args: func() []ArgInfo {
					n := field.Type.NumIn()
					ret := make([]ArgInfo, n)
					for i := 0; i < n; i++ {
						arg := field.Type.In(i)
						ret[i] = ArgInfo{
							Type: SignatureOfType(arg).String(),
						}
					}
					return ret
				}(),
			})
		} else if field.PkgPath == "" {
			access := field.Tag.Get("access")
			if access != "read" {
				access = "readwrite"
			}
			ifc_info.Properties = append(ifc_info.Properties, PropertyInfo{
				Name:   field.Name,
				Type:   SignatureOfType(field.Type).String(),
				Access: access,
			})
		}
	}

	return ifc_info
}

func InstallOnSession(obj DBusObject) error {
	info := obj.GetDBusInfo()
	path := ObjectPath(info.ObjectPath)
	if path.IsValid() {
		return InstallOnSessionAny(obj, info.Dest, path, info.Interface)
	}
	return errors.New("ObjectPath " + info.ObjectPath + " is invalid")
}

//TODO: Need exported?
func InstallOnSessionAny(v interface{}, dest_name string, path ObjectPath, iface string) error {
	conn, err := SessionBus()
	if err != nil {
		return err
	}
	return export(conn, v, dest_name, path, iface)
}

func setupSignalHandler(c *Conn, v interface{}, path ObjectPath, iface string) {
	value := reflect.ValueOf(v).Elem()
	n := value.NumField()
	for i := 0; i < n; i++ {
		fn := value.Field(i)
		if fn.Type().Kind() == reflect.Func {
			name := iface + "." + reflect.TypeOf(v).Elem().Field(i).Name
			fn.Set(reflect.MakeFunc(fn.Type(), func(in []reflect.Value) []reflect.Value {
				inputs := make([]interface{}, len(in))
				for i, v := range in {
					inputs[i] = v.Interface()
				}
				c.Emit(path, name, inputs...)
				return nil
			}))
		}
	}
}

//TODO: Need exported?
func export(c *Conn, v interface{}, name string, path ObjectPath, iface string) error {
	not_registered := true
	for _, _name := range c.Names() {
		if _name == name {
			not_registered = false
			break
		}

	}
	if not_registered {
		reply, _ := c.RequestName(name, NameFlagDoNotQueue)
		if reply != RequestNameReplyPrimaryOwner {
			return errors.New("name " + name + " already taken")
		}
	}
	setupSignalHandler(c, v, path, iface)

	err := c.Export(v, path, iface)
	if err != nil {
		return err
	}
	infos := c.handlers[path]
	parentpath, basepath := splitObjectPath(path)
	if parent, ok := c.handlers[ObjectPath(parentpath)]; ok {
		intro := parent["org.freedesktop.DBus.Introspectable"]
		if reflect.TypeOf(intro).AssignableTo(introspectProxyType) {
			intro.(IntrospectProxy).child[basepath] = true
		}
	}
	if _, ok := infos["org.freedesktop.DBus.Introspectable"]; !ok {
		infos["org.freedesktop.DBus.Introspectable"] = IntrospectProxy{infos, make(map[string]bool)}
	}
	if _, ok := infos["org.freedesktop.DBus.Properties"]; !ok {
		infos["org.freedesktop.DBus.Properties"] = PropertiesProxy{infos}
	}
	return nil
}
