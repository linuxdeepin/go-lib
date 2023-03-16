// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusutil

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
	"sync"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

var logger *log.Logger

func init() {
	// setup logger
	logOut := ioutil.Discard
	if os.Getenv("DEBUG_DBUSUTIL") == "1" {
		logOut = os.Stderr
	}
	logger = log.New(logOut, "[dbusutil]", log.Lshortfile)
}

const orgFreedesktopDBus = "org.freedesktop.DBus"

type accessType uint

const (
	accessRead      accessType = 1
	accessWrite                = 2
	accessReadWrite            = accessRead | accessWrite
)

func (a accessType) String() string {
	switch a {
	case accessRead:
		return "read"
	case accessWrite:
		return "write"
	case accessReadWrite:
		return "readwrite"
	default:
		return fmt.Sprintf("invalid(%d)", a)
	}
}

type emitType uint

const (
	emitFalse emitType = iota
	emitTrue
	emitInvalidates
)

func (e emitType) String() string {
	switch e {
	case emitFalse:
		return "false"
	case emitTrue:
		return "true"
	case emitInvalidates:
		return "invalidates"
	default:
		return fmt.Sprintf("invalid(%d)", e)
	}
}

// struct field prop
type fieldProp struct {
	rValue  reflect.Value
	valueMu *sync.RWMutex

	cbMu       sync.Mutex
	writeCb    PropertyWriteCallback
	readCb     PropertyReadCallback
	changedCbs []PropertyChangedCallback
}

type fieldPropStatic struct {
	name      string
	rType     reflect.Type
	valueType fieldPropValueType
	signature dbus.Signature
	hasStruct bool
	emit      emitType
	access    accessType
}

type fieldPropValueType uint

const (
	fieldPropValueNotProp fieldPropValueType = iota
	fieldPropValueImplProp
	fieldPropValuePtrImplProp
)

func (p *fieldProp) getValue(propRead *PropertyRead) (value interface{}, err *dbus.Error) {
	readCb := p.getReadCallback()
	if readCb != nil {
		err = readCb(propRead)
		if err != nil {
			return
		}
	}

	if p.valueMu != nil {
		p.valueMu.RLock()
	}

	value = p.rValue.Interface()
	if propValue, ok := value.(Property); ok {
		value, err = propValue.GetValue()
	}

	if p.valueMu != nil {
		p.valueMu.RUnlock()
	}
	return
}

func (p *fieldProp) GetValueVariant(propRead *PropertyRead,
	signature dbus.Signature) (dbus.Variant, *dbus.Error) {

	value, err := p.getValue(propRead)
	if err != nil {
		return dbus.Variant{}, err
	}
	return dbus.MakeVariantWithSignature(value, signature), nil
}

func (p *fieldProp) SetValue(propWrite *PropertyWrite) (changed bool, err *dbus.Error) {
	writeCb := p.getWriteCallback()
	if writeCb != nil {
		err = writeCb(propWrite)
		if err != nil {
			return
		}
	}

	if p.valueMu != nil {
		p.valueMu.Lock()
	}

	newVal := propWrite.Value

	value := p.rValue.Interface()
	propValue, ok := value.(Property)
	if ok {
		changed, err = propValue.SetValue(newVal)
	} else {
		newValRV := reflect.ValueOf(newVal)
		newValRT := reflect.TypeOf(newVal)
		valueRT := reflect.TypeOf(value)
		if valueRT != newValRT {
			// type not equal, try convert
			if newValRT.ConvertibleTo(valueRT) {
				newValRV = newValRV.Convert(valueRT)
			} else {
				err = dbus.MakeFailedError(errors.New("type not convertible"))
			}
		}

		if err == nil && !reflect.DeepEqual(value, newValRV.Interface()) {
			p.rValue.Set(newValRV)
			changed = true
		}
	}

	if p.valueMu != nil {
		p.valueMu.Unlock()
	}
	return
}

func (p *fieldProp) getWriteCallback() PropertyWriteCallback {
	p.cbMu.Lock()
	cb := p.writeCb
	p.cbMu.Unlock()
	return cb
}

func (p *fieldProp) getReadCallback() PropertyReadCallback {
	p.cbMu.Lock()
	cb := p.readCb
	p.cbMu.Unlock()
	return cb
}

