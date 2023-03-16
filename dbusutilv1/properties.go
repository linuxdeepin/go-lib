// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
// 
// SPDX-License-Identifier: GPL-3.0-or-later
package dbusutilv1

import (
	"errors"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/prop"
)

func (so *ServerObject) propertiesGet(sender dbus.Sender,
	interfaceName, propertyName string) (dbus.Variant, *dbus.Error) {
	so.service.DelayAutoQuit()

	impl := so.getImplementerByInterface(interfaceName)
	if impl == nil {
		return dbus.Variant{}, prop.ErrIfaceNotFound
	}
	p := impl.props[propertyName]
	if p == nil {
		return dbus.Variant{}, prop.ErrPropNotFound
	}

	implStatic := impl.getStatic(so.service)
	propStatic := implStatic.props[propertyName]

	if propStatic.access&accessRead == 0 {
		return dbus.Variant{}, dbus.MakeFailedError(errors.New("property can not be read"))
	}

	propRead := newPropertyRead(sender, so, interfaceName, propertyName)

	variant, err := p.GetValueVariant(propRead, propStatic.signature)
	if err != nil {
		return dbus.Variant{}, err
	}
	return variant, nil
}

func (so *ServerObject) propertiesGetAll(sender dbus.Sender, interfaceName string) (map[string]dbus.Variant, *dbus.Error) {
	so.service.DelayAutoQuit()

	impl := so.getImplementerByInterface(interfaceName)
	if impl == nil {
		return nil, prop.ErrIfaceNotFound
	}

	implStatic := impl.getStatic(so.service)

	result := make(map[string]dbus.Variant, len(impl.props))
	for propName, p := range impl.props {
		propStatic := implStatic.props[propName]

		if propStatic.access&accessRead != 0 {
			propRead := newPropertyRead(sender, so, interfaceName, propName)
			variant, err := p.GetValueVariant(propRead, propStatic.signature)
			if err != nil {
				// ignore err
				continue
			}
			result[propName] = variant
		}
	}
	return result, nil
}

func (so *ServerObject) propertiesSet(sender dbus.Sender, interfaceName, propertyName string,
	newVar dbus.Variant) *dbus.Error {
	so.service.DelayAutoQuit()

	impl := so.getImplementerByInterface(interfaceName)
	if impl == nil {
		return prop.ErrIfaceNotFound
	}

	p := impl.props[propertyName]
	if p == nil {
		return prop.ErrPropNotFound
	}

	implStatic := impl.getStatic(so.service)
	propStatic := implStatic.props[propertyName]

	if propStatic.access&accessWrite == 0 {
		return dbus.MakeFailedError(errors.New("property can not be written"))
	}

	if newVar.Signature() != propStatic.signature {
		return prop.ErrInvalidArg
	}
	newVarValue := newVar.Value()

	//fix newVarValue []interface{} to struct
	if propStatic.hasStruct {
		fixedRV, err := valueFromBus(newVarValue, propStatic.rType)
		if err != nil {
			return dbus.MakeFailedError(err)
		}
		newVarValue = fixedRV.Interface()
	}

	propWrite := newPropertyWrite(sender, so, interfaceName, propertyName, newVarValue)

	changed, setErr := p.SetValue(propWrite)
	if setErr != nil {
		return setErr
	}
	if changed {
		impl.notifyChanged(so.service, so.path, p, propStatic, propWrite.Value)
	}
	return nil
}
