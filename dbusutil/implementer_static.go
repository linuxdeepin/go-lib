package dbusutil

import (
	"errors"
	"reflect"

	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
)

func newImplementerStatic(impl Implementer, structValue reflect.Value) *implementerStatic {
	s := &implementerStatic{}
	structType := structValue.Type()
	s.props = getFieldPropStaticMap(structType, structValue)

	s.introspectInterface = introspect.Interface{
		Name:       impl.GetInterfaceName(),
		Signals:    getSignals(structType),
		Properties: getPropsIntrospection(s.props),
		Methods:    getMethods(impl, getMethodDetailMap(structType)),
	}
	return s
}

type implementerStatic struct {
	props               map[string]*fieldPropStatic
	introspectInterface introspect.Interface
}

func (is *implementerStatic) checkProperty(propName string) error {
	_, ok := is.props[propName]
	if ok {
		return nil
	}
	return errors.New("property not found")
}

func (is *implementerStatic) checkPropertyValue(propName string, value interface{}) error {
	p := is.props[propName]
	if p == nil {
		return errors.New("property not found")
	}

	if p.signature != dbus.SignatureOf(value) {
		return errors.New("property signature not equal")
	}

	return nil
}
