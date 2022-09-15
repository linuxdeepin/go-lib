// SPDX-FileCopyrightText: 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package dbusutil

import (
	"errors"

	"github.com/godbus/dbus"
)

type ServerObject struct {
	service      *Service
	path         dbus.ObjectPath
	implementers []*implementer
}

func (so *ServerObject) Path() dbus.ObjectPath {
	return so.path
}

func (so *ServerObject) Export() error {
	s := so.service
	conn := s.conn
	path := so.path

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.objMap[path]
	if ok {
		return errors.New("server object is exported")
	}

	for _, impl := range so.implementers {
		core := impl.core
		corePtr := getImplementerPointer(core)
		_, ok := s.implObjMap[corePtr]
		if ok {
			return errors.New("implementer is exported")
		}

		// 对旧代码实现兼容
		if coreExt, ok := core.(ImplementerExt); ok {
			err := conn.ExportMethodTable(coreExt.GetExportedMethods().toMethodTable(), so.path, core.GetInterfaceName())
			if err != nil {
				return err
			}
		} else {
			err := conn.Export(core, so.path, core.GetInterfaceName())
			if err != nil {
				return err
			}
		}

		s.implObjMap[corePtr] = so
	}

	methodTable := make(map[string]interface{}, 3)
	methodTable["Introspect"] = so.introspectableIntrospect

	err := conn.ExportMethodTable(methodTable, so.path,
		orgFreedesktopDBus+".Introspectable")
	if err != nil {
		return err
	}

	delete(methodTable, "Introspect")
	methodTable["Get"] = so.propertiesGet
	methodTable["GetAll"] = so.propertiesGetAll
	methodTable["Set"] = so.propertiesSet

	err = conn.ExportMethodTable(methodTable, so.path,
		orgFreedesktopDBus+".Properties")
	if err != nil {
		return err
	}

	s.objMap[so.path] = so
	return nil
}

func (so *ServerObject) StopExport() error {
	s := so.service
	conn := s.conn
	path := so.path

	s.mu.Lock()
	defer s.mu.Unlock()

	_, ok := s.objMap[path]
	if !ok {
		return errors.New("server object is not exported")
	}

	// TODO: 等 github.com/godbus/dbus 升级之后，需要把所有 conn.Export 都换成 conn.ExportMethodTable。
	// 目前是由于 conn.ExportMethodTable 方法在 methods 参数为 nil 时有 bug，才用了 conn.Export 方法。
	err := conn.Export(nil, path, orgFreedesktopDBus+".Properties")
	if err != nil {
		return err
	}

	err = conn.Export(nil, path, orgFreedesktopDBus+".Introspectable")
	if err != nil {
		return err
	}

	for _, impl := range so.implementers {
		core := impl.core
		err := conn.Export(nil, so.path, core.GetInterfaceName())
		if err != nil {
			return err
		}
		corePtr := getImplementerPointer(core)
		delete(s.implObjMap, corePtr)
	}

	delete(s.objMap, path)
	return nil
}

func (so *ServerObject) getImplementer(interfaceName string) *implementer {
	for _, impl := range so.implementers {
		if impl.core.GetInterfaceName() == interfaceName {
			return impl
		}
	}
	return nil
}

func (so *ServerObject) SetWriteCallback(v Implementer, propertyName string,
	cb PropertyWriteCallback) error {
	impl := so.getImplementer(v.GetInterfaceName())
	if impl == nil {
		return errors.New("not exported")
	}

	return impl.setWriteCallback(propertyName, cb)
}

func (so *ServerObject) SetReadCallback(v Implementer, propertyName string,
	cb PropertyReadCallback) error {
	impl := so.getImplementer(v.GetInterfaceName())
	if impl == nil {
		return errors.New("not exported")
	}

	return impl.setReadCallback(propertyName, cb)
}

func (so *ServerObject) ConnectChanged(v Implementer, propertyName string,
	cb PropertyChangedCallback) error {

	impl := so.getImplementer(v.GetInterfaceName())
	if impl == nil {
		return errors.New("not exported")
	}

	return impl.connectChanged(propertyName, cb)
}
