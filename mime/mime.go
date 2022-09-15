// SPDX-FileCopyrightText: 2018 - 2022 UnionTech Software Technology Co., Ltd.
//
// SPDX-License-Identifier: GPL-3.0-or-later

package mime

import (
	"fmt"
	"os"
	"path/filepath"

	dutils "github.com/linuxdeepin/go-lib/utils"
	gio "github.com/linuxdeepin/go-gir/gio-2.0"
)

const (
	MimeTypeGtk    = "application/x-gtk-theme"
	MimeTypeIcon   = "application/x-icon-theme"
	MimeTypeCursor = "application/x-cursor-theme"

	MimeUserConfig = ".config/mimeapps.list"
)

// Query query file mime type
func Query(uri string) (string, error) {
	file := dutils.DecodeURI(uri)

	// determine whether theme type
	mime, err := queryThemeMime(file)
	if err == nil {
		return mime, nil
	}

	return doQueryFile(file)
}

// Set 'mime' default app to 'desktopId'
//
// desktopId: the basename of the desktop file
func SetDefaultApp(mime, desktopId string) error {
	// If the default app is in /usr/share/applications/, still go to set
	cur := GetUserDefaultApp(mime)
	if cur == desktopId {
		return nil
	}

	app := gio.NewDesktopAppInfo(desktopId)
	if app == nil {
		return fmt.Errorf("Invalid id '%v'", desktopId)
	}
	defer app.Unref()

	_, err := app.SetAsDefaultForType(mime)
	return err
}

// get default from ~/.config/mimeapps.list
func GetUserDefaultApp(ty string) string {
	v, exist := dutils.ReadKeyFromKeyFile(filepath.Join(os.Getenv("HOME"), MimeUserConfig),
		"Default Applications", ty, "")
	if !exist {
		return ""
	}

	ret, ok := v.(string)
	if !ok {
		return ""
	}

	return ret
}

// Get default app for 'mime'
// gio.AppInfoGetDefaultForType can get default from ~/.config/mimeapps.list ， /usr/share/applications/mimeinfo.cache , /usr/share/applications/mimeapps.list
// ret0: desktopId
func GetDefaultApp(mime string, mustSupportURIs bool) (string, error) {
	app := gio.AppInfoGetDefaultForType(mime, false)
	if app == nil {
		return "", fmt.Errorf("Invalid mime '%v'", mime)
	}
	defer app.Unref()

	if mustSupportURIs {
		if !app.SupportsUris() {
			return "", fmt.Errorf("Not found app supported '%s' and uris", mime)
		}
	}

	return app.GetId(), nil
}

// Get app list of supported the 'mime'
// ret0: desktopId list
func GetAppList(mime string) []string {
	apps := gio.AppInfoGetAllForType(mime)

	var list []string
	for _, app := range apps {
		list = append(list, app.GetId())
		app.Unref()
	}
	return list
}

func doQueryFile(file string) (string, error) {
	if !dutils.IsFileExist(file) {
		return "", fmt.Errorf("Not found the file '%s'", file)
	}

	gf := gio.FileNewForPath(file)
	defer gf.Unref()

	info, err := gf.QueryInfo(gio.FileAttributeStandardContentType,
		gio.FileQueryInfoFlagsNone, nil)
	if err != nil {
		return "", err
	}
	defer info.Unref()

	return info.GetAttributeString(
		gio.FileAttributeStandardContentType), nil
}
