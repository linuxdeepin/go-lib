package dbusutil

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"path"
	"strings"
	"sync"
	"time"

	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbus1/introspect"
)

type Service struct {
	conn      *dbus.Conn
	objects   map[dbus.ObjectPath]*object
	objectsMu sync.RWMutex

	hasCall   bool
	hasCallMu sync.Mutex

	quit              chan struct{}
	canQuit           func() bool
	quitCheckInterval time.Duration
}

func NewService(conn *dbus.Conn) *Service {
	return &Service{
		conn:    conn,
		objects: make(map[dbus.ObjectPath]*object),
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
		panic(err)
	}
	if reply != dbus.RequestNameReplyPrimaryOwner {
		return errors.New("name already taken")
	}
	return nil
}

func (s *Service) ReleaseName(name string) error {
	_, err := s.conn.ReleaseName(name)
	return err
}

func (s *Service) Export(v Exportable) error {
	exportInfo := v.GetDBusExportInfo()
	log.Println("service.Export", exportInfo)
	path := dbus.ObjectPath(exportInfo.Path)
	if !path.IsValid() {
		return errors.New("path is invalid")
	}

	structType, structValue := getStructTypeValue(v)
	if structType == nil {
		return errors.New("v is not a struct pointer")
	}

	implementer := &implementer{
		service:       s,
		path:          path,
		interfaceName: exportInfo.Interface,
		core:          v,
	}

	implementer.corePropsMu = getPropsMutex(structValue)
	implementer.props = getProps(implementer, structType, structValue)
	implementer.methods = getMethods(v, getMethodDetailMap(structType))
	implementer.signals = getSignals(structType)

	s.objectsMu.RLock()
	obj, ok := s.objects[path]
	s.objectsMu.RUnlock()
	if !ok {
		// path not exist
		obj = newObject(path, s)

		err := s.conn.Export(v, path, exportInfo.Interface)
		if err != nil {
			return err
		}

		err = obj.exportProperties(s.conn)
		if err != nil {
			return err
		}

		err = obj.exportIntrospectable(s.conn)
		if err != nil {
			return err
		}

		s.objectsMu.Lock()
		s.objects[path] = obj
		s.addPath(path)
		s.objectsMu.Unlock()
	} else {

		err := s.conn.Export(v, path, exportInfo.Interface)
		if err != nil {
			return err
		}
	}
	obj.addImplementer(implementer)

	return nil
}

func (s *Service) StopExport(exportInfo ExportInfo) error {
	path := dbus.ObjectPath(exportInfo.Path)

	s.objectsMu.RLock()
	obj, ok := s.objects[path]
	s.objectsMu.RUnlock()
	if ok {
		if obj.hasImplementer(exportInfo.Interface) {
			err := s.conn.Export(nil, path, exportInfo.Interface)
			if err != nil {
				return err
			}
			obj.deleteImplementer(exportInfo.Interface)
		}
		if obj.numImplementer() == 0 {
			err := obj.stopExportIntrospectable(s.conn)
			if err != nil {
				return err
			}
			err = obj.stopExportProperties(s.conn)
			if err != nil {
				return err
			}
			s.objectsMu.Lock()
			delete(s.objects, path)
			s.removePath(path)
			s.objectsMu.Unlock()
		}
	}
	return nil
}

func (s *Service) IsExported(exportInfo ExportInfo) bool {
	path := dbus.ObjectPath(exportInfo.Path)

	s.objectsMu.RLock()
	obj, ok := s.objects[path]
	s.objectsMu.RUnlock()
	if ok {
		if obj.hasImplementer(exportInfo.Interface) {
			return true
		}
	}
	return false
}

func (s *Service) addPath(p dbus.ObjectPath) {
	for p != "/" {
		parentPath := dbus.ObjectPath(path.Dir(string(p)))
		parentObj, ok := s.objects[parentPath]
		if ok {
			child := path.Base(string(p))
			log.Println("parent", parentPath, "add child", child)
			intro := parentObj.introspectableImpl
			intro.addChild(child)
			break
		}
		p = parentPath
	}
}

func (s *Service) removePath(p dbus.ObjectPath) {
	for p != "/" {
		parentPath := dbus.ObjectPath(path.Dir(string(p)))
		parentObj, ok := s.objects[parentPath]
		if ok {
			if !s.pathInUse(p) {
				child := path.Base(string(p))
				intro := parentObj.introspectableImpl
				log.Println("parent", parentPath, "delete child", child)
				intro.deleteChild(child)
			}

			break
		}
		p = parentPath
	}
}

func (s *Service) pathInUse(p dbus.ObjectPath) bool {
	for p0 := range s.objects {
		if strings.HasPrefix(string(p0), string(p)+"/") {
			return true
		}
	}
	return false
}

func (s *Service) getImplementer(exportInfo ExportInfo) *implementer {
	s.objectsMu.RLock()
	obj := s.objects[dbus.ObjectPath(exportInfo.Path)]
	s.objectsMu.RUnlock()
	if obj == nil {
		return nil
	}
	return obj.getImplementer(exportInfo.Interface)
}