func (p *fieldProp) setWriteCallback(cb PropertyWriteCallback) {
	p.cbMu.Lock()
	p.writeCb = cb
	p.cbMu.Unlock()
}

func (p *fieldProp) setReadCallback(cb PropertyReadCallback) {
	p.cbMu.Lock()
	p.readCb = cb
	p.cbMu.Unlock()
}

func (p *fieldProp) connectChanged(cb PropertyChangedCallback) {
	p.cbMu.Lock()

	// copy on write
	newCbs := make([]PropertyChangedCallback, len(p.changedCbs)+1)
	copy(newCbs, p.changedCbs)
	newCbs[len(newCbs)-1] = cb
	p.changedCbs = newCbs

	p.cbMu.Unlock()
}

// do changed callbacks
func (p *fieldProp) notifyChanged(change *PropertyChanged) {
	p.cbMu.Lock()
	callbacks := p.changedCbs
	p.cbMu.Unlock()
	for _, cb := range callbacks {
		cb(change)
	}
}

// emit DBus signal Properties.PropertiesChanged
func emitPropertiesChanged(conn *dbus.Conn, path dbus.ObjectPath, interfaceName string,
	propName string, value interface{}, emit emitType) (err error) {
	const signal = orgFreedesktopDBus + ".Properties.PropertiesChanged"
	var changedProps map[string]dbus.Variant
	switch emit {
	case emitFalse:
		// do nothing
	case emitInvalidates:
		err = conn.Emit(path, signal, interfaceName, changedProps, []string{propName})
	case emitTrue:
		changedProps = map[string]dbus.Variant{
			propName: dbus.MakeVariant(value),
		}
		err = conn.Emit(path, signal, interfaceName, changedProps, []string{})
	default:
		panic("invalid value for emitType")
	}
	return
}

func getPropsIntrospection(props map[string]*fieldPropStatic) []introspect.Property {
	var result = make([]introspect.Property, len(props))
	idx := 0
	for _, p := range props {

		var access string
		switch p.access {
		case accessWrite:
			access = "write"
		case accessRead:
			access = "read"
		case accessReadWrite:
			access = "readwrite"
		default:
			panic("invalid access")
		}

		result[idx] = introspect.Property{
			Name:   p.name,
			Type:   p.signature.String(),
			Access: access,
		}
		idx++
	}

	return result
}

func getSignals(structType reflect.Type) []introspect.Signal {
	signalsField, ok := structType.FieldByName("signals")
	if !ok {
		return nil
	}

	if signalsField.Type.Kind() != reflect.Ptr {
		return nil
	}

	signalsFieldElemType := signalsField.Type.Elem()
	if signalsFieldElemType.Kind() != reflect.Struct {
		return nil
	}

	var signals []introspect.Signal
	numField := signalsFieldElemType.NumField()
	for i := 0; i < numField; i++ {
		signalItem := signalsFieldElemType.Field(i)
		signalItemType := signalItem.Type

		if signalItemType.Kind() == reflect.Struct {
			var args []introspect.Arg
			numArg := signalItemType.NumField()
			for j := 0; j < numArg; j++ {
				signalArg := signalItemType.Field(j)
				args = append(args, introspect.Arg{
					Name: signalArg.Name,
					Type: dbus.SignatureOfType(signalArg.Type).String(),
				})
			}
			signals = append(signals, introspect.Signal{
				Name: signalItem.Name,
				Args: args,
			})
		}
	}
	return signals
}

const propsMuField = "PropsMu"

func getCorePropsMu(structValue reflect.Value) *sync.RWMutex {
	propsMasterRV := structValue.FieldByName(propsMuField)
	if !propsMasterRV.IsValid() {
		return nil
	}
	return propsMasterRV.Addr().Interface().(*sync.RWMutex)
}

func getStructValue(m interface{}) (reflect.Value, bool) {
	type0 := reflect.TypeOf(m)
	value0 := reflect.ValueOf(m)

	if type0.Kind() != reflect.Ptr {
		return reflect.Value{}, false
	}

	elemType := type0.Elem()
	elemValue := value0.Elem()

	if elemType.Kind() != reflect.Struct {
		return reflect.Value{}, false
	}
	if !elemValue.IsValid() {
		return reflect.Value{}, false
	}
	return elemValue, true
}

