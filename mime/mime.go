package mime

import (
	"fmt"
	"pkg.deepin.io/lib/gio-2.0"
	"pkg.deepin.io/lib/theme/checker"
	dutils "pkg.deepin.io/lib/utils"
)

const (
	MimeTypeGtk    = "application/x-gtk-theme"
	MimeTypeIcon   = "application/x-icon-theme"
	MimeTypeCursor = "application/x-cursor-theme"

	mimeTypeTheme = "application/x-theme"
)

func Query(uri string) (string, error) {
	file := dutils.DecodeURI(uri)
	if !dutils.IsFileExist(file) {
		// 'cursor.theme' may not exist
		cursor, _ := checker.IsCursorTheme(file)
		if cursor {
			return MimeTypeCursor, nil
		}
		return "", fmt.Errorf("Not found the file '%s'", uri)
	}

	mime, err := doQueryFile(file)
	if err != nil {
		return "", err
	}

	if mime != mimeTypeTheme {
		return mime, nil
	}

	return queryThemeMime(file)
}

// Set 'mime' default app to 'desktopId'
//
// desktopId: the basename of the desktop file
func SetDefaultApp(mime, desktopId string) error {
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

func queryThemeMime(file string) (string, error) {
	gtk, _ := checker.IsGtkTheme(file)
	if gtk {
		return MimeTypeGtk, nil
	}

	icon, _ := checker.IsIconTheme(file)
	if icon {
		return MimeTypeIcon, nil
	}

	cursor, _ := checker.IsCursorTheme(file)
	if cursor {
		return MimeTypeCursor, nil
	}

	return "", fmt.Errorf("The mime of '%s' not supported", file)
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
