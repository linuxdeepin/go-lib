package gio

import "testing"

func TestAppInfo(t *testing.T) {
	apps := AppInfoGetAll()
	for _, app := range apps {
		app.GetSupportedTypes()
		app.GetIcon()
	}
}
