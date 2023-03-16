// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
// 
// SPDX-License-Identifier: GPL-3.0-or-later
package proxy

import (
	"errors"
	"fmt"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/linuxdeepin/go-lib/dbusutilv1"
)

type Object interface {
	Path_() dbus.ObjectPath
	ServiceName_() string
	InitSignalExt(sigLoop *dbusutilv1.SignalLoop, ruleAuto bool)
	RemoveAllHandlers()
	RemoveHandler(handlerId dbusutilv1.SignalHandlerId)
	ConnectPropertiesChanged(
		cb func(interfaceName string, changedProperties map[string]dbus.Variant,
			invalidatedProperties []string)) (dbusutilv1.SignalHandlerId, error)
}

type ImplObject struct {
	obj      dbus.BusObject
	conn     *dbus.Conn
	mu       sync.Mutex
	extraMap map[string]interface{}
	*objectSignalExt
}

func (o *ImplObject) SetExtra(key string, v interface{}) {
	o.mu.Lock()

	if o.extraMap == nil {
		o.extraMap = make(map[string]interface{})
	}
	o.extraMap[key] = v

	o.mu.Unlock()
}

func (o *ImplObject) GetExtra(key string) (interface{}, bool) {
	o.mu.Lock()
	v, ok := o.extraMap[key]
	o.mu.Unlock()
	return v, ok
}

type objectSignalExt struct {
	sigLoop  *dbusutilv1.SignalLoop
	ruleAuto bool

	// used when ruleAuto is true
	ruleHandlersMap map[string][]dbusutilv1.SignalHandlerId //nolint
	//                  ^rule  ^handler ids

	// used when ruleAuto is false
	handleIds []dbusutilv1.SignalHandlerId //nolint:

	propChangedHandlerId dbusutilv1.SignalHandlerId               //nolint:
	propChangedCallbacks map[propChangedKey][]PropChangedCallback //nolint:
}

type propChangedKey struct {
	interfaceName, propName string
}

type PropChangedCallback func(hasValue bool, value interface{})

func (o *ImplObject) Init_(conn *dbus.Conn, serviceName string,
	path dbus.ObjectPath) {
	o.conn = conn
	o.obj = conn.Object(serviceName, path)
}

func (o *ImplObject) Path_() dbus.ObjectPath {
	return o.obj.Path()
}

func (o *ImplObject) ServiceName_() string {
	return o.obj.Destination()
}

func (o *ImplObject) Go_(method string, flags dbus.Flags,
	ch chan *dbus.Call, args ...interface{}) *dbus.Call {
	return o.obj.Go(method, flags, ch, args...)
}

func (o *ImplObject) GetProperty_(flags dbus.Flags, interfaceName, propName string,
	value interface{}) error {
	call := o.obj.Call("org.freedesktop.DBus.Properties.Get", flags, interfaceName, propName)
	return storeGetProperty(call, value)
}

func storeGetProperty(call *dbus.Call, value interface{}) error {
	var variant dbus.Variant
	err := call.Store(&variant)
	if err != nil {
		return err
	}
	return dbus.Store([]interface{}{variant.Value()}, value)
}

func (o *ImplObject) SetProperty_(flags dbus.Flags, interfaceName, propName string,
	value interface{}) error {
	return o.obj.Call("org.freedesktop.DBus.Properties.Set", flags, interfaceName,
		propName, dbus.MakeVariant(value)).Err
}

func (o *ImplObject) addRuleHandlerId(rule string, handlerId dbusutilv1.SignalHandlerId) {
	if o.ruleHandlersMap == nil {
		o.ruleHandlersMap = make(map[string][]dbusutilv1.SignalHandlerId)
	}
	o.ruleHandlersMap[rule] = append(o.ruleHandlersMap[rule], handlerId)
}

func (o *ImplObject) getMatchRulePropertiesChanged() string {
	return fmt.Sprintf("type='signal',interface='org.freedesktop.DBus.Properties',"+
		"member='PropertiesChanged',path='%s',sender='%s'", o.obj.Path(), o.obj.Destination())
}