func getFieldPropStaticMap(structType reflect.Type,
	structValue reflect.Value) map[string]*fieldPropStatic {

	props := make(map[string]*fieldPropStatic)

	var prevField reflect.StructField
	numField := structType.NumField()
	for i := 0; i < numField; i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)

		if field.Name == propsMuField {
			prevField = field
			continue
		}

		if !fieldValue.CanSet() {
			prevField = field
			continue
		}

		tag := field.Tag.Get("prop")
		if tag == "-" {
			prevField = field
			continue
		}

		if prevField.Name+"Mu" == field.Name {
			prevField = field
			continue
		}

		prop0 := newFieldPropStatic(field, fieldValue, tag)
		props[field.Name] = prop0
		prevField = field
	}
	return props
}

func getFieldPropMap(impl *implementer, implStatic *implementerStatic,
	structValue reflect.Value, s *Service, path dbus.ObjectPath) map[string]*fieldProp {

	structType := structValue.Type()
	props := make(map[string]*fieldProp)

	corePropsMu := getCorePropsMu(structValue)

	numField := structType.NumField()
	var prevField reflect.StructField
	for i := 0; i < numField; i++ {
		field := structType.Field(i)
		fieldValue := structValue.Field(i)

		// ex:
		// prevField: Prop1
		// current Field: Prop1Mu
		if prevField.Name+"Mu" == field.Name &&
			props[prevField.Name] != nil {

			mu, ok := fieldValue.Addr().Interface().(*sync.RWMutex)
			if ok {
				// override prev fieldProp.ValueMu
				props[prevField.Name].valueMu = mu
			}

			prevField = field
			continue
		}

		propStatic, ok := implStatic.props[field.Name]
		if !ok {
			prevField = field
			continue
		}

		p := &fieldProp{
			rValue: fieldValue,
		}

		var propValue Property
		switch propStatic.valueType {
		case fieldPropValueNotProp:
			p.valueMu = corePropsMu

		case fieldPropValueImplProp:
			propValue = fieldValue.Interface().(Property)

		case fieldPropValuePtrImplProp:
			fieldValuePtr := fieldValue.Addr()
			propValue = fieldValuePtr.Interface().(Property)
			p.rValue = fieldValuePtr
		}

		if propValue != nil {
			propValue.SetNotifyChangedFunc(func(val interface{}) {
				impl.notifyChanged(s, path, p, propStatic, val)
			})
		}

		props[field.Name] = p
		prevField = field
	}
	return props
}

func parsePropTag(tag string) (accessType, emitType) {
	access := accessRead
	emit := emitTrue
	tagParts := strings.Split(tag, ",")
	for _, tagPart := range tagParts {
		if strings.HasPrefix(tagPart, "access:") {
			accessStr := tagPart[len("access:"):]
			switch accessStr {
			case "r", "read":
				access = accessRead
			case "w", "write":
				access = accessWrite
			case "rw", "readwrite":
				access = accessReadWrite
			default:
				panic(fmt.Errorf("invalid access %q", accessStr))
			}
			continue
		} else if strings.HasPrefix(tagPart, "emit:") {
			emitStr := tagPart[len("emit:"):]
			switch emitStr {
			case "true":
				emit = emitTrue
			case "false":
				emit = emitFalse
			case "invalidates":
				emit = emitInvalidates
			default:
				panic(fmt.Errorf("invalid emit %q", emitStr))
			}
			continue
		}
	}
	return access, emit
}

func toProperty(value reflect.Value) (Property, fieldPropValueType) {
	propValue, ok := value.Interface().(Property)
	if ok {
		return propValue, fieldPropValueImplProp
	}

	// try value.Addr
	if value.Kind() == reflect.Struct {
		propValue, ok = value.Addr().Interface().(Property)
		if ok {
			return propValue, fieldPropValuePtrImplProp
		}
	}
	return nil, fieldPropValueNotProp
}

func newFieldPropStatic(field reflect.StructField, fieldValue reflect.Value,
	tag string) *fieldPropStatic {

	access, emit := parsePropTag(tag)
	p := &fieldPropStatic{
		name:   field.Name,
		access: access,
		emit:   emit,
	}
	var rType reflect.Type

	propValue, valueType := toProperty(fieldValue)
	p.valueType = valueType
	if valueType == fieldPropValueNotProp {
		rType = field.Type
	} else {
		rType = propValue.GetType()
	}

	p.rType = rType
	p.signature = dbus.SignatureOfType(rType)
	if strings.Contains(p.signature.String(), "(") {
		p.hasStruct = true
	}
	return p
}

