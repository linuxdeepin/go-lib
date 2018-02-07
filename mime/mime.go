/*
 * Copyright (C) 2014 ~ 2018 Deepin Technology Co., Ltd.
 *
 * Author:     jouyouyun <jouyouwen717@gmail.com>
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program.  If not, see <http://www.gnu.org/licenses/>.
 */

package mime

import (
	"fmt"
	"gir/gio-2.0"
	dutils "pkg.deepin.io/lib/utils"
)

const (
	MimeTypeGtk    = "application/x-gtk-theme"
	MimeTypeIcon   = "application/x-icon-theme"
	MimeTypeCursor = "application/x-cursor-theme"

	mimeTypeTheme = "application/x-theme"
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
	cur, _ := GetDefaultApp(mime, false)
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

// Get default app for 'mime'
//
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
