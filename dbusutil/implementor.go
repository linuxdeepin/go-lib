package dbusutil

import (
	"errors"
	"log"
	"sync"

	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbus1/introspect"
	"pkg.deepin.io/lib/dbus1/prop"
	"pkg.deepin.io/lib/strv"
)

type implementer struct {
	service         *Service
	path            dbus.ObjectPath
	interfaceName   string
	core            Exportable
	corePropsLocker propLocker // it pointer to core.PropsMaster

	methods []introspect.Method
	props   map[string]*fieldProp
	propsMu sync.RWMutex

	signals []introspect.Signal
}

func (impl *implementer) Service() *Service {
	return impl.service
}

func (impl *implementer) Path() dbus.ObjectPath {
	return impl.path
}

func (impl *implementer) Interface() string {
	return impl.interfaceName
}

func (impl *implementer) Core() Exportable {
	return impl.core
}

func (impl *implementer) setWriteCallback(propertyName string, cb PropertyWriteCallback) error {
	prop0 := impl.getProperty(propertyName)
	if prop0 == nil {
		return errors.New("property not found")
	}
	prop0.setWriteCallback(cb)
	return nil
}

func (impl *implementer) setReadCallback(propertyName string, cb PropertyReadCallback) error {
	prop0 := impl.getProperty(propertyName)
	if prop0 == nil {
		return errors.New("property not found")
	}
	prop0.setReadCallback(cb)
	return nil
}

func (impl *implementer) connectChanged(propertyName string, cb PropertyChangedCallback) error {
	prop0 := impl.getProperty(propertyName)
	if prop0 == nil {
		return errors.New("property not found")
	}

	prop0.connectChanged(cb)
	return nil
}

func (impl *implementer) NotifyChanged(propertyName string, value interface{}) error {
	prop0 := impl.getProperty(propertyName)
	if prop0 == nil {
		return errors.New("not found property")
	}
	impl.notifyChanged(prop0, value)
	return nil
}

func (impl *implementer) notifyChanged(prop0 *fieldProp, value interface{}) {
	propChanged := newPropertyChanged(prop0.name, impl, value)
	prop0.notifyChanged(propChanged)
	prop0.emitChanged(impl.service.conn, impl.path, impl.interfaceName, value)
}

func (impl *implementer) getProperty(name string) *fieldProp {
	impl.propsMu.RLock()
	p := impl.props[name]
	impl.propsMu.RUnlock()
	return p
}

func (impl *implementer) checkProperty(name string) error {
	prop0 := impl.getProperty(name)
	if prop0 == nil {
		return errors.New("property not found")
	}
	return nil
}

func (impl *implementer) checkPropertyValue(name string, value interface{}) error {
	prop0 := impl.getProperty(name)
	if prop0 == nil {
		return errors.New("not found property")
	}

	if prop0.signature != dbus.SignatureOf(value) {
		return errors.New("property signature not equal")
	}

	return nil
}

type introspectableImplementer struct {
	service    *Service
	path       dbus.ObjectPath
	mu         sync.RWMutex
	data       string
	interfaces map[string]introspect.Interface
	//              ^interfaceName
	children strv.Strv // node name of children
}

var peerIntrospectData = introspect.Interface{
	Name: orgFreedesktopDBus + ".Peer",
	Methods: []introspect.Method{
		{
			Name: "Ping",
		},
		{
			Name: "GetMachineId",
			Args: []introspect.Arg{
				{"machine_uuid", "s", "out"},
			},
		},
	},
}

func (oi *introspectableImplementer) addImplementer(impl *implementer) {
	oi.mu.Lock()

	empty := len(oi.interfaces) == 0
	oi.interfaces[impl.interfaceName] = introspect.Interface{
		Name:       impl.interfaceName,
		Methods:    impl.methods,
		Properties: getPropsIntrospection(impl.props),
		Signals:    impl.signals,
	}
	if empty {
		oi.interfaces[peerIntrospectData.Name] = peerIntrospectData
		oi.interfaces[introspect.IntrospectData.Name] = introspect.IntrospectData
		oi.interfaces[prop.IntrospectData.Name] = prop.IntrospectData
	}
	oi.data = ""

	oi.mu.Unlock()
}

func (oi *introspectableImplementer) deleteImplementer(interfaceName string) {
	oi.mu.Lock()
	delete(oi.interfaces, interfaceName)
	oi.data = ""
	oi.mu.Unlock()
}

