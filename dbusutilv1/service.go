// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
// 
// SPDX-License-Identifier: GPL-3.0-or-later
package dbusutilv1

import (
	"errors"
	"fmt"
	"sync"
	"time"
	"unsafe"

	"github.com/godbus/dbus/v5"
	"github.com/godbus/dbus/v5/introspect"
)

type Service struct {
	conn *dbus.Conn
	mu   sync.RWMutex

	hasCall bool

	quit              chan struct{}
	canQuit           func() bool
	quitCheckInterval time.Duration

	objMap        map[dbus.ObjectPath]*ServerObject
	implStaticMap map[string]*implementerStatic
	//                ^interface name
	implObjMap map[unsafe.Pointer][]*ServerObject
}

func NewService(conn *dbus.Conn) *Service {
	return &Service{
		conn:          conn,
		objMap:        make(map[dbus.ObjectPath]*ServerObject),
		implStaticMap: make(map[string]*implementerStatic),
		implObjMap:    make(map[unsafe.Pointer][]*ServerObject),
	}
}

func NewSessionService() (*Service, error) {
	bus, err := dbus.SessionBus()
	if err != nil {
		return nil, err
	}
	return NewService(bus), nil
}

func NewSystemService() (*Service, error) {
	bus, err := dbus.SystemBus()
	if err != nil {
		return nil, err
	}
	return NewService(bus), nil
}

func (s *Service) Conn() *dbus.Conn {
	return s.conn
}

func (s *Service) RequestName(name string) error {
	reply, err := s.conn.RequestName(name, dbus.NameFlagDoNotQueue)
	if err != nil {
		return err
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		return fmt.Errorf("name %s already taken", name)
	}
	return nil
}

func (s *Service) ReleaseName(name string) error {
	_, err := s.conn.ReleaseName(name)
	return err
}

func (s *Service) NewServerObject(path dbus.ObjectPath) (*ServerObject, error) {
	obj := &ServerObject{
		service: s,
		path:    path,
	}
	return obj, nil
}

func (s *Service) ServerObjectAddImpl(so *ServerObject, interfaceName string, implCore Implementer) error {
	isFind := false
	for _, impl := range so.implementers {
		if impl.getInterfaceName() == interfaceName {
			isFind = true
		}
	}
	if isFind {
		return errors.New("interface duplicated")
	}
	impl, err := newImplementer(implCore, interfaceName, s, so.path)
	if err != nil {
		return err
	}
	so.implementers = append(so.implementers, impl)

	return nil
}

func (s *Service) GetServerObject(impl Implementer) []*ServerObject {
	ptr := getImplementerPointer(impl)
	s.mu.RLock()
	so := s.implObjMap[ptr]
	s.mu.RUnlock()
	return so
}

func (s *Service) GetServerObjectByPath(path dbus.ObjectPath) *ServerObject {
	s.mu.RLock()
	so := s.objMap[path]
	s.mu.RUnlock()
	return so
}

func (s *Service) Export(path dbus.ObjectPath, interfaceName string, impl Implementer) error {
	if !path.IsValid() {
		return errors.New("path invalid")
	}
	so := s.GetServerObjectByPath(path)
	if so == nil {
		var err error
		so, err = s.NewServerObject(path)
		if err != nil {
			return err
		}
	}
	err := s.ServerObjectAddImpl(so, interfaceName, impl)
	if err != nil {
		return err
	}
	return so.Export()
}