func (o *ImplObject) propChangedHandler(sig *dbus.Signal) {
	var interfaceName string
	var changedProps map[string]dbus.Variant
	var invalidatedProps []string

	err := dbus.Store(sig.Body, &interfaceName, &changedProps, &invalidatedProps)
	if err != nil {
		return
	}
	for propName, variant := range changedProps {
		key := propChangedKey{interfaceName, propName}
		o.mu.Lock()
		callbacks := o.propChangedCallbacks[key]
		o.mu.Unlock()

		for _, callback := range callbacks {
			callback(true, variant.Value())
		}
	}

	for _, propName := range invalidatedProps {
		key := propChangedKey{interfaceName, propName}
		o.mu.Lock()
		callbacks := o.propChangedCallbacks[key]
		o.mu.Unlock()

		for _, callback := range callbacks {
			callback(false, nil)
		}
	}
}

func (o *ImplObject) initPropertiesChangedHandler() error {
	o.checkSignalExt()
	if o.propChangedHandlerId != 0 {
		return nil
	}
	var handlerId dbusutilv1.SignalHandlerId
	sigRule := &dbusutilv1.SignalRule{
		Path: o.obj.Path(),
		Name: "org.freedesktop.DBus.Properties.PropertiesChanged",
	}
	if o.ruleAuto {
		rule := o.getMatchRulePropertiesChanged()
		err := o.addMatch(rule)
		if err != nil {
			return err
		}
		handlerId = o.sigLoop.AddHandler(sigRule, o.propChangedHandler)
		o.addRuleHandlerId(rule, handlerId)
	} else {
		handlerId = o.sigLoop.AddHandler(sigRule, o.propChangedHandler)
	}
	o.propChangedHandlerId = handlerId
	return nil
}

func (o *ImplObject) ConnectPropertyChanged_(interfaceName, propName string, callback PropChangedCallback) error {
	o.mu.Lock()
	defer o.mu.Unlock()

	err := o.initPropertiesChangedHandler()
	if err != nil {
		return err
	}

	if o.propChangedCallbacks == nil {
		o.propChangedCallbacks = make(map[propChangedKey][]PropChangedCallback)
	}

	key := propChangedKey{interfaceName, propName}
	o.propChangedCallbacks[key] = append(o.propChangedCallbacks[key], callback)
	return nil
}

func (o *ImplObject) checkSignalExt() {
	if o.objectSignalExt == nil {
		panic(fmt.Sprintf("you should call %T.InitSignalExt() first", o))
	}
}

func (o *ImplObject) addMatch(rule string) error {
	handlerIds := o.ruleHandlersMap[rule]
	if len(handlerIds) == 0 {
		return globalRuleCounter.addMatch(o.conn, rule)
	}
	return nil
}

func handlerIdSliceDelete(slice []dbusutilv1.SignalHandlerId,
	index int) []dbusutilv1.SignalHandlerId {
	return append(slice[:index], slice[index+1:]...)
}

func handlerIdSliceIndex(slice []dbusutilv1.SignalHandlerId, hId dbusutilv1.SignalHandlerId) int {
	for idx, value := range slice {
		if value == hId {
			return idx
		}
	}
	return -1
}

func (o *ImplObject) removeHandler(handlerId dbusutilv1.SignalHandlerId) {
	o.sigLoop.RemoveHandler(handlerId)
	if o.ruleAuto {
		for rule, handlerIds := range o.ruleHandlersMap {
			for idx, hId := range handlerIds {
				if hId == handlerId {
					if len(handlerIds) == 1 {
						_ = globalRuleCounter.removeMatch(o.conn, rule)
					}
					o.ruleHandlersMap[rule] = handlerIdSliceDelete(handlerIds, idx)
					return
				}
			}
		}
	} else {
		idx := handlerIdSliceIndex(o.handleIds, handlerId)
		if idx != -1 {
			o.handleIds = handlerIdSliceDelete(o.handleIds, idx)
		}
	}
}

