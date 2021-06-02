package dbusutil

import (
	"bytes"
	"errors"
	"fmt"
	"sort"
	"sync"
	"time"
	"unsafe"

	"github.com/godbus/dbus"
	"github.com/godbus/dbus/introspect"
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
	implObjMap map[unsafe.Pointer]*ServerObject
}

func NewService(conn *dbus.Conn) *Service {
	return &Service{
		conn:          conn,
		objMap:        make(map[dbus.ObjectPath]*ServerObject),
		implStaticMap: make(map[string]*implementerStatic),
		implObjMap:    make(map[unsafe.Pointer]*ServerObject),
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

func (s *Service) NewServerObject(path dbus.ObjectPath,
	implementers ...Implementer) (*ServerObject, error) {

	if !path.IsValid() {
		return nil, errors.New("path invalid")
	}

	if len(implementers) == 0 {
		return nil, errors.New("no implementer")
	}

	implMap := make(map[string]Implementer)
	implSlice := make([]*implementer, 0, len(implementers))

	for _, implCore := range implementers {
		ifcName := implCore.GetInterfaceName()
		_, ok := implMap[ifcName]
		if ok {
			return nil, errors.New("interface duplicated")
		}

		impl, err := newImplementer(implCore, s, path)
		if err != nil {
			return nil, err
		}
		implSlice = append(implSlice, impl)
		implMap[ifcName] = implCore
	}

	obj := &ServerObject{
		service:      s,
		path:         path,
		implementers: implSlice,
	}

	return obj, nil
}

func (s *Service) GetServerObject(impl Implementer) *ServerObject {
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

func (s *Service) Export(path dbus.ObjectPath, implements ...Implementer) error {
	so, err := s.NewServerObject(path, implements...)
	if err != nil {
		return err
	}
	return so.Export()
}

func (s *Service) StopExport(impl Implementer) error {
	so := s.GetServerObject(impl)
	if so == nil {
		return errors.New("server object not found")
	}
	return so.StopExport()
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

func (s *Service) getImplementerStatic(impl Implementer) *implementerStatic {
	ifcName := impl.GetInterfaceName()
	s.mu.RLock()
	implStatic := s.implStaticMap[ifcName]
	s.mu.RUnlock()
	return implStatic
}

func (s *Service) Emit(v Implementer, signalName string, values ...interface{}) error {
	so := s.GetServerObject(v)
	if so == nil {
		return errors.New("v is not exported")
	}
	implStatic := s.getImplementerStatic(v)

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

	return s.conn.Emit(so.path,
		v.GetInterfaceName()+"."+signalName, values...)
}

func (s *Service) EmitPropertyChanged(v Implementer, propertyName string, value interface{}) error {
	so := s.GetServerObject(v)
	if so == nil {
		return errors.New("v is not exported")
	}

	impl := so.getImplementer(v.GetInterfaceName())

	implStatic := s.getImplementerStatic(v)

	err := implStatic.checkPropertyValue(propertyName, value)
	if err != nil {
		return err
	}

	return impl.emitPropChanged(s, so.path, propertyName, value)
}

func (s *Service) DelayEmitPropertyChanged(v Implementer) func() error {
	so := s.GetServerObject(v)
	if so == nil {
		return nil
	}

	impl := so.getImplementer(v.GetInterfaceName())
	if impl == nil {
		return nil
	}

	impl.delayEmitPropChanged()
	return func() error {
		return impl.stopDelayEmitPropChanged(s, so.path)
	}
}

func (s *Service) EmitPropertiesChanged(v Implementer, propValMap map[string]interface{},
	invalidatedProps ...string) error {
	so := s.GetServerObject(v)
	if so == nil {
		return errors.New("v is not exported")
	}

	implStatic := s.getImplementerStatic(v)

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
	return s.conn.Emit(so.path, signalName, v.GetInterfaceName(), changedProps, invalidatedProps)
}

func (s *Service) Quit() {
	s.conn.Close()
	close(s.quit)
}

func (s *Service) Wait() {
	s.quit = make(chan struct{})
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
	so := s.GetServerObject(v)
	if so == nil {
		return "", errors.New("not exported")
	}

	impl := so.getImplementer(v.GetInterfaceName())
	implStatic := impl.getStatic(s)

	var buf bytes.Buffer

	var propNames []string
	for name := range impl.props {
		propNames = append(propNames, name)
	}
	sort.Strings(propNames)

	for _, name := range propNames {
		p := impl.props[name]
		fmt.Fprintln(&buf, "property name:", name)
		fmt.Fprintf(&buf, "valueMu: %p\n", p.valueMu)

		propStatic := implStatic.props[name]
		fmt.Fprintln(&buf, "signature:", propStatic.signature)
		fmt.Fprintln(&buf, "access:", propStatic.access)
		fmt.Fprintln(&buf, "emit:", propStatic.emit)

		buf.WriteString("\n")
	}

	return buf.String(), nil
}
