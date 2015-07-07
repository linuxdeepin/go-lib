package dbus

import "encoding/xml"
import "bytes"
import "fmt"
import "reflect"
import "sync"
import "pkg.deepin.io/lib/dbus/introspect"

const InterfaceLifeManager = "com.deepin.DBus.LifeManager"

type LifeManager struct {
	name        string
	path        ObjectPath
	count       int32
	onDestroy   func()
	countLocker sync.Mutex
}

func RegisterLifeManagerCallback(m DBusObject, cb func()) {
	if con := detectConnByDBusObject(m); con != nil {
		path := ObjectPath(m.GetDBusInfo().ObjectPath)
		if infos, ok := con.handlers[path]; ok {
			if i, ok := (infos[InterfaceLifeManager]).(*LifeManager); ok {
				i.onDestroy = cb
			}
		}
	}
}

func NewLifeManager(name string, path ObjectPath) *LifeManager {
	return &LifeManager{name: name, path: path, count: 1}
}

func (ifc *LifeManager) InterfaceName() string {
	return InterfaceLifeManager
}

func (ifc *LifeManager) Ref() {
	ifc.countLocker.Lock()

	ifc.count++

	ifc.countLocker.Unlock()
}
func (ifc *LifeManager) Unref() {
	ifc.countLocker.Lock()

	ifc.count--
	if ifc.count == 0 && ifc.onDestroy != nil {
		ifc.onDestroy()
	}

	ifc.countLocker.Unlock()
}

const InterfaceIntrospectProxy = "org.freedesktop.DBus.Introspectable"

type IntrospectProxy struct {
	infos map[string]interface{}
	child map[string]bool
}

func NewIntrospectProxy(infos map[string]interface{}) *IntrospectProxy {
	return &IntrospectProxy{
		infos: infos,
		child: make(map[string]bool),
	}
}

func (ifc IntrospectProxy) InterfaceName() string {
	return InterfaceIntrospectProxy
}

func (ifc IntrospectProxy) Introspect() (string, error) {
	var node = new(introspect.NodeInfo)
	for k, _ := range ifc.child {
		node.Children = append(node.Children, introspect.NodeInfo{
			Name: k,
		})
	}
	for name, ifc := range ifc.infos {
		info := BuildInterfaceInfo(ifc)
		info.Name = name
		node.Interfaces = append(node.Interfaces, *info)
	}
	var buffer bytes.Buffer

	writer := xml.NewEncoder(&buffer)
	writer.Indent("", "     ")
	writer.Encode(node)
	return buffer.String(), nil
}

const InterfacePropertiesProxy = "org.freedesktop.DBus.Properties"

type PropertiesProxy struct {
	infos             map[string]interface{}
	PropertiesChanged func(string, map[string]Variant, []string)
}

func NewPropertiesProxy(infos map[string]interface{}) *PropertiesProxy {
	return &PropertiesProxy{
		infos: infos,
	}
}

func (PropertiesProxy) InterfaceName() string {
	return InterfacePropertiesProxy
}

func (propProxy PropertiesProxy) GetAll(ifcName string) (props map[string]Variant, err error) {
	props = make(map[string]Variant)
	if ifc, ok := propProxy.infos[ifcName]; ok {
		o_type := getTypeOf(ifc)
		n := o_type.NumField()
		for i := 0; i < n; i++ {
			field := o_type.Field(i)
			if field.Type.Kind() != reflect.Func && isExportedStructField(field) {
				props[field.Name], err = propProxy.Get(ifcName, field.Name)
				if err != nil {
					return nil, err
				}
			}
		}
	}
	return
}