func (o *ImplObject) removePropertiesChangedHandler() {
	if o.propChangedHandlerId != 0 {
		o.removeHandler(o.propChangedHandlerId)
		o.propChangedHandlerId = 0
	}
}

func (o *ImplObject) removeAllHandlers() {
	if o.ruleAuto {
		for rule, handlerIds := range o.ruleHandlersMap {
			for _, hId := range handlerIds {
				o.sigLoop.RemoveHandler(hId)
			}
			_ = globalRuleCounter.removeMatch(o.conn, rule)
		}
		o.ruleHandlersMap = nil
	} else {
		for _, hId := range o.handleIds {
			o.sigLoop.RemoveHandler(hId)
		}
	}
	o.removePropertiesChangedHandler()
}

func (o *ImplObject) ConnectSignal_(rule string, sigRule *dbusutilv1.SignalRule, cb dbusutilv1.SignalHandlerFunc) (dbusutilv1.SignalHandlerId, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.checkSignalExt()
	return o.connectSignal(rule, sigRule, cb)
}

func (o *ImplObject) connectSignal(rule string, sigRule *dbusutilv1.SignalRule, cb dbusutilv1.SignalHandlerFunc) (dbusutilv1.SignalHandlerId, error) {
	if o.ruleAuto {
		err := o.addMatch(rule)
		if err != nil {
			return 0, err
		}

		handlerId := o.sigLoop.AddHandler(sigRule, cb)
		o.addRuleHandlerId(rule, handlerId)
		return handlerId, nil

	} else {
		handlerId := o.sigLoop.AddHandler(sigRule, cb)
		o.handleIds = append(o.handleIds, handlerId)
		return handlerId, nil
	}
}

// Object public methods:

func (o *ImplObject) InitSignalExt(sigLoop *dbusutilv1.SignalLoop, ruleAuto bool) {
	if o.conn != sigLoop.Conn() {
		panic("dbus conn not same")
	}
	o.mu.Lock()
	if o.objectSignalExt == nil {
		o.objectSignalExt = &objectSignalExt{
			sigLoop:  sigLoop,
			ruleAuto: ruleAuto,
		}
	}
	o.mu.Unlock()
}

const (
	RemoveAllHandlers              = -1
	RemovePropertiesChangedHandler = -2
)

func (o *ImplObject) RemoveAllHandlers() {
	o.removeAllHandlers()
}

func (o *ImplObject) RemoveHandler(handlerId dbusutilv1.SignalHandlerId) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.checkSignalExt()

	switch handlerId {
	case RemoveAllHandlers:
		o.removeAllHandlers()
	case RemovePropertiesChangedHandler:
		o.removePropertiesChangedHandler()
	default:
		o.removeHandler(handlerId)
	}
}

func (o *ImplObject) ConnectPropertiesChanged(
	cb func(interfaceName string, changedProperties map[string]dbus.Variant,
		invalidatedProperties []string)) (dbusutilv1.SignalHandlerId, error) {

	if cb == nil {
		return 0, errors.New("nil callback")
	}
	var rule string

	o.mu.Lock()
	defer o.mu.Unlock()

	o.checkSignalExt()
	if o.ruleAuto {
		rule = o.getMatchRulePropertiesChanged()
	}

	return o.connectSignal(rule, &dbusutilv1.SignalRule{
		Path: o.obj.Path(),
		Name: "org.freedesktop.DBus.Properties.PropertiesChanged",
	}, func(sig *dbus.Signal) {
		var interfaceName string
		var changedProps map[string]dbus.Variant
		var invalidatedProps []string

		err := dbus.Store(sig.Body, &interfaceName, &changedProps, &invalidatedProps)
		if err == nil {
			cb(interfaceName, changedProps, invalidatedProps)
		}
	})
}

type Implementer interface {
	GetObject_() *ImplObject
	GetInterfaceName_() string
}
