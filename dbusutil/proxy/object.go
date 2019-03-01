package proxy

import (
	"errors"
	"fmt"
	"sync"

	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbusutil"
)

type Object struct {
	obj  dbus.BusObject
	conn *dbus.Conn
	mu   sync.Mutex

	*objectSignalExt
}

type objectSignalExt struct {
	sigLoop         *dbusutil.SignalLoop
	ruleAuto        bool
	ruleHandlersMap map[string][]dbusutil.SignalHandlerId
	//                  ^rule  ^handler ids
	propChangedHandlerId dbusutil.SignalHandlerId
	propChangedCallbacks map[propChangedKey][]PropChangedCallback
}

type propChangedKey struct {
	interfaceName, propName string
}

type PropChangedCallback func(hasValue bool, value interface{})

func (o *Object) Init_(conn *dbus.Conn, serviceName string,
	path dbus.ObjectPath) {
	o.conn = conn
	o.obj = conn.Object(serviceName, path)
}

func (o *Object) Path_() dbus.ObjectPath {
	return o.obj.Path()
}

func (o *Object) ServiceName_() string {
	return o.obj.Destination()
}

func (o *Object) Go_(method string, flags dbus.Flags,
	ch chan *dbus.Call, args ...interface{}) *dbus.Call {
	return o.obj.Go(method, flags, ch, args...)
}

func (o *Object) GetProperty_(flags dbus.Flags, interfaceName, propName string,
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

func (o *Object) SetProperty_(flags dbus.Flags, interfaceName, propName string,
	value interface{}) error {
	return o.obj.Call("org.freedesktop.DBus.Properties.Set", flags, interfaceName,
		propName, dbus.MakeVariant(value)).Err
}

func (o *Object) addRuleHandlerId(rule string, handlerId dbusutil.SignalHandlerId) {
	if o.ruleHandlersMap == nil {
		o.ruleHandlersMap = make(map[string][]dbusutil.SignalHandlerId)
	}
	o.ruleHandlersMap[rule] = append(o.ruleHandlersMap[rule], handlerId)
}

func (o *Object) getMatchRulePropertiesChanged() string {
	return fmt.Sprintf("type='signal',interface='org.freedesktop.DBus.Properties',"+
		"member='PropertiesChanged',path='%s',sender='%s'", o.obj.Path(), o.obj.Destination())
}

func (o *Object) propChangedHandler(sig *dbus.Signal) {
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

func (o *Object) initPropertiesChangedHandler() error {
	o.checkSignalExt()
	if o.propChangedHandlerId != 0 {
		return nil
	}
	var handlerId dbusutil.SignalHandlerId
	sigRule := &dbusutil.SignalRule{
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

func (o *Object) ConnectPropertyChanged_(interfaceName, propName string, callback PropChangedCallback) error {
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

func (o *Object) checkSignalExt() {
	if o.objectSignalExt == nil {
		panic(fmt.Sprintf("you should call %T.InitSignalExt() first", o))
	}
}

func (o *Object) addMatch(rule string) error {
	handlerIds := o.ruleHandlersMap[rule]
	if len(handlerIds) == 0 {
		return globalRuleCounter.addMatch(o.conn, rule)
	}
	return nil
}

func handlerIdSliceDelete(slice []dbusutil.SignalHandlerId,
	index int) []dbusutil.SignalHandlerId {
	return append(slice[:index], slice[index+1:]...)
}

func (o *Object) removeHandler(handlerId dbusutil.SignalHandlerId) {
	for rule, handlerIds := range o.ruleHandlersMap {
		for idx, hId := range handlerIds {
			if hId == handlerId {
				o.sigLoop.RemoveHandler(hId)
				if len(handlerIds) == 1 {
					globalRuleCounter.removeMatch(o.conn, rule)
				}
				o.ruleHandlersMap[rule] = handlerIdSliceDelete(handlerIds, idx)
				return
			}
		}
	}
}

func (o *Object) removePropertiesChangedHandler() {
	if o.propChangedHandlerId != 0 {
		o.removeHandler(o.propChangedHandlerId)
		o.propChangedHandlerId = 0
	}
}

func (o *Object) removeAllHandlers() {
	for rule, handlerIds := range o.ruleHandlersMap {
		for _, hId := range handlerIds {
			o.sigLoop.RemoveHandler(hId)
		}
		globalRuleCounter.removeMatch(o.conn, rule)
	}
	o.ruleHandlersMap = nil
	o.propChangedHandlerId = 0
}

func (o *Object) ConnectSignal_(rule string, sigRule *dbusutil.SignalRule, cb dbusutil.SignalHandlerFunc) (dbusutil.SignalHandlerId, error) {
	o.mu.Lock()
	defer o.mu.Unlock()

	o.checkSignalExt()
	return o.connectSignal(rule, sigRule, cb)
}

func (o *Object) connectSignal(rule string, sigRule *dbusutil.SignalRule, cb dbusutil.SignalHandlerFunc) (dbusutil.SignalHandlerId, error) {
	if !o.ruleAuto {
		return o.sigLoop.AddHandler(sigRule, cb), nil
	}
	err := o.addMatch(rule)
	if err != nil {
		return 0, err
	}

	handlerId := o.sigLoop.AddHandler(sigRule, cb)
	o.addRuleHandlerId(rule, handlerId)
	return handlerId, nil
}

// Object public methods:

func (o *Object) InitSignalExt(sigLoop *dbusutil.SignalLoop, ruleAuto bool) {
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

func (o *Object) RemoveHandler(handlerId dbusutil.SignalHandlerId) {
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

func (o *Object) ConnectPropertiesChanged(
	cb func(interfaceName string, changedProperties map[string]dbus.Variant,
		invalidatedProperties []string)) (dbusutil.SignalHandlerId, error) {

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

	return o.connectSignal(rule, &dbusutil.SignalRule{
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
	GetObject_() *Object
	GetInterfaceName_() string
}
