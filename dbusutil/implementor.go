// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusutil

import (
	"errors"
	"reflect"
	"sync"
	"unsafe"

	"github.com/godbus/dbus/v5"
)

// Implementer 对象实例
type Implementer interface {
}

// ImplementerExt 对象实例拓展
type ImplementerExt interface {
	Implementer
	GetExportedMethods() ExportedMethods
}

// ImplementerV20 兼容v20之前的方式
type ImplementerV20 interface {
	GetInterfaceName() string
}

type ExportedMethods []ExportedMethod

func (methods ExportedMethods) Len() int {
	return len(methods)
}

func (methods ExportedMethods) Less(i, j int) bool {
	return methods[i].Name < methods[j].Name
}

func (methods ExportedMethods) Swap(i, j int) {
	methods[i], methods[j] = methods[j], methods[i]
}

func (methods ExportedMethods) toMethodTable() map[string]interface{} {
	result := make(map[string]interface{}, len(methods))
	for _, method := range methods {
		result[method.Name] = method.Fn
	}
	return result
}

type ExportedMethod struct {
	Name    string
	Fn      interface{}
	InArgs  []string
	OutArgs []string
}

// #nosec G103
func getImplementerPointer(impl Implementer) unsafe.Pointer {
	return unsafe.Pointer(reflect.ValueOf(impl).Pointer())
}

type implementer struct {
	core          Implementer
	props         map[string]*fieldProp
	propChanges   propChanges
	interfaceName string
}

func (impl *implementer) getStatic(s *Service) *implementerStatic {
	return s.getImplementerStatic(impl.getInterfaceName())
}

func (impl *implementer) getInterfaceName() string {
	return impl.interfaceName
}

func newImplementer(core Implementer, ifcName string, service *Service, path dbus.ObjectPath) (*implementer, error) {
	impl := &implementer{
		core:          core,
		interfaceName: ifcName,
	}

	structValue, ok := getStructValue(core)
	if !ok {
		return nil, errors.New("v is not a struct pointer")
	}

	service.mu.Lock()
	implStatic, ok := service.implStaticMap[ifcName]
	if !ok {
		implStatic = newImplementerStatic(core, ifcName, structValue)
		service.implStaticMap[ifcName] = implStatic
	}
	service.mu.Unlock()

	impl.props = getFieldPropMap(impl, implStatic, structValue, service, path)
	return impl, nil
}

type propChanges struct {
	mu        sync.Mutex
	modeMu    sync.Mutex
	delayMode bool
	items     []implPropChanged
}

type implPropChanged struct {
	name  string
	value interface{}
}

func (impl *implementer) setWriteCallback(propertyName string, cb PropertyWriteCallback) error {
	p := impl.props[propertyName]
	if p == nil {
		return errors.New("property not found")
	}
	p.setWriteCallback(cb)
	return nil
}

func (impl *implementer) setReadCallback(propertyName string, cb PropertyReadCallback) error {
	p := impl.props[propertyName]
	if p == nil {
		return errors.New("property not found")
	}
	p.setReadCallback(cb)
	return nil
}

func (impl *implementer) connectChanged(propertyName string, cb PropertyChangedCallback) error {
	p := impl.props[propertyName]
	if p == nil {
		return errors.New("property not found")
	}

	p.connectChanged(cb)
	return nil
}

func (impl *implementer) notifyChanged(s *Service, path dbus.ObjectPath,
	p *fieldProp, propStatic *fieldPropStatic, value interface{}) {

	interfaceName := impl.getInterfaceName()
	propChanged := newPropertyChanged(path, interfaceName, propStatic.name, value)
	p.notifyChanged(propChanged)
	_ = emitPropertiesChanged(s.conn, path, interfaceName,
		propStatic.name, value, propStatic.emit)
}

func (impl *implementer) delayEmitPropChanged() {
	// start delay mode
	impl.propChanges.modeMu.Lock()

	impl.propChanges.mu.Lock()
	impl.propChanges.delayMode = true
	impl.propChanges.mu.Unlock()
}

func (impl *implementer) emitPropChanged(s *Service, path dbus.ObjectPath,
	propName string, value interface{}) (err error) {
	impl.propChanges.mu.Lock()
	defer impl.propChanges.mu.Unlock()

	if impl.propChanges.delayMode {
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
		implStatic := impl.getStatic(s)
		propStatic := implStatic.props[propName]
		err = emitPropertiesChanged(s.conn, path, impl.getInterfaceName(),
			propName, value, propStatic.emit)
	}
	return
}

func (impl *implementer) stopDelayEmitPropChanged(s *Service, path dbus.ObjectPath) (err error) {
	impl.propChanges.mu.Lock()

	var changedProps map[string]dbus.Variant
	var invalidatedProps []string
	items := impl.propChanges.items
	if len(items) > 0 {
		changedProps = make(map[string]dbus.Variant)
	}

	implStatic := impl.getStatic(s)

	for _, change := range items {
		p := implStatic.props[change.name]
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
		err = s.conn.Emit(path, signalName, impl.getInterfaceName(),
			changedProps, invalidatedProps)
	}

	impl.propChanges.items = nil
	impl.propChanges.delayMode = false
	impl.propChanges.mu.Unlock()

	// end delay mode
	impl.propChanges.modeMu.Unlock()
	return
}