func (s *Service) Emit(v Exportable, signalName string, values ...interface{}) error {
	exportInfo := v.GetDBusExportInfo()
	impl := s.getImplementer(exportInfo)
	var signal introspect.Signal
	for _, sig := range impl.signals {
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

	return s.conn.Emit(dbus.ObjectPath(exportInfo.Path),
		exportInfo.Interface+"."+signalName, values...)
}

func (s *Service) EmitPropertyChanged(v Exportable, propertyName string, value interface{}) error {
	exportInfo := v.GetDBusExportInfo()
	impl := s.getImplementer(exportInfo)
	err := impl.checkPropertyValue(propertyName, value)
	if err != nil {
		return err
	}

	path := dbus.ObjectPath(exportInfo.Path)
	signalName := orgFreedesktopDBus + ".Properties.PropertiesChanged"
	propMap := make(map[string]dbus.Variant)
	propMap[propertyName] = dbus.MakeVariant(value)
	return s.conn.Emit(path, signalName, exportInfo.Interface, propMap, []string{})
}

func (s *Service) EmitPropertiesChanged(v Exportable, propValMap map[string]interface{},
	invalidatedProps ...string) error {

	exportInfo := v.GetDBusExportInfo()
	impl := s.getImplementer(exportInfo)
	path := dbus.ObjectPath(exportInfo.Path)
	signalName := orgFreedesktopDBus + ".Properties.PropertiesChanged"
	var changedProps map[string]dbus.Variant
	if len(propValMap) > 0 {
		changedProps = make(map[string]dbus.Variant)
	}
	for propName, val := range propValMap {
		err := impl.checkPropertyValue(propName, val)
		if err != nil {
			return err
		}
		changedProps[propName] = dbus.MakeVariant(val)
	}
	for _, propName := range invalidatedProps {
		if _, ok := propValMap[propName]; ok {
			return errors.New("property appears in both propValMap and invalidateProps")
		}

		err := impl.checkProperty(propName)
		if err != nil {
			return err
		}
	}
	return s.conn.Emit(path, signalName, exportInfo.Interface, changedProps, invalidatedProps)
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
					s.hasCallMu.Lock()
					hasCall := s.hasCall
					s.hasCallMu.Unlock()

					if !hasCall {
						if s.canQuit == nil || s.canQuit() {
							s.Quit()
							return
						}
					} else {
						s.hasCallMu.Lock()
						s.hasCall = false
						s.hasCallMu.Unlock()
					}
				}
			}
		}()
	}
	<-s.quit
}

func (s *Service) DelayAutoQuit() {
	s.hasCallMu.Lock()
	s.hasCall = true
	s.hasCallMu.Unlock()
}

func (s *Service) SetAutoQuitHandler(interval time.Duration, canQuit func() bool) {
	s.quitCheckInterval = interval
	s.canQuit = canQuit
}

func (s *Service) SetWriteCallback(v Exportable, propertyName string,
	cb PropertyWriteCallback) error {

	exportInfo := v.GetDBusExportInfo()
	impl := s.getImplementer(exportInfo)
	if impl == nil {
		return errors.New("not exported")
	}

	return impl.setWriteCallback(propertyName, cb)
}

func (s *Service) SetReadCallback(v Exportable, propertyName string,
	cb PropertyReadCallback) error {

	exportInfo := v.GetDBusExportInfo()
	impl := s.getImplementer(exportInfo)
	if impl == nil {
		return errors.New("not exported")
	}

	return impl.setReadCallback(propertyName, cb)
}

func (s *Service) ConnectChanged(v Exportable, propertyName string,
	cb PropertyChangedCallback) error {

	exportInfo := v.GetDBusExportInfo()
	impl := s.getImplementer(exportInfo)
	if impl == nil {
		return errors.New("not exported")
	}

	return impl.connectChanged(propertyName, cb)
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

func (s *Service) DumpProperties(v Exportable) (string, error) {
	exportInfo := v.GetDBusExportInfo()
	impl := s.getImplementer(exportInfo)
	if impl == nil {
		return "", errors.New("not exported")
	}

	var buf bytes.Buffer
	impl.propsMu.RLock()

	for propName, fieldProp := range impl.props {
		fmt.Fprintln(&buf, "property name:", propName)
		fmt.Fprintf(&buf, "mu: %p\n", fieldProp.valueMu)
		if fieldProp.valueMu != nil {
			fmt.Fprintln(&buf, "mu is PropsMu?", fieldProp.valueMu == impl.corePropsMu)
		}

		fmt.Fprintln(&buf, "signature:", fieldProp.signature)
		fmt.Fprintln(&buf, "access:", fieldProp.access)
		fmt.Fprintln(&buf, "emit:", fieldProp.emit)
		buf.WriteString("\n")
	}

	impl.propsMu.RUnlock()

	return buf.String(), nil
}
