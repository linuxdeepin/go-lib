// Theme settings.
package theme

// #cgo pkg-config: x11 xcursor xfixes gtk+-3.0
// #include <stdlib.h>
// #include "cursor.h"
import "C"

import (
	"fmt"
	"path"
	"pkg.deepin.io/lib/glib-2.0"
	dutils "pkg.deepin.io/lib/utils"
	"unsafe"
)

const (
	wmSchema        = "com.deepin.wrap.gnome.desktop.wm.preferences"
	xsettingsSchema = "com.deepin.xsettings"

	xsKeyTheme      = "theme-name"
	xsKeyIconTheme  = "icon-theme-name"
	xsKeyCursorName = "gtk-cursor-theme-name"
)

func SetGtkTheme(name string) error {
	if !IsThemeInList(name, ListGtkTheme()) {
		return fmt.Errorf("Invalid theme '%s'", name)
	}

	old := getXSettingsValue(xsKeyTheme)
	if old == name {
		return nil
	}

	if !setXSettingsKey(xsKeyTheme, name) {
		return fmt.Errorf("Set theme to '%s' by xsettings failed",
			name)
	}

	if !setWMTheme(name) {
		setXSettingsKey(xsKeyTheme, old)
		return fmt.Errorf("Set wm theme to '%s' failed", name)
	}

	if !setQTTheme(name) {
		setXSettingsKey(xsKeyTheme, old)
		setWMTheme(old)
		return fmt.Errorf("Set qt theme to '%s' failed", name)
	}
	return nil
}

func SetIconTheme(name string) error {
	if !IsThemeInList(name, ListIconTheme()) {
		return fmt.Errorf("Invalid theme '%s'", name)
	}

	old := getXSettingsValue(xsKeyIconTheme)
	if old == name {
		return nil
	}

	if !setXSettingsKey(xsKeyIconTheme, name) {
		return fmt.Errorf("Set theme to '%s' by xsettings failed",
			name)
	}
	return nil
}

func SetCursorTheme(name string) error {
	if !IsThemeInList(name, ListCursorTheme()) {
		return fmt.Errorf("Invalid theme '%s'", name)
	}

	old := getXSettingsValue(xsKeyCursorName)
	if old == name {
		return nil
	}

	if !setXSettingsKey(xsKeyCursorName, name) {
		return fmt.Errorf("Set theme to '%s' by xsettings failed",
			name)
	}

	if !setQtCursor(name) {
		setXSettingsKey(xsKeyCursorName, name)
		return fmt.Errorf("Set qt theme to '%s' failed", name)
	}

	handleGtkCursorChanged()
	return nil
}

func getXSettingsValue(key string) string {
	xs, err := dutils.CheckAndNewGSettings(xsettingsSchema)
	if err != nil {
		return ""
	}
	defer xs.Unref()
	return xs.GetString(key)
}

func setXSettingsKey(key, value string) bool {
	xs, err := dutils.CheckAndNewGSettings(xsettingsSchema)
	if err != nil {
		return false
	}
	defer xs.Unref()
	return xs.SetString(key, value)
}

func setWMTheme(name string) bool {
	wm, err := dutils.CheckAndNewGSettings(wmSchema)
	if err != nil {
		return false
	}
	defer wm.Unref()
	return wm.SetString("theme", name)
}

func setQTTheme(name string) bool {
	config := path.Join(glib.GetUserConfigDir(), "Trolltech.conf")
	return setQt4Theme(config)
}

func setQt4Theme(config string) bool {
	value, _ := dutils.ReadKeyFromKeyFile(config, "Qt", "style", "")
	if value == "GTK+" {
		return true
	}
	return dutils.WriteKeyToKeyFile(config, "Qt", "style", "GTK+")
}

func setQtCursor(name string) bool {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))

	ret := C.set_qt_cursor(cName)
	if ret != 0 {
		return false
	}
	return true
}

func handleGtkCursorChanged() {
	C.handle_gtk_cursor_changed()
}