type methodDetail struct {
	In  []string
	Out []string
}

func (md methodDetail) getInArgName(index int, type0 reflect.Type, methodName string) string {
	if index >= len(md.In) {
		panic(fmt.Errorf("failed to get %s.%s in[%d] argument name",
			type0, methodName, index))
	}
	return md.In[index]
}

func (md methodDetail) getOutArgName(index int, type0 reflect.Type, methodName string) string {
	if index >= len(md.Out) {
		panic(fmt.Errorf("failed to get %s.%s out[%d] argument name",
			type0, methodName, index))
	}
	return md.Out[index]
}

func getMethodDetailMap(structType reflect.Type) map[string]methodDetail {
	result := make(map[string]methodDetail)
	methodsField, ok := structType.FieldByName("methods")
	if !ok {
		return nil
	}

	if methodsField.Type.Kind() != reflect.Ptr {
		return nil
	}

	methodsFieldElemType := methodsField.Type.Elem()
	if methodsFieldElemType.Kind() != reflect.Struct {
		return nil
	}

	numField := methodsFieldElemType.NumField()
	for i := 0; i < numField; i++ {
		methodItem := methodsFieldElemType.Field(i)
		tagIn := methodItem.Tag.Get("in")
		tagOut := methodItem.Tag.Get("out")

		result[methodItem.Name] = methodDetail{
			In:  splitArg(tagIn),
			Out: splitArg(tagOut),
		}
	}
	return result
}

func splitArg(str string) (result []string) {
	parts := strings.Split(str, ",")
	for _, part := range parts {
		part = strings.TrimSpace(part)
		if part != "" {
			result = append(result, part)
		}
	}
	return
}

// Methods returns the description of the methods of v. This can be used to
// create a Node which can be passed to NewIntrospectable.
func getMethodsOld(v interface{}, methodDetailMap map[string]methodDetail) []introspect.Method {
	t := reflect.TypeOf(v)
	ms := make([]introspect.Method, 0, t.NumMethod())
	for i := 0; i < t.NumMethod(); i++ {
		if t.Method(i).PkgPath != "" {
			continue
		}
		mt := t.Method(i).Type
		if mt.NumOut() == 0 ||
			mt.Out(mt.NumOut()-1) != reflect.TypeOf(&dbus.Error{}) {

			continue
		}
		var m introspect.Method
		m.Name = t.Method(i).Name
		m.Args = make([]introspect.Arg, 0, mt.NumIn()+mt.NumOut()-2)

		methodDetail := methodDetailMap[m.Name]
		inArgIndex := 0
		for j := 1; j < mt.NumIn(); j++ {
			if mt.In(j) != reflect.TypeOf((*dbus.Sender)(nil)).Elem() &&
				mt.In(j) != reflect.TypeOf((*dbus.Message)(nil)).Elem() {

				argName := methodDetail.getInArgName(inArgIndex, t, m.Name)
				inArgIndex++
				arg := introspect.Arg{Name: argName,
					Type:      dbus.SignatureOfType(mt.In(j)).String(),
					Direction: "in",
				}
				m.Args = append(m.Args, arg)
			}
		}
		for j := 0; j < mt.NumOut()-1; j++ {
			argName := methodDetail.getOutArgName(j, t, m.Name)
			arg := introspect.Arg{
				Name:      argName,
				Type:      dbus.SignatureOfType(mt.Out(j)).String(),
				Direction: "out",
			}
			m.Args = append(m.Args, arg)
		}
		m.Annotations = make([]introspect.Annotation, 0)
		ms = append(ms, m)
	}
	return ms
}

func (em *ExportedMethod) getInArgName(index int, type0 reflect.Type, methodName string) string {
	if index >= len(em.InArgs) {
		panic(fmt.Errorf("failed to get %s.%s in[%d] argument name",
			type0, methodName, index))
	}
	return em.InArgs[index]
}

func (em *ExportedMethod) getOutArgName(index int, type0 reflect.Type, methodName string) string {
	if index >= len(em.OutArgs) {
		panic(fmt.Errorf("failed to get %s.%s out[%d] argument name",
			type0, methodName, index))
	}
	return em.OutArgs[index]
}

