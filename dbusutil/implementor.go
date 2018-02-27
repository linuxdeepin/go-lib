package dbusutil

import (
	"errors"
	"strings"
	"sync"

	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbus1/introspect"
	"pkg.deepin.io/lib/dbus1/prop"
	"pkg.deepin.io/lib/strv"
)

type implementer struct {
	service       *Service
	path          dbus.ObjectPath
	interfaceName string
	core          Exportable
	corePropsMu   *sync.RWMutex // it pointer to core.PropsMu

	methods     []introspect.Method
	signals     []introspect.Signal
	props       map[string]*fieldProp
	propsMu     sync.RWMutex
	propChanges *propChanges
}

type propChanges struct {
	mu        sync.Mutex
	delayMode bool
	items     []implPropChanged
}

type implPropChanged struct {
	name  string
	value interface{}
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

func (impl *implementer) delayEmitPropChanged() {
	impl.propsMu.Lock()
	defer impl.propsMu.Unlock()

	if impl.propChanges == nil {
		impl.propChanges = &propChanges{}
	}

	impl.propChanges.mu.Lock()
	impl.propChanges.delayMode = true
}

func (impl *implementer) emitPropChanged(propName string, value interface{}) (err error) {
	impl.propsMu.Lock()
	defer impl.propsMu.Unlock()

	fieldProp, ok := impl.props[propName]
	if !ok {
		return errors.New("property not found")
	}
	if fieldProp.signature != dbus.SignatureOf(value) {
		return errors.New("property signature not equal")
	}

	if impl.propChanges != nil && impl.propChanges.delayMode {
		// in delay mode
		items := impl.propChanges.items
		found := false
		for i := 0; i < len(items); i++ {
			item := &items[i]
			if item.name == propName {
				item.value = value
				found = true
				break
			}
		}
		if !found {
			impl.propChanges.items = append(items, implPropChanged{
				name:  propName,
				value: value,
			})
		}

	} else {
		// not in delay mode
		err = fieldProp.emitChanged(impl.service.conn, impl.path, impl.interfaceName, value)
	}
	return
}

func (impl *implementer) stopDelayEmitPropChanged() (err error) {
	impl.propsMu.Lock()
	defer impl.propsMu.Unlock()

	if impl.propChanges == nil {
		panic("failed to assert impl.propChanges != nil")
	}

	var changedProps map[string]dbus.Variant
	var invalidatedProps []string
	items := impl.propChanges.items
	if len(items) > 0 {
		changedProps = make(map[string]dbus.Variant)
	}

	for _, change := range items {
		p := impl.props[change.name]
		switch p.emit {
		case emitFalse:
			// do nothing
		case emitInvalidates:
			invalidatedProps = append(invalidatedProps, change.name)
		case emitTrue:
			changedProps[change.name] = dbus.MakeVariant(change.value)
		}
	}

	const signalName = orgFreedesktopDBus + ".Properties.PropertiesChanged"
	if len(changedProps)+len(invalidatedProps) > 0 {
		err = impl.service.conn.Emit(impl.path, signalName, impl.interfaceName,
			changedProps, invalidatedProps)
	}

	impl.propChanges.items = nil
	impl.propChanges.delayMode = false
	impl.propChanges.mu.Unlock()
	return
}

type introspectableImplementer struct {
	service    *Service
	path       dbus.ObjectPath
	mu         sync.RWMutex
	data       string
	interfaces map[string]introspect.Interface
	//              ^interfaceName
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

func (oi *introspectableImplementer) clearCache() {
	oi.mu.Lock()
	oi.data = ""
	oi.mu.Unlock()
}

func (oi *introspectableImplementer) getChildren() (children strv.Strv) {
	var target string
	if oi.path == "/" {
		target = "/"
	} else {
		target = string(oi.path) + "/"
	}
	for objPath := range oi.service.objects {
		if objPath == oi.path {
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

	for _, child := range oi.getChildren() {
		node.Children = append(node.Children, introspect.Node{Name: child})
	}
	oi.mu.RUnlock()

	// marshal xml
	newData := string(introspect.NewIntrospectable(node))

	oi.mu.Lock()
	oi.data = newData
	oi.mu.Unlock()
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