func (s *Service) StopExport(impl Implementer) error {
	sos := s.GetServerObject(impl)
	if sos == nil {
		return errors.New("server object not found")
	}
	for _, so := range sos {
		err := so.StopExport()
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Service) StopExportByPath(path dbus.ObjectPath) error {
	so := s.GetServerObjectByPath(path)
	if so == nil {
		return errors.New("server object not found")
	}
	return so.StopExport()
}

func (s *Service) IsExported(impl Implementer) bool {
	so := s.GetServerObject(impl)
	return so != nil
}

func (s *Service) getImplementerStatic(ifcName string) *implementerStatic {
	s.mu.RLock()
	implStatic := s.implStaticMap[ifcName]
	s.mu.RUnlock()
	return implStatic
}

func (s *Service) Emit(v Implementer, signalName string, values ...interface{}) error {
	sos := s.GetServerObject(v)
	if sos == nil {
		return errors.New("v is not exported")
	}

	for _, so := range sos {
		impl := so.getImplementer(v)
		if impl == nil {
			return errors.New("impl not exist")
		}

		implStatic := s.getImplementerStatic(impl.getInterfaceName())

		var signal introspect.Signal
		for _, sig := range implStatic.introspectInterface.Signals {
			if sig.Name == signalName {
				signal = sig
				break
			}
		}

		if signal.Name == "" {
			return errors.New("not found signal")
		}
		if len(values) != len(signal.Args) {
			return errors.New("signal args length not equal")
		}
		for idx, arg := range signal.Args {
			valueType := dbus.SignatureOf(values[idx]).String()
			if valueType != arg.Type {
				return fmt.Errorf("signal arg[%d] type not match", idx)
			}
		}

		err := s.conn.Emit(so.path, impl.getInterfaceName()+"."+signalName, values...)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) EmitPropertyChanged(v Implementer, propertyName string, value interface{}) error {
	sos := s.GetServerObject(v)
	if sos == nil {
		return errors.New("v is not exported")
	}

	for _, so := range sos {
		impl := so.getImplementer(v)
		if impl == nil {
			return errors.New("impl not exist")
		}

		implStatic := s.getImplementerStatic(impl.getInterfaceName())

		err := implStatic.checkPropertyValue(propertyName, value)
		if err != nil {
			return err
		}
		err = impl.emitPropChanged(s, so.path, propertyName, value)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) DelayEmitPropertyChanged(v Implementer) func() error {
	sos := s.GetServerObject(v)
	if sos == nil {
		return nil
	}

	for _, so := range sos {
		impl := so.getImplementer(v)
		if impl == nil {
			return nil
		}

		impl.delayEmitPropChanged()
	}

	cb := func() error {
		for _, so := range sos {
			impl := so.getImplementer(v)
			if impl == nil {
				return errors.New("impl not exist")
			}
			err := impl.stopDelayEmitPropChanged(s, so.path)
			if err != nil {
				return err
			}
		}
		return nil
	}

	return cb
}

func (s *Service) EmitPropertiesChanged(v Implementer, propValMap map[string]interface{},
	invalidatedProps ...string) error {
	sos := s.GetServerObject(v)
	if sos == nil {
		return errors.New("v is not exported")
	}

	for _, so := range sos {
		impl := so.getImplementer(v)
		if impl == nil {
			return errors.New("impl not exist")
		}

		implStatic := s.getImplementerStatic(impl.getInterfaceName())

		const signalName = orgFreedesktopDBus + ".Properties.PropertiesChanged"
		var changedProps map[string]dbus.Variant
		if len(propValMap) > 0 {
			changedProps = make(map[string]dbus.Variant)
		}
		for propName, val := range propValMap {
			err := implStatic.checkPropertyValue(propName, val)
			if err != nil {
				return err
			}
			changedProps[propName] = dbus.MakeVariant(val)
		}
		for _, propName := range invalidatedProps {
			if _, ok := propValMap[propName]; ok {
				return errors.New("property appears in both propValMap and invalidateProps")
			}

			err := implStatic.checkProperty(propName)
			if err != nil {
				return err
			}
		}

		err := s.conn.Emit(so.path, signalName, impl.getInterfaceName(), changedProps, invalidatedProps)
		if err != nil {
			return err
		}
	}

	return nil
}

func (s *Service) Quit() {
	s.mu.Lock()
	if s.quit != nil {
		close(s.quit)
		s.quit = nil
	}
	s.mu.Unlock()
}

func (s *Service) Wait() {
	s.mu.Lock()
	s.quit = make(chan struct{})
	s.mu.Unlock()
	if s.quitCheckInterval > 0 {
		go func() {
			ticker := time.NewTicker(s.quitCheckInterval)
			for {
				select {
				case <-s.quit:
					return
				case <-ticker.C:
					s.mu.RLock()
					hasCall := s.hasCall
					s.mu.RUnlock()
					logger.Println("Service.Wait tick hasCall:", hasCall)

					if !hasCall {
						if s.canQuit == nil || s.canQuit() {
							s.Quit()
							return
						}
					} else {
						s.mu.Lock()
						s.hasCall = false
						s.mu.Unlock()
					}
				}
			}
		}()
	}
	<-s.quit
}

func (s *Service) DelayAutoQuit() {
	s.mu.Lock()
	s.hasCall = true
	s.mu.Unlock()
}

func (s *Service) SetAutoQuitHandler(interval time.Duration, canQuit func() bool) {
	s.quitCheckInterval = interval
	s.canQuit = canQuit
}

func (s *Service) GetConnPID(name string) (pid uint32, err error) {
	err = s.conn.BusObject().Call(orgFreedesktopDBus+".GetConnectionUnixProcessID",
		0, name).Store(&pid)
	return
}

func (s *Service) GetConnUID(name string) (uid uint32, err error) {
	err = s.conn.BusObject().Call(orgFreedesktopDBus+".GetConnectionUnixUser",
		0, name).Store(&uid)
	return
}

func (s *Service) GetNameOwner(name string) (owner string, err error) {
	err = s.conn.BusObject().Call(orgFreedesktopDBus+".GetNameOwner",
		0, name).Store(&owner)
	return
}

func (s *Service) NameHasOwner(name string) (hasOwner bool, err error) {
	err = s.conn.BusObject().Call(orgFreedesktopDBus+".NameHasOwner",
		0, name).Store(&hasOwner)
	return
}

func (s *Service) DumpProperties(v Implementer) (string, error) {
	return "", nil
}