var (
	typeOfDBusErrorPtr = reflect.TypeOf((*dbus.Error)(nil))
	typeOfDBusSender   = reflect.TypeOf((*dbus.Sender)(nil)).Elem()
	typeOfDBusMessage  = reflect.TypeOf((*dbus.Message)(nil)).Elem()
)

// Methods returns the description of the methods of v. This can be used to
// create a Node which can be passed to NewIntrospectable.
func getMethods(v interface{}, methods ExportedMethods) []introspect.Method {
	t := reflect.TypeOf(v)
	result := make([]introspect.Method, 0, len(methods))
	for _, method := range methods {
		methodType := reflect.TypeOf(method.Fn)

		numIn := methodType.NumIn()
		numOut := methodType.NumOut()
		if numOut == 0 ||
			methodType.Out(numOut-1) != typeOfDBusErrorPtr {

			continue
		}
		var m introspect.Method
		m.Name = method.Name
		m.Args = make([]introspect.Arg, 0, numIn+numOut-1)
		inArgIndex := 0
		for j := 0; j < numIn; j++ {
			argType := methodType.In(j)
			if argType == typeOfDBusSender || argType == typeOfDBusMessage {
				// 忽略类型为 dbus.Sender 或 dbus.Message 的参数
				continue
			}

			argName := method.getInArgName(inArgIndex, t, m.Name)
			inArgIndex++
			arg := introspect.Arg{
				Name:      argName,
				Type:      dbus.SignatureOfType(methodType.In(j)).String(),
				Direction: "in",
			}
			m.Args = append(m.Args, arg)
		}
		for j := 0; j < numOut-1; j++ {
			argName := method.getOutArgName(j, t, m.Name)
			arg := introspect.Arg{
				Name:      argName,
				Type:      dbus.SignatureOfType(methodType.Out(j)).String(),
				Direction: "out",
			}
			m.Args = append(m.Args, arg)
		}
		m.Annotations = make([]introspect.Annotation, 0)
		result = append(result, m)
	}
	return result
}

type PropertyReadCallback func(read *PropertyRead) *dbus.Error

type PropertyWriteCallback func(write *PropertyWrite) *dbus.Error

type PropertyChangedCallback func(change *PropertyChanged)

type Property interface {
	SetValue(val interface{}) (changed bool, err *dbus.Error)
	GetValue() (val interface{}, err *dbus.Error)
	SetNotifyChangedFunc(func(val interface{}))
	GetType() reflect.Type
}

type PropertyInfo struct {
	Path      dbus.ObjectPath
	Interface string
	Name      string
}

type PropertyAccess struct {
	PropertyInfo
	Sender  dbus.Sender
	service *Service
}

func (pa *PropertyAccess) GetPID() (uint32, error) {
	return pa.service.GetConnPID(string(pa.Sender))
}

func (pa *PropertyAccess) GetUID() (uint32, error) {
	return pa.service.GetConnUID(string(pa.Sender))
}

type PropertyRead struct {
	PropertyAccess
}

func newPropertyRead(sender dbus.Sender, so *ServerObject,
	interfaceName, name string) *PropertyRead {

	pr := new(PropertyRead)
	pr.Sender = sender
	pr.service = so.service
	pr.Name = name
	pr.Interface = interfaceName
	pr.Path = so.path
	return pr
}

type PropertyWrite struct {
	PropertyAccess
	Value interface{} // new value
}

func newPropertyWrite(sender dbus.Sender, so *ServerObject,
	interfaceName, name string, value interface{}) *PropertyWrite {

	pw := new(PropertyWrite)
	pw.Sender = sender
	pw.service = so.service
	pw.Name = name
	pw.Interface = interfaceName
	pw.Path = so.path
	pw.Value = value
	return pw
}

type PropertyChanged struct {
	PropertyInfo
	Value interface{} // new value
}

func newPropertyChanged(path dbus.ObjectPath, interfaceName, name string,
	value interface{}) *PropertyChanged {
	pc := new(PropertyChanged)
	pc.Name = name
	pc.Interface = interfaceName
	pc.Path = path
	pc.Value = value
	return pc
}

func valueFromBus(src interface{}, valueRT reflect.Type) (reflect.Value, error) {
	newValueRV := reflect.New(valueRT)
	err := dbus.Store([]interface{}{src}, newValueRV.Interface())
	if err != nil {
		return reflect.Value{}, err
	}
	return newValueRV.Elem(), nil
}
