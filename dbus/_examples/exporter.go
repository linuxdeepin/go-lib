package main

import "fmt"
import "dlib/dbus"

type DesktopManager struct {
	Changed func(int32)
	Name    string
}

func (m *DesktopManager) ListAutoStart() []string {
	return []string{}
}

func (m DesktopManager) OnPropertiesChanged(name string, old interface{}) {
	fmt.Println("Property ", name, " changed  ", old, " =====>", m.Name)
}
func (m *DesktopManager) GetDBusInfo() dbus.DBusInfo {
	return dbus.DBusInfo{
		"com.deepin.dde.Desktopmanager",
		"/com/deepin/dde/Desktopmananger",
		"com.deepin.dde.Desktopmananger",
	}
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
	return dbus.DBusInfo{
		"com.deepin.dde.Desktopmanager",
		"/com/deepin/dde/Desktopmananger/Entry" + e.dbusID,
		"com.deepin.dde.Desktopmananger.DesktopEntry",
	}
}

func (e *DesktopEntry) GetLocaleString(field string) {
}

func main() {
	m := DesktopManager{}
	m.Name = "snyh"
	dbus.InstallOnSession(&m)
	/*m.Changed(3) //emit an signal*/
	select {}
}