func (propProxy PropertiesProxy) Set(ifcName string, propName string, value Variant) error {
	if ifc, ok := propProxy.infos[ifcName]; ok {
		ifc_t := getTypeOf(ifc)
		t, ok := ifc_t.FieldByName(propName)
		v := getValueOf(ifc).FieldByName(propName)
		if ok && v.IsValid() && "read" != t.Tag.Get("access") {
			if !v.CanAddr() {
				return NewPropertyNotWritableError(propName)
			}
			if v.Type().Implements(propertyType) {
				if reflect.TypeOf(value.Value()) == v.MethodByName("GetType").Interface().(func() reflect.Type)() {
					v.MethodByName("SetValue").Interface().(func(interface{}))(value.Value())
					fn := reflect.ValueOf(ifc).MethodByName("OnPropertiesChanged")
					if fn.IsValid() && !fn.IsNil() {
						fn.Call([]reflect.Value{reflect.ValueOf(propName), reflect.Zero(reflect.TypeOf(value.Value()))})
					}
					return nil
				} else {
					return NewPropertyNotWritableError(propName)
				}
			} else if v.Type() == reflect.TypeOf(value.Value()) {
				prop_val := reflect.ValueOf(value.Value())
				prop_old_val := v.Interface()
				v.Set(prop_val)
				fn := reflect.ValueOf(ifc).MethodByName("OnPropertiesChanged")
				if fn.IsValid() && !fn.IsNil() {
					fn.Call([]reflect.Value{reflect.ValueOf(propName), reflect.ValueOf(prop_old_val)})
				}
				return nil
			} else if isStructureMatched(v.Interface(), value.Value()) {
				t := reflect.TypeOf(v.Interface())
				if t.Kind() == reflect.Ptr {
					t = t.Elem()
					v = v.Elem()
				}
				dValues := value.Value().([]interface{})
				for i := 0; i < t.NumField(); i++ {
					if isExportedStructField(t.Field(i)) {
						field := v.Field(i)
						if t.Field(i).Type.Kind() == reflect.Ptr {
							field = field.Elem()
						}
						field.Set(reflect.ValueOf(dValues[i]))
					}
				}
				return nil
			}
		} else {
			return NewUnknowPropertyError(propName)
		}
	}
	return NewUnknowInterfaceError(ifcName)
}
func (propProxy PropertiesProxy) Get(ifcName string, propName string) (Variant, error) {
	if ifc, ok := propProxy.infos[ifcName]; ok {
		t, ok := getTypeOf(ifc).FieldByName(propName)
		if !ok || !isExportedStructField(t) {
			return MakeVariant(""), NewUnknowPropertyError(propName)
		}
		value := getValueOf(ifc).FieldByName(propName)
		if value.IsValid() == false {
			return MakeVariant(""), NewUnknowPropertyError(propName)
		}

		if value.Type().Implements(propertyType) {
			if value.IsNil() {
				return MakeVariant(""), fmt.Errorf("nil dbus.Property(%s:%s)", ifcName, propName)
			} else if fun, ok := value.MethodByName("GetValue").Interface().(func() interface{}); ok {
				t := fun()
				return MakeVariant(t), nil
			} else {
				return MakeVariant(""), NewUnknowPropertyError(propName)
			}
		} else if reflect.TypeOf(ifc).Implements(dbusObjectInterface) {
			value = tryTranslateDBusObjectToObjectPath(detectConnByDBusObject(ifc.(DBusObject)), value)
			//TODO: if ifc is not an DBusObject then we need try get the Conn object by other way
		}

		if value.IsValid() {
			if path, ok := value.Interface().(*ObjectPath); ok {
				if path == nil || !path.IsValid() {
					return MakeVariant(""), NewUnknowPropertyError(propName)
				}
			} else if path, ok := value.Interface().(ObjectPath); ok {
				if !path.IsValid() {
					return MakeVariant(""), NewUnknowPropertyError(propName)
				}
			} else if str, ok := value.Interface().(*string); ok && str == nil {
				//TODO: Why only an nil ptr string will cause we lost dbus connection?
				return MakeVariant(""), nil
			}
			return MakeVariant(value.Interface()), nil
		} else {
			return MakeVariant(""), NewUnknowPropertyError(propName)
		}
	} else {
		return MakeVariant(""), NewUnknowInterfaceError(ifcName)
	}
}
