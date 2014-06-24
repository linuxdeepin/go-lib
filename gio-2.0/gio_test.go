package gio

import "fmt"
import "testing"

import "pkg.linuxdeepin.com/lib/glib-2.0"
import "pkg.linuxdeepin.com/lib/gio-2.0"

func TestAppInfo(t *testing.T) {
	apps := gio.AppInfoGetAll()
	for i, app := range apps {
		fmt.Println(i, app.GetSupportedTypes())
	}
}

func TestKeyFile(t *testing.T) {
	f := glib.NewKeyFile()
	f.LoadFromFile("/usr/share/applications/qtcreator.desktop", glib.KeyFileFlagsNone)
	s, _ := f.GetString("Desktop Entry", "MimeType")
	fmt.Println(s)
}
