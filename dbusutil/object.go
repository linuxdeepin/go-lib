package dbusutil

import (
	"sync"

	"pkg.deepin.io/lib/dbus1"
	"pkg.deepin.io/lib/dbus1/introspect"
	"pkg.deepin.io/lib/strv"
)

type object struct {
	path dbus.ObjectPath

	implementers map[string]*implementer
	//                ^interfaceName
	implementersMu sync.RWMutex

	propertiesImpl     *propertiesImplementer
	introspectableImpl *introspectableImplementer

	children strv.Strv // node name of children
}

func newObject(path dbus.ObjectPath, service *Service) *object {
	obj := &object{
		path:         path,
		implementers: make(map[string]*implementer),

		propertiesImpl: &propertiesImplementer{
			service: service,
		},

		introspectableImpl: &introspectableImplementer{
			service:    service,
			path:       path,
			interfaces: make(map[string]introspect.Interface),
		},
	}

	obj.propertiesImpl.object = obj
	return obj
}

func (obj *object) addImplementer(impl *implementer) {
	obj.implementersMu.Lock()
	obj.implementers[impl.interfaceName] = impl
	obj.implementersMu.Unlock()

	obj.introspectableImpl.addImplementer(impl)
}

func (obj *object) deleteImplementer(interfaceName string) {
	obj.implementersMu.Lock()
	delete(obj.implementers, interfaceName)
	obj.implementersMu.Unlock()

	obj.introspectableImpl.deleteImplementer(interfaceName)
}

func (obj *object) numImplementer() int {
	obj.implementersMu.RLock()
	n := len(obj.implementers)
	obj.implementersMu.RUnlock()
	return n
}

func (obj *object) hasImplementer(interfaceName string) bool {
	obj.implementersMu.RLock()
	_, ok := obj.implementers[interfaceName]
	obj.implementersMu.RUnlock()
	return ok
}

func (obj *object) getImplementer(interfaceName string) *implementer {
	obj.implementersMu.RLock()
	impl := obj.implementers[interfaceName]
	obj.implementersMu.RUnlock()
	return impl
}

func (obj *object) exportProperties(conn *dbus.Conn) error {
	return conn.Export(obj.propertiesImpl, obj.path, orgFreedesktopDBus+".Properties")
}

func (obj *object) stopExportProperties(conn *dbus.Conn) error {
	return conn.Export(nil, obj.path, orgFreedesktopDBus+".Properties")
}

func (obj *object) exportIntrospectable(conn *dbus.Conn) error {
	return conn.Export(obj.introspectableImpl, obj.path, orgFreedesktopDBus+".Introspectable")
}

func (obj *object) stopExportIntrospectable(conn *dbus.Conn) error {
	return conn.Export(nil, obj.path, orgFreedesktopDBus+".Introspectable")
}
