package dlib

import "fmt"
import "testing"

import "dlib/glib-2.0"
import "dlib/gio-2.0"

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
