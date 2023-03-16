// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusutil

import (
	"errors"
	"reflect"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

func newImplementerStatic(impl Implementer, interfaceName string, structValue reflect.Value) *implementerStatic {
	s := &implementerStatic{}
	structType := structValue.Type()
	s.props = getFieldPropStaticMap(structType, structValue)

	// 对旧代码实现兼容
	var methods []introspect.Method
	if implExt, ok := impl.(ImplementerExt); ok {
		methods = getMethods(implExt, implExt.GetExportedMethods())
	} else {
		methods = getMethodsOld(impl, getMethodDetailMap(structType))
	}

	s.introspectInterface = introspect.Interface{
		Name:       interfaceName,
		Signals:    getSignals(structType),
		Properties: getPropsIntrospection(s.props),
		Methods:    methods,
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