func (oi *introspectableImplementer) addChild(child string) {
	oi.mu.Lock()
	oi.children, _ = oi.children.Add(child)
	oi.data = ""
	oi.mu.Unlock()
}

func (oi *introspectableImplementer) deleteChild(child string) {
	oi.mu.Lock()
	oi.children, _ = oi.children.Delete(child)
	oi.data = ""
	oi.mu.Unlock()
}

func (oi *introspectableImplementer) Introspect() (string, *dbus.Error) {
	oi.service.DelayAutoQuit()
	oi.mu.RLock()
	data := oi.data

	if data != "" {
		oi.mu.RUnlock()
		return data, nil
	}

	node := &introspect.Node{
		Interfaces: oi.getInterfaces(),
	}

	for _, child := range oi.children {
		node.Children = append(node.Children, introspect.Node{Name: child})
	}
	oi.mu.RUnlock()

	// marshal xml
	newData := string(introspect.NewIntrospectable(node))

	oi.mu.Lock()
	oi.data = newData
	oi.mu.Unlock()
	log.Println("encoding xml", oi.path)
	return newData, nil
}

func (oi *introspectableImplementer) getInterfaces() []introspect.Interface {
	interfaces := make([]introspect.Interface, len(oi.interfaces))
	idx := 0
	for _, ifc := range oi.interfaces {
		interfaces[idx] = ifc
		idx++
	}
	return interfaces
}

type propertiesImplementer struct {
	object  *object
	service *Service
}

func (p *propertiesImplementer) Get(sender dbus.Sender,
	interfaceName, propertyName string) (dbus.Variant, *dbus.Error) {
	p.service.DelayAutoQuit()

	impl := p.object.getImplementer(interfaceName)
	if impl == nil {
		return dbus.Variant{}, prop.ErrIfaceNotFound
	}
	prop0 := impl.getProperty(propertyName)
	if prop0 == nil {
		return dbus.Variant{}, prop.ErrPropNotFound
	}

	if prop0.access&accessRead == 0 {
		return dbus.Variant{}, dbus.MakeFailedError(errors.New("property can not be read"))
	}

	propRead := newPropertyRead(propertyName, impl, sender)

	variant, err := prop0.GetValueVariant(propRead)
	if err != nil {
		return dbus.Variant{}, err
	}
	return variant, nil
}

func (p *propertiesImplementer) GetAll(sender dbus.Sender, interfaceName string) (map[string]dbus.Variant, *dbus.Error) {
	p.service.DelayAutoQuit()

	impl := p.object.getImplementer(interfaceName)
	if impl == nil {
		return nil, prop.ErrIfaceNotFound
	}

	result := make(map[string]dbus.Variant, len(impl.props))
	impl.propsMu.RLock()
	for propName, prop0 := range impl.props {
		if prop0.access&accessRead != 0 {
			propRead := newPropertyRead(propName, impl, sender)
			variant, err := prop0.GetValueVariant(propRead)
			if err != nil {
				// ignore err
				continue
			}
			result[propName] = variant
		}
	}
	impl.propsMu.RUnlock()
	return result, nil
}

func (p *propertiesImplementer) Set(sender dbus.Sender, interfaceName, propertyName string,
	newVar dbus.Variant) *dbus.Error {
	p.service.DelayAutoQuit()

	impl := p.object.getImplementer(interfaceName)
	if impl == nil {
		return prop.ErrIfaceNotFound
	}

	prop0 := impl.getProperty(propertyName)
	if prop0 == nil {
		return prop.ErrPropNotFound
	}

	if prop0.access&accessWrite == 0 {
		return dbus.MakeFailedError(errors.New("property can not be written"))
	}

	if newVar.Signature() != prop0.signature {
		return prop.ErrInvalidArg
	}
	newVarValue := newVar.Value()

	// fix newVarValue []interface{} to struct
	if prop0.hasStruct {
		fixedRV, err := valueFromBus(newVarValue, prop0.rType)
		if err != nil {
			return dbus.MakeFailedError(err)
		}
		newVarValue = fixedRV.Interface()
	}

	propWrite := newPropertyWrite(propertyName, impl, newVarValue, sender)

	changed, setErr := prop0.SetValue(propWrite)
	if setErr != nil {
		return setErr
	}
	if changed {
		impl.notifyChanged(prop0, newVarValue)
	}
	return nil
}
