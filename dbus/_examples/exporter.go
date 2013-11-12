package main

import "dlib/dbus"
import "dlib/dbus/property"
import "dlib"
import "fmt"

const (
	_DEST = "com.deepin.dde.DestkopManager"
	_PATH = "/com/deepin/dde/DesktopManager"
	_IFC  = "com.deepin.dde.DesktopManager"
)

type DesktopManager struct {
	Changed  func(int32)
	Name     string
	DockMode dbus.Property
}

func (m *DesktopManager) ListAutoStart() []string {
	return []string{}
}

func (m DesktopManager) OnPropertiesChanged(name string, old interface{}) {
	fmt.Println("Property ", name, " changed  ", old, " =====>", m.Name)
}
func (m *DesktopManager) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{_DEST, _PATH, _IFC}
}

//auto return an sub object to dbus
func (m *DesktopManager) GetDesktopEntryById() []*DesktopEntry {
	return []*DesktopEntry{NewDesktopEntry("b"), NewDesktopEntry("c")}
}
func (m *DesktopManager) GetOne() *DesktopEntry {
	return NewDesktopEntry("a")
}

type DesktopEntry struct {
	ID          string
	Name        string
	Exec        string
	Description string
	IsAutoStart bool
	IsTerimal   bool
	dbusID      string
}

func NewDesktopEntry(id string) *DesktopEntry {
	return &DesktopEntry{
		dbusID: id,
	}
}

func (e *DesktopEntry) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{_DEST, _PATH + "/Entry" + e.dbusID, _IFC + ".DesktopEntry"}
}

func (e *DesktopEntry) GetLocaleString(field string) {
}

func main() {
	c, _ := dbus.SessionBus()
	s := dlib.NewSettings("com.deepin.dde.dock")
	m := DesktopManager{}

	// auto proxy an gsettings property
	m.DockMode = property.NewGSettingsPropertyFull(s, "hide-mode", "" /*this value's type must be the second paramter type*/, c, _PATH, _IFC, "DockMode" /*export Name*/)

	m.Name = "snyh"
	dbus.InstallOnSession(&m)
	ca := c.Object(_DEST, _PATH).Call("org.freedesktop.DBus.Introspectable.Introspect", 0)
	fmt.Println(ca)
	/*m.Changed(3) //emit an signal*/
	/*select {}*/
	dlib.StartLoop()
}
